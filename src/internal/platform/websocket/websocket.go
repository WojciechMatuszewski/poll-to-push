package websocket

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi/apigatewaymanagementapiiface"
)

type Service struct {
	api apigatewaymanagementapiiface.ClientAPI
}

func NewService(api apigatewaymanagementapiiface.ClientAPI) *Service {
	return &Service{api: api}
}

func (s *Service) Send(ctx context.Context, ID string, data []byte) error {
	req := s.api.PostToConnectionRequest(&apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(ID),
		Data:         data,
	})

	_, err := req.Send(ctx)
	return err
}
