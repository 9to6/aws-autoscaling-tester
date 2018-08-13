package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	llog "github.com/sirupsen/logrus"
	"github.com/9to6/aws-autoscaling-tester/pkg/server/conn"
	"strconv"
)

func PostConnectionHandler(c *gin.Context) {
	l_, ok := c.Get("logger")
	if ok == false {
		fmt.Println("error--")
		c.AbortWithStatusJSON(500, gin.H{
			"message": "logger error",
		})
		c.Header("Content-Type", "application/json")
		return
	}
	l := l_.(*llog.Logger)

	connStr := c.PostForm("conn")
	connection, err := strconv.ParseInt(connStr, 10, 32)
	if err != nil {
		l.Error(err)
		c.JSON(500, gin.H{
			"message": "you have to input 'conn' params",
		})
		c.AbortWithError(500, err)
		c.Header("Content-Type", "application/json")
		return
	}
	l.Info("connection parameter is ", connection)
	conn.IncreaseConnectionCount(float64(connection))
	l.Info("complete put connection metrics ")
}
