package client

import (
	"fmt"
	"github.com/9to6/aws-autoscaling-tester/pkg/client/config"
	"github.com/9to6/aws-autoscaling-tester/pkg/client/handler"
	"github.com/9to6/aws-autoscaling-tester/pkg/client/worker"
	"github.com/9to6/aws-autoscaling-tester/pkg/log"
	"github.com/9to6/gin-logrus"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	llog "github.com/sirupsen/logrus"
)

type Client struct {
	logger *llog.Entry
	router *gin.Engine
	worker *worker.Worker
	config *config.Config
}

func (s *Client) Logger() *llog.Entry {
	return s.logger
}

func (s *Client) Config() *config.Config {
	return s.config
}

func (s *Client) Run() error {
	s.worker = worker.NewWorker(*s.config, s.Logger())
	s.worker.StartWork()
	return s.router.Run(fmt.Sprintf(":%d", s.config.Port))
}

func (s *Client) SetWorker(w *worker.Worker) {
	s.worker.StopWork()
	s.worker = w
	s.worker.StartWork()
}

func NewClient(config *config.Config) *Client {
	client := &Client{
		logger: log.Logger.WithFields(llog.Fields{"client": "rest"}),
		config: config,
	}
	router := gin.New()
	router.Use(location.Default())
	router.Use(ginlogrus.Logger(log.Logger))
	router.Use(middlewarePredispatch(client))
	client.router = router

	router.GET("/status", func(c *gin.Context) { c.Writer.Write([]byte(`success`)) })
	router.POST("/worker", handler.PostCreateWorker)
	return client
}

func middlewarePredispatch(s *Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("logger", s.Logger())
		ctx.Set("client", s)
		ctx.Next()
	}
}
