package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("Minha primeira API em GO")

	router := gin.Default()

	router.SetTrustedProxies(nil)

	RegisterRoutes(router)

	router.Run(":3000")
}
