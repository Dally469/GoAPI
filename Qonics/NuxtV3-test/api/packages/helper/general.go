package helper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/mail"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/dally469/api/packages/config"
	"github.com/dally469/api/packages/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type DynamicSelect struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

var ctx = context.Background()
var SessionExpirationTime time.Duration = 30

func RequestAppendHeader(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if c.Request.Method == "OPTIONS" {
		c.JSON(200, gin.H{"success": 1})
		panic("done")
	}
}
func SecurePath(c *gin.Context) *models.Payload {
	token := c.GetHeader("Authorization")
	RequestAppendHeader(c)
	client := []byte(config.Redis.Get(ctx, token).Val())
	if client == nil {
		c.JSON(401, gin.H{"message": "Token not found or expired", "status": 401})
		panic("Token not found or expired")
	}
	var logger models.Payload
	err := json.Unmarshal(client, &logger)
	if err != nil {
		c.JSON(401, gin.H{"message": "Authentication failed, invalid token", "status": 401})
		panic("done, secure path failed #unmarshal" + err.Error())
	}
	config.Redis.Expire(ctx, token, time.Duration(SessionExpirationTime*time.Minute))
	return &logger

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func DynamicSelectData(c *gin.Context) {
	payload := SecurePath(c)
	data, _ := json.Marshal(payload)
	fmt.Printf("Payload %v", data)
	var list []DynamicSelect
	list = append(list, DynamicSelect{Value: "search-staff", Label: "Staffs"})
	list = append(list, DynamicSelect{Value: "search-student", Label: "Students"})
	c.JSON(200, gin.H{"status": 200, "records": list})
	return
}

func GenerateRandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789$&*!@%#"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func SendEmail(email string, subject string, message string, attachment *string) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "UHTGL<no-reply@uhtgl.org>")
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", message)
	if attachment != nil {
		msg.Attach(*attachment)
	}

	n := gomail.NewDialer("mail.uhtgl.org", 465, "no-reply@uhtgl.org", "qt*NL2$%1sEBC/u")

	//Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
}
func ParseTemplate(templateFileName string, data interface{}) (*string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return nil, err
	}
	body := buf.String()
	return &body, nil
}
