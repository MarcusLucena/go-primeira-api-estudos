package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("Minha primeira API em GO")

	router := gin.Default()

	router.SetTrustedProxies(nil)

	DB = initDB("tasks.db")
	defer DB.Close()

	RegisterRoutes(router)

	router.Run(":3000")
}
