package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	currentSession *session.Session = nil
)

func NewAwsConnection(region, accesskey, secretkey string) (*session.Session, error) {
	conf := newConfig(region, accesskey, secretkey)
	conf.DisableSSL = aws.Bool(true)
	options := session.Options{
		Config: *config,
	}

	// os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
	// options.Profile = cfg.AWS.Profile

	currentSession, err = session.NewSessionWithOptions(options)
	if err != nil {
		return nil, err
	}

	return currentSession, nil
}
