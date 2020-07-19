package mock

//go:generate mockgen -package=mock -destination=apigatewaymanagementapiiface.go github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi/apigatewaymanagementapiiface ClientAPI
