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

    r.GET("/led", handleLed)
    handleSocket(r)

    r.Run(":8080")
    fmt.Println("Websocket App End.")
}
