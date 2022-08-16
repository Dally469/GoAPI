package main

import (
	"fmt"

	"github.com/dally469/api/packages/config"
	"github.com/dally469/api/packages/routes"
	
	"time"
    "math/rand"
)

func main() {
	fmt.Println("Hello - Qollege")
	config.InitializeConfig()
	config.ConnectDb()
	defer config.DB.Close()
	// controller.SocketConnection()
	rand.Seed(time.Now().UnixNano())
	server := routes.InitRoutes()
	server.Run(":8080")
}
