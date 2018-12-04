package main

import (
    "fmt"
    "log"
    "github.com/gin-gonic/gin"
)

func main() {
    log.Println("Websocket App start.")
    r := gin.Default()

    r.GET("/ping", handlePing)
    r.GET("/chime", handleChime)

    r.GET("/led", handleLed)
    handleSocket(r)

    // http://dex.local:8080/public/test.mp3
    r.Static("/public", "./public")
    r.Run(":8080")
    fmt.Println("Websocket App End.")
}
