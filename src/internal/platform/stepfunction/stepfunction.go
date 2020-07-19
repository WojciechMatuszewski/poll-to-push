package stepfunction

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/sfn/sfniface"
	"github.com/aws/aws-sdk-go/aws"
)

type Service struct {
	api        sfniface.ClientAPI
	machineArn string
}

func NewService(api sfniface.ClientAPI, machineArn string) *Service {
	return &Service{
		api:        api,
		machineArn: machineArn,
	}
}

func (s *Service) Start(ctx context.Context, input, id string) (string, error) {
	req := s.api.StartExecutionRequest(&sfn.StartExecutionInput{
		Input:           aws.String(input),
		Name:            aws.String(id),
		StateMachineArn: aws.String(s.machineArn),
	})

	out, err := req.Send(ctx)
	if err != nil {
		return "", err
	}

	return *out.ExecutionArn, nil
}
