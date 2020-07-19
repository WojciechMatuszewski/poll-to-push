package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pkg/errors"
	"poll-to-push/src/internal/task"
)

type Response struct {
	ID string `json:"id"`
}

type Handler func(ctx context.Context, request events.APIGatewayV2HTTPRequest) (
	events.APIGatewayV2HTTPResponse, error)

type StepFunctionStarter interface {
	Start(ctx context.Context, input, id string) (string, error)
}

func newHandler(starter StepFunctionStarter) Handler {
	return func(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		var input task.Input
		err := json.Unmarshal([]byte(request.Body), &input)
		if err != nil {
			return events.APIGatewayV2HTTPResponse{
					StatusCode: http.StatusInternalServerError,
					Body:       http.StatusText(http.StatusInternalServerError),
				},
				errors.Wrap(err, "error: while unmarshalling")
		}
		if !input.IsValid() {
			return events.APIGatewayV2HTTPResponse{
					StatusCode: http.StatusBadRequest,
					Body:       http.StatusText(http.StatusBadRequest),
				},
				fmt.Errorf("error: wrong payload %v", request.Body)
		}

		id, err := starter.Start(ctx, request.Body, request.RequestContext.RequestID)
		if err != nil {
			return events.APIGatewayV2HTTPResponse{
					StatusCode: http.StatusInternalServerError,
					Body:       http.StatusText(http.StatusInternalServerError),
				},
				errors.Wrap(err, "error: starting stepfunction")
		}

		resp := Response{
			ID: id,
		}
		buf, err := json.Marshal(&resp)
		if err != nil {
			return events.APIGatewayV2HTTPResponse{
					StatusCode: http.StatusInternalServerError,
					Body:       http.StatusText(http.StatusInternalServerError),
				},
				errors.Wrap(err, "error: marshaling input")
		}

		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusOK,
			Body:       string(buf),
		}, nil
	}
}
