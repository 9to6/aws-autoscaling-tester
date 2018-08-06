package aws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func newConfig(region, accesskey, secretkey string) *aws.Config {
	awsConfig := aws.NewConfig()
	awsConfig.WithRegion(region)

	if accesskey != "" {
		creds := credentials.NewStaticCredentials(accesskey, secretkey, "")
		awsConfig.Credentials = creds
	} else if os.Getenv("AWS_ACCESS_KEY") != "" || os.Getenv("AWS_ACCESS_KEY_ID") != "" {
		awsConfig.Credentials = credentials.NewEnvCredentials()
	}
	return awsConfig
}
