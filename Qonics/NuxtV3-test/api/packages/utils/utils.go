package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dally469/api/packages/helper"
	"github.com/gin-gonic/gin"
)

func ParseBody(r *http.Request, x interface{}){
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if  err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
func CorsReply(c *gin.Context) {
	// time.Sleep(5 * time.Second)
	helper.RequestAppendHeader(c)
}
