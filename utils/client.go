package utils

import (
	cfg "aws-proxy/config"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/spf13/viper"
)

func EC2Client() *ec2.Client {

	Path := viper.GetString("config")
	Config := cfg.NewConfig(Path)

	KeyId := Config.Aws.AccessKeyID
	SecretKey := Config.Aws.SecretKeyID
	Region := Config.Aws.Region

	// Create a new AWS session with the provided credentials and region
	configs, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(Region),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     KeyId,
				SecretAccessKey: SecretKey,
			},
		},
		))
	if err != nil {
		panic(err)
	}
	client := ec2.NewFromConfig(configs)
	return client
}
