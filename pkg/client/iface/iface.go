package iface

import (
	llog "github.com/sirupsen/logrus"
	"aws-autoscaling-tester/pkg/client/worker"
)

type Client interface {
	Logger() *llog.Entry
	Run() error
	SetWorker(*worker.Worker)
}
