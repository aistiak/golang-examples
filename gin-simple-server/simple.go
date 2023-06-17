package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

func RunSimpleGinServer() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	fmt.Println("---running server ---")
	router.Run("localhost:8000")
}
