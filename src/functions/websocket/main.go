package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"poll-to-push/src/internal/platform/config"
	"poll-to-push/src/internal/platform/env"
	"poll-to-push/src/internal/task"
)

func main() {
	cfg := config.Must(config.New())
	db := dynamodb.New(cfg)
	taskStore := task.NewStore(db, env.Get(env.TABLE_NAME))

	lambda.Start(NewHandler(taskStore))
}
