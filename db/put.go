package db

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func PutDynamodb(id int, title string) {
	ddb := dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))

	param := &dynamodb.PutItemInput{
		TableName: aws.String("Golang-CreateTable"),
		Item: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String(strconv.Itoa(id)), // データ型(Number:N)
			},
			"TITLE": {
				S: aws.String(title), //データ型(String:S)
			},
		},
	}

	_, err := ddb.PutItem(param) //実行

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Put a first Item")
}
