package timer

import (
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"golang.org/x/net/context"
	"github.com/9to6/aws-autoscaling-tester/pkg/log"
	"github.com/9to6/aws-autoscaling-tester/pkg/server/conn"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var ticker *time.Ticker

func StartTimer(cloudWatch *cloudwatch.CloudWatch) {
	// ctx, cancelFunc := context.WithCancel(context.Background())
	// shutdown(cancelFunc)

	ticker := time.NewTicker(60 * time.Second)
	go func() {
		for t := range ticker.C {
			log.Logger.Info("Ticker executed: ", t)
			if err := conn.SendConnectionMetric(cloudWatch); err != nil {
				log.Logger.Error("error: ", err)
			} else {
				log.Logger.Info("To send connection metrics is done ")
			}
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
