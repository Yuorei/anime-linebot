package main

import (
    "log"
    //"os"
    "strings"

    "github.com/gin-gonic/gin"

    "github.com/line/line-bot-sdk-go/linebot"
)

func main() {
    port := "3000"//os.Getenv("PORT")

    if port == "" {
        log.Fatal("$PORT must be set")
    }
    bot, err := linebot.New(
        "ac1cef1ef1b189e2a25cc20cb6a666ce",
        "vJtxR7bXEHb4UNWMPeMv+kLgjdXMGGNUk8cFB26qQix+R/MIc+3E9fLiLvqiHPHEN/Bz8N2YikIDDoUGip5ZS/ZHeEjslBS4ouzTJUFX8y4saWp2Z/v7H9CAJbM3Uy6b7TjavQybANOpjV/PEDqsYgdB04t89/1O/w1cDnyilFU=",
    )
    if err != nil {
        log.Fatal(err)
    }

    router := gin.New()
    router.Use(gin.Logger())

    // LINE Messaging API ルーティング
    //router.POST("/callback", func(c *gin.Context) {
	router.POST("/webhook", func(c *gin.Context) {
        events, err := bot.ParseRequest(c.Request)
        if err != nil {
            if err == linebot.ErrInvalidSignature {
                log.Print(err)
            }
            return
        }
    })
    router.Run(":" + port)
}