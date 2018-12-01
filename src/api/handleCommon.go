package main

import "github.com/gin-gonic/gin"

func handlePing(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "ping",
    })
}
