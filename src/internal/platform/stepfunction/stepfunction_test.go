package stepfunction_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"poll-to-push/src/internal/platform/stepfunction"
	"poll-to-push/src/internal/platform/stepfunction/mock"
)

func TestService_Start(t *testing.T) {
	ctx := context.Background()

	t.Run("start error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockAPI := mock.NewMockClientAPI(ctrl)
		service := stepfunction.NewService(mockAPI, "arn")

		req := sfn.StartExecutionRequest{
			Request: &aws.Request{
				Data:        &sfn.StartExecutionOutput{},
				Error:       errors.New("boom"),
				HTTPRequest: &http.Request{},
				Retryer:     aws.NoOpRetryer{},
			},
		}
		mockAPI.EXPECT().StartExecutionRequest(&sfn.StartExecutionInput{
			Input:           aws.String("foo"),
			Name:            aws.String("bar"),
			StateMachineArn: aws.String("arn"),
		}).Return(req)

		out, err := service.Start(ctx, "foo", "bar")
		assert.Error(t, err)
		assert.Empty(t, out)
	})

	t.Run("success", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockAPI := mock.NewMockClientAPI(ctrl)
		service := stepfunction.NewService(mockAPI, "arn")

		req := sfn.StartExecutionRequest{
			Request: &aws.Request{
				Data:        &sfn.StartExecutionOutput{ExecutionArn: aws.String("executionArn")},
				Error:       nil,
				HTTPRequest: &http.Request{},
				Retryer:     aws.NoOpRetryer{},
			},
		}
		mockAPI.EXPECT().StartExecutionRequest(&sfn.StartExecutionInput{
			Input:           aws.String("foo"),
			Name:            aws.String("bar"),
			StateMachineArn: aws.String("arn"),
		}).Return(req)

		out, err := service.Start(ctx, "foo", "bar")
		assert.NoError(t, err)
		assert.Equal(t, "executionArn", out)
	})
}
