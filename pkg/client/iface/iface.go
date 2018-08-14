package iface

import (
	"github.com/9to6/aws-autoscaling-tester/pkg/client/config"
	"github.com/9to6/aws-autoscaling-tester/pkg/client/worker"
	llog "github.com/sirupsen/logrus"
)

type Client interface {
	Logger() *llog.Entry
	Config() *config.Config
	Run() error
	SetWorker(*worker.Worker)
}
