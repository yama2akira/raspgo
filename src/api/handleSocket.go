package main

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "gopkg.in/olahol/melody.v1"
)

func handlePage(ctx *gin.Context) {
    http.ServeFile(ctx.Writer, ctx.Request, "index.html")
}

func handleSocket(engine *gin.Engine) {
    m := melody.New()
    rg := engine.Group("/chat")
    rg.GET("/", handlePage);
    rg.GET("/ws", func(ctx *gin.Context) {
        m.HandleRequest(ctx.Writer, ctx.Request)
    })

    m.HandleMessage(func(s *melody.Session, msg []byte) {
        m.Broadcast(msg)
    })

    m.HandleConnect(func(s *melody.Session) {
        log.Printf("websocket connection open. [session: %#v]\n", s)
    })

    m.HandleDisconnect(func(s *melody.Session) {
        log.Printf("websocket connection close. [session: %#v]\n", s)
    })
}

