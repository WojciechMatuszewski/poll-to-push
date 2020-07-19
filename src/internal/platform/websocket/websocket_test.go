package websocket_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"poll-to-push/src/internal/platform/websocket"
	"poll-to-push/src/internal/platform/websocket/mock"
)

func TestService_Send(t *testing.T) {
	ctx := context.Background()

	t.Run("failure", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		api := mock.NewMockClientAPI(ctrl)
		service := websocket.NewService(api)

		req := apigatewaymanagementapi.PostToConnectionRequest{
			Request: &aws.Request{
				Data:        &apigatewaymanagementapi.PostToConnectionOutput{},
				Error:       errors.New("boom"),
				HTTPRequest: &http.Request{},
				Retryer:     aws.NoOpRetryer{},
			},
			Input: nil,
			Copy:  nil,
		}
		api.EXPECT().PostToConnectionRequest(&apigatewaymanagementapi.PostToConnectionInput{
			ConnectionId: aws.String("123"),
			Data:         []byte(`{"greeting":"Hi Wojtek!"}`),
		}).Return(req)

		err := service.Send(ctx, "123", []byte(`{"greeting":"Hi Wojtek!"}`))
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		api := mock.NewMockClientAPI(ctrl)
		service := websocket.NewService(api)

		req := apigatewaymanagementapi.PostToConnectionRequest{
			Request: &aws.Request{
				Data:        &apigatewaymanagementapi.PostToConnectionOutput{},
				Error:       nil,
				HTTPRequest: &http.Request{},
				Retryer:     aws.NoOpRetryer{},
			},
			Input: nil,
			Copy:  nil,
		}
		api.EXPECT().PostToConnectionRequest(&apigatewaymanagementapi.PostToConnectionInput{
			ConnectionId: aws.String("123"),
			Data:         []byte(`{"greeting":"Hi Wojtek!"}`),
		}).Return(req)

		err := service.Send(ctx, "123", []byte(`{"greeting":"Hi Wojtek!"}`))
		assert.NoError(t, err)
	})
}
