package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"poll-to-push/src/internal/task"
)

func handler(ctx context.Context, input task.Input) (task.Result, error) {
	time.Sleep(30 * time.Second)
	return task.Result{Greeting: fmt.Sprintf("Greeting %v!", input.Name)}, nil
}

func main() {
	lambda.Start(handler)
}
