package mock

//go:generate mockgen  -package=mock -destination=./sfniface.go github.com/aws/aws-sdk-go-v2/service/sfn/sfniface ClientAPI
