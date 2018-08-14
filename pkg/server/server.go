package server

import (
	"github.com/9to6/aws-autoscaling-tester/pkg/log"
	"github.com/9to6/gin-logrus"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	llog "github.com/sirupsen/logrus"

	"fmt"
	"github.com/9to6/aws-autoscaling-tester/pkg/server/config"
	"github.com/9to6/aws-autoscaling-tester/pkg/server/conn"
	"github.com/9to6/aws-autoscaling-tester/pkg/server/handler"
	"github.com/9to6/aws-autoscaling-tester/pkg/server/timer"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
)

type Server struct {
	router     *gin.Engine
	logger     *llog.Logger
	config     *config.Config
	cloudWatch *cloudwatch.CloudWatch
}

func (s *Server) Logger() *llog.Logger {
	return s.logger
}

func (s *Server) Config() *config.Config {
	return s.config
}

func (s *Server) CloudWatch() *cloudwatch.CloudWatch {
	return s.cloudWatch
}

func middlewarePredispatch(s *Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("logger", s.Logger())
		ctx.Set("cloudWatch", s.CloudWatch())
		ctx.Set("config", s.Config())
		ctx.Next()
	}
}

func (s *Server) Run() error {
	timer.StartTimer(s.cloudWatch)
	return s.router.Run(fmt.Sprintf(":%d", s.config.Port))
}

func NewServer(conf config.Config) (*Server, error) {
	server := &Server{
		logger: log.Logger,
		config: &conf,
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

	conn.InitConnectionCount(float64(conf.ConnectionValue))
	return server, nil
}
