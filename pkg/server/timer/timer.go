package timer

import (
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"syscall"
	"time"
		"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/9to6/cloudwatch-metrics-tester/pkg/server/conn"
	"github.com/9to6/cloudwatch-metrics-tester/pkg/log"
)

var ticker *time.Ticker

func StartTimer(cloudWatch *cloudwatch.CloudWatch) {
	// ctx, cancelFunc := context.WithCancel(context.Background())
	// shutdown(cancelFunc)

	ticker := time.NewTicker(60 * time.Second)
	go func() {
		for t := range ticker.C {
			log.Logger.Info("Ticker executed", t)
			conn.SendConnectionMetric(cloudWatch)
		}
	}()

}

func StopTimer() {
	ticker.Stop()
}

func shutdown(cancelFunc context.CancelFunc) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cancelFunc()
	}()
}
