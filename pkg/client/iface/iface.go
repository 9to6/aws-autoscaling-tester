package iface

import (
	llog "github.com/sirupsen/logrus"
	"github.com/9to6/aws-autoscaling-tester/pkg/client/worker"
)

type Client interface {
	Logger() *llog.Entry
	Run() error
	SetWorker(*worker.Worker)
}
