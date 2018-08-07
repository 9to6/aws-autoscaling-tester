package conn

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"sync"
	"time"
)

const (
	namespace       = "TESTMetrics"
	metricName      = "connection"
	connectionValue = 1000
)

var (
	currentConnection float64 = connectionValue
	mutex                     = &sync.Mutex{}
)

func SendConnectionMetric(cloudWatch *cloudwatch.CloudWatch) error {
	mutex.Lock()
	val := currentConnection
	mutex.Unlock()
	params := &cloudwatch.PutMetricDataInput{
		MetricData: []*cloudwatch.MetricDatum{
			{
				Dimensions: make([]*cloudwatch.Dimension, 0),
				MetricName: aws.String(metricName),
				Timestamp:  aws.Time(time.Now()),
				Value:      aws.Float64(val),
			}},
		Namespace: aws.String(namespace),
	}

	// initialize connection value
	mutex.Lock()
	currentConnection = connectionValue
	mutex.Unlock()
	if _, err := cloudWatch.PutMetricData(params); err != nil {
		return err
	}
	return nil
}

func IncreaseConnectionCount(count float64) {
	mutex.Lock()
	currentConnection += count
	mutex.Unlock()
}

func GetConnectionCount() float64 {
	mutex.Lock()
	ret := currentConnection
	mutex.Unlock()
	return ret
}
