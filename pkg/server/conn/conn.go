package conn

import (
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/aws"
	"time"
)

const (
	namespace  = "TESTMetrics"
	metricName = "connection"
	connectionValue = 1000
)

func SendConnectionMetric(cloudWatch *cloudwatch.CloudWatch) error {
	params := &cloudwatch.PutMetricDataInput{
		MetricData: []*cloudwatch.MetricDatum{
			&cloudwatch.MetricDatum{
				Dimensions: make([]*cloudwatch.Dimension, 0),
				MetricName: aws.String(metricName),
				Timestamp:  aws.Time(time.Now()),
				Value:      aws.Float64(connectionValue),
			}},
		Namespace: aws.String(namespace),
	}

	if _, err := cloudWatch.PutMetricData(params); err != nil {
		return err
	}
	return nil
}
