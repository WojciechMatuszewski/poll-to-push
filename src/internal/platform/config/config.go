package config

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
)

func New() (aws.Config, error) {
	return external.LoadDefaultAWSConfig()
}

func Must(cfg aws.Config, err error) aws.Config {
	if err != nil {
		panic(err.Error())
	}

	return cfg
}
