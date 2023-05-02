package transport

import (
	"log"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	log.Println(message)

	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*",)
	c.JSON(statusCode, Response{message})
}