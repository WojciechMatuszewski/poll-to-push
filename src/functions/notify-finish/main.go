package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"poll-to-push/src/internal/platform/config"
	"poll-to-push/src/internal/platform/env"
	"poll-to-push/src/internal/platform/websocket"
	"poll-to-push/src/internal/task"
)

func main() {
	cfg := config.Must(config.New())

	db := dynamodb.New(cfg)
	taskStore := task.NewStore(db, env.Get(env.TABLE_NAME))

	apigwCfg := cfg
	apigwCfg.EndpointResolver = aws.ResolveWithEndpointURL(env.Get(env.WEBSOCKET_ADDRESS))
	wsAPI := apigatewaymanagementapi.New(apigwCfg)
	wsService := websocket.NewService(wsAPI)

	lambda.Start(NewHandler(taskStore, wsService))
}
