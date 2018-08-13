package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	llog "github.com/sirupsen/logrus"
	"github.com/9to6/aws-autoscaling-tester/pkg/server/conn"
)

const (
	connectionValue = 100
)

func GetConnectionHandler(c *gin.Context) {
	l_, ok := c.Get("logger")
	if ok == false {
		fmt.Println("error--")
		return
	}
	l := l_.(*llog.Logger)

	conn.IncreaseConnectionCount(connectionValue)
	// cloudwatch_, ok := c.Get("cloudWatch")
	// if ok == false {
	// 	return
	// }
	// cloudWatch := cloudwatch_.(*cloudwatch.CloudWatch)
	// if err := setConnectionMetric(cloudWatch); err != nil {
	// 	l.Error(err)
	// }
	l.Info("complete put connection metrics, current count: ", conn.GetConnectionCount())
}
