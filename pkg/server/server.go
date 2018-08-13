package server

import (
	"github.com/9to6/gin-logrus"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	llog "github.com/sirupsen/logrus"
	"metric-generator/pkg/log"

	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"metric-generator/pkg/server/handler"
	"metric-generator/pkg/server/timer"
	"os"
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
	router.POST("/connection", handler.PostConnectionHandler)
	return server, nil
}
