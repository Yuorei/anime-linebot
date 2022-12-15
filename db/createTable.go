package db

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateDynamodbTable() {
	ddb := dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))

	params := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("ID"),
				AttributeType: aws.String("N"), // データ型(Number:N)
			},
			{
				AttributeName: aws.String("TITLE"),
				AttributeType: aws.String("S"), // データ型(String:S)
			},
		},
		BillingMode: aws.String("PAY_PER_REQUEST"), // キャパシティーモードをオンデマンドに指定
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       aws.String("HASH"), // HASH(Partition key)を設定
			},
			{
				AttributeName: aws.String("TITLE"),
				KeyType:       aws.String("RANGE"), // RANGE(Sort key)を設定
			},
		},
		TableName: aws.String("LineMovie"), // テーブル名
	}

	_, err := ddb.CreateTable(params)

	if err != nil {
		fmt.Println("Got error calling CreateTable:")
	} else {

		fmt.Println("Created the table")
	}
}
