package main

import (
	"log"
	"os"

	//"os"
	"strings"

	"github.com/Yuorei/withdrawal/db"
	"github.com/Yuorei/withdrawal/get_api"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/line/line-bot-sdk-go/linebot"
)
var id=0
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_ACCESS_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// dynamodbテーブルがない場合に作成
	db.CreateDynamodbTable()

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

		// "表示" 単語を含む場合、返信される
		replyImage := "表示"
		// 失敗
		//apiMiss := "失敗しました"
		// チャットの回答
		//responseSave := "を保存しました"

		for _, event := range events {
			// イベントがメッセージの受信だった場合
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				// メッセージがテキスト形式の場合
				case *linebot.TextMessage:
					replyMessage := message.Text
					// ここでDBの中身を表示する予定
					if strings.Contains(replyMessage, replyImage) {
						// todo
					}
					data := get_api.SearchTvGET(replyMessage)
					for _, item := range data.Results {
						// bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(item.Originalname)).Do()
						// 仕様でここで終了
						
						// 写真表示
						registerImage := linebot.NewImageMessage("https://image.tmdb.org/t/p/w500"+item.Posterpath, "https://image.tmdb.org/t/p/w500"+item.Posterpath)
						bot.ReplyMessage(event.ReplyToken, registerImage).Do()
						id++
						//db.PutDynamodb(id, item.Originalname)

					}
				}
			}
		}
	})
	router.Run(":" + port)
}