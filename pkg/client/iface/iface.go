package iface

import (
	llog "github.com/sirupsen/logrus"
	"metric-generator/pkg/client/worker"
)

type Client interface {
	Logger() *llog.Entry
	Run() error
	SetWorker(*worker.Worker)
}
