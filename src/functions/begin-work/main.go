package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"poll-to-push/src/internal/platform/config"
	"poll-to-push/src/internal/platform/env"
	"poll-to-push/src/internal/platform/stepfunction"
)

func main() {
	cfg := config.Must(config.New())

	sfnAPI := sfn.New(cfg)
	sfnService := stepfunction.NewService(sfnAPI, env.Get(env.MACHINE_ARN))
	h := newHandler(sfnService)
	lambda.Start(h)
}
