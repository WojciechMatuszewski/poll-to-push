package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pkg/errors"
	"poll-to-push/src/internal/task"
)

type Payload struct {
	Data struct {
		TaskID string `json:"taskID"`
	} `json:"data"`
}

type Handler func(ctx context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error)

type ConnectionSaver interface {
	Save(ctx context.Context, conn task.Connection) error
}

func NewHandler(saver ConnectionSaver) Handler {
	return func(ctx context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
		if request.Body == "" {
			return events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
		}
		var payload Payload
		err := json.Unmarshal([]byte(request.Body), &payload)
		if err != nil {
			return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, errors.Wrap(err, "error: unmarshalling clients payload")
		}

		conn := task.Connection{
			ID:     request.RequestContext.ConnectionID,
			TaskID: payload.Data.TaskID,
		}
		err = saver.Save(ctx, conn)
		if err != nil {
			return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, errors.Wrap(err, "error: save the connection")
		}

		return events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
	}
}
