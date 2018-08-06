package handler

import (
	"time"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gin-gonic/gin"
	llog "github.com/sirupsen/logrus"
	"fmt"
)

const (
	namespace  = "TESTMetrics"
	metricName = "connection"
	connectionValue = 100
)

func GetConnectionHandler(c *gin.Context) {
	l_, ok := c.Get("logger")
	if ok == false {
		fmt.Println("error--")
		return
	}
	l := l_.(*llog.Logger)

	cloudwatch_, ok := c.Get("cloudWatch")
	if ok == false {
		return
	}
	cloudWatch := cloudwatch_.(*cloudwatch.CloudWatch)
	if err := setConnectionMetric(cloudWatch); err != nil {
		l.Error(err)
	}
	l.Info("complete put connection metrics ")
}

func setConnectionMetric(cloudWatch *cloudwatch.CloudWatch) error {
	var val float64 = 1 * connectionValue;
	params := &cloudwatch.PutMetricDataInput{
		MetricData: []*cloudwatch.MetricDatum{
			&cloudwatch.MetricDatum{
				Dimensions: make([]*cloudwatch.Dimension, 0),
				MetricName: aws.String(metricName),
				Timestamp:  aws.Time(time.Now()),
				Value:      aws.Float64(val),
			}},
		Namespace: aws.String(namespace),
	}

	if _, err := cloudWatch.PutMetricData(params); err != nil {
		return err
	}
	return nil
}
