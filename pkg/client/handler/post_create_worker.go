package handler

import (
	"github.com/9to6/aws-autoscaling-tester/pkg/client/iface"
	"github.com/9to6/aws-autoscaling-tester/pkg/client/worker"
	"fmt"
	"github.com/gin-gonic/gin"
	llog "github.com/sirupsen/logrus"
	"strconv"
)

func PostCreateWorker(c *gin.Context) {
	l_, ok := c.Get("logger")
	if ok == false {
		c.AbortWithError(500, fmt.Errorf("get logger error"))
		return
	}
	l := l_.(*llog.Entry).WithFields(llog.Fields{"handler": "create_worker"})
	l.Info("post_create_worker")

	client_, ok := c.Get("client")
	if ok == false {
		c.AbortWithError(500, fmt.Errorf("get client error"))
		return
	}
	var period, conn int32
	var url string
	{
		periodStr, ok := c.GetPostForm("period")
		if !ok {
			c.AbortWithError(500, fmt.Errorf("you have to input 'period' params"))
			return
		}
		val, err := strconv.ParseInt(periodStr, 10, 32)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		period = int32(val)
	}
	{
		connStr, ok := c.GetPostForm("conn")
		if !ok {
			c.AbortWithError(500, fmt.Errorf("you have to input 'conn' params"))
			return
		}
		val, err := strconv.ParseInt(connStr, 10, 32)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		conn = int32(val)
	}
	{
		url_, ok := c.GetPostForm("url")
		if !ok {
			c.AbortWithError(500, fmt.Errorf("you have to input 'conn' params"))
			return
		}
		url = url_
	}

	w := worker.NewWorker(period, conn, url)
	cli := client_.(iface.Client)
	cli.SetWorker(w)
}
