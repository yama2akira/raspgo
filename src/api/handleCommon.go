package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
   // "gobot.io/x/gobot"
    "gobot.io/x/gobot/drivers/gpio"
    "gobot.io/x/gobot/platforms/raspi"
    "os/exec"
   // "fmt"
)

func handlePing(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "pingです",
    })
}

func handleChime(c *gin.Context) {
    //start_gin.sh のあるフォルダがカレント
    err := exec.Command("sudo", "mpg321", "-a", "hw:1,0", "public/test.mp3").Start()
    if(err == nil) {
        c.String(http.StatusOK, "chime start!!")
    } else {
        c.String(http.StatusForbidden, "chime error!!")
    }
}

// https://make.kosakalab.com/rpi/raspberry-pi_golang/
var toggle bool

func handleLed(c *gin.Context) {
    toggle = !toggle
    r := raspi.NewAdaptor()
    led := gpio.NewLedDriver(r, "7")
    if(toggle) {
        led.On() 
    } else {
        led.Off()
    }
    //led.Toggle()

    status := map[bool]string{true: "On", false: "Off"}[toggle]
    c.String(http.StatusOK, "LED: " + status)
}
