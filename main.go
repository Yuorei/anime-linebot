package main

import (
	"log"

	//"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	port := "" //os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	bot, err := linebot.New(
		"CHANNEL_SECRET",
		"CHANNEL_ACCESS_TOKEN",
    )
	if err != nil {
		log.Fatal(err)
	}

	router := gin.New()
	router.Use(gin.Logger())

	// LINE Messaging API ルーティング
	router.POST("/webhook", func(c *gin.Context) {
		events, err := bot.ParseRequest(c.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				log.Print(err)
			}
			return
		}
		// "保存" 単語を含む場合、返信される
		replySave := "保存"
		// "保存" 単語を含む場合、返信される
		replyImage := "表示"
		// チャットの回答
		responseSave := "を保存しました"

		// DBからとってきた写真
		responseImage := linebot.NewImageMessage("", "")

		for _, event := range events {
			// イベントがメッセージの受信だった場合
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				// メッセージがテキスト形式の場合
				case *linebot.TextMessage:
					replyMessage := message.Text
					// テキストで返信されるケース
					if strings.Contains(replyMessage, replySave) {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(responseSave)).Do()

					} else if strings.Contains(replyMessage, replyImage) {
						bot.ReplyMessage(event.ReplyToken, responseImage).Do()
					}

					registerImage := linebot.NewImageMessage("", "")
					bot.ReplyMessage(event.ReplyToken, registerImage).Do()
				}
			}
		}
	})
	router.Run(":" + port)
}
