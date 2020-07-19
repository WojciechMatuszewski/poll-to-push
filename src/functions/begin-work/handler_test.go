package main

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"poll-to-push/src/functions/begin-work/mock"
)

func TestHandler(t *testing.T) {
	ctx := context.Background()

	t.Run("empty payload", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		starter := mock.NewMockStepFunctionStarter(ctrl)
		h := newHandler(starter)

		out, err := h(ctx, events.APIGatewayV2HTTPRequest{
			Body: `{"bar": "fo"}`,
			RequestContext: events.APIGatewayV2HTTPRequestContext{
				RequestID: "123",
			},
		})

		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, out.StatusCode)
	})

	t.Run("starter failure", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		starter := mock.NewMockStepFunctionStarter(ctrl)
		h := newHandler(starter)

		starter.EXPECT().Start(ctx, `{"name": "Wojtek"}`, "123").Return("", errors.New("boom"))
		out, err := h(ctx, events.APIGatewayV2HTTPRequest{
			Body: `{"name": "Wojtek"}`,
			RequestContext: events.APIGatewayV2HTTPRequestContext{
				RequestID: "123",
			},
		})

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, out.StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		starter := mock.NewMockStepFunctionStarter(ctrl)
		h := newHandler(starter)

		starter.EXPECT().Start(ctx, `{"name": "Wojtek"}`, "123").Return("997", nil)
		out, err := h(ctx, events.APIGatewayV2HTTPRequest{
			Body: `{"name": "Wojtek"}`,
			RequestContext: events.APIGatewayV2HTTPRequestContext{
				RequestID: "123",
			},
		})

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, out.StatusCode)
		assert.Equal(t, `{"id":"997"}`, out.Body)
	})
}
