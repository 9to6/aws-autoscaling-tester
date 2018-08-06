package server

import (
	"github.com/9to6/cloudwatch-metrics-tester/pkg/log"
	"github.com/9to6/gin-logrus"
	llog "github.com/sirupsen/logrus"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/aws/aws-sdk-go/service/cloudwatch"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"os"
	"fmt"
	"github.com/9to6/cloudwatch-metrics-tester/pkg/server/handler"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/9to6/cloudwatch-metrics-tester/pkg/server/timer"
)

type Server struct {
	router     *gin.Engine
	logger     *llog.Logger
	cloudWatch *cloudwatch.CloudWatch
}

func (s Server) Logger() *llog.Logger {
	return s.logger
}

func (s Server) CloudWatch() *cloudwatch.CloudWatch {
	return s.cloudWatch
}

func middlewarePredispatch(s *Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("logger", s.Logger())
		ctx.Set("cloudWatch", s.CloudWatch())
		ctx.Next()
	}
}
func (s *Server) Run() error {
	timer.StartTimer(s.cloudWatch)
	return s.router.Run(":8080")
}

func NewServer() (*Server, error) {
	server := &Server{
		logger: log.Logger,
	}
	server.logger.Info("aaaaaaa")

	router := gin.New()
	router.Use(location.Default())
	router.Use(ginlogrus.Logger(log.Logger))

	{
		awsConfig := &aws.Config{
			Region: aws.String("us-east-1"),
		}

		if os.Getenv("AWS_ACCESS_KEY") != "" || os.Getenv("AWS_ACCESS_KEY_ID") != "" {
			awsConfig.Credentials = credentials.NewEnvCredentials()
		}

		awsConfig.DisableSSL = aws.Bool(true)
		options := session.Options{
			Config: *awsConfig,
		}

		sess, err := session.NewSessionWithOptions(options)
		if err != nil {
			fmt.Println("sess error")
			return nil, err
		}
		server.cloudWatch = cloudwatch.New(sess)
	}

	server.router = router
	router.Use(middlewarePredispatch(server))

	router.GET("/status", func(c *gin.Context) { c.Writer.Write([]byte(`success`)) })
	router.GET("/connection", handler.GetConnectionHandler)
	return server, nil
}
