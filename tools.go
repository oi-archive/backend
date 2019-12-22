package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func ThrowUnknownError(c *gin.Context, err error) {
	log.Println(err)
	c.JSON(500, gin.H{
		"errorCode":    500,
		"errorMessage": "Unknown error.",
	})
}
