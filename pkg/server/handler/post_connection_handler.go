package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	llog "github.com/sirupsen/logrus"
	"metric-generator/pkg/server/conn"
	"strconv"
)

func PostConnectionHandler(c *gin.Context) {
	l_, ok := c.Get("logger")
	if ok == false {
		fmt.Println("error--")
		return
	}
	l := l_.(*llog.Logger)

	connStr := c.PostForm("conn")
	connection, err := strconv.ParseInt(connStr, 10, 32)
	if err != nil {
		l.Error(err)
		return
	}
	l.Info("connection parameter is ", connection)
	conn.IncreaseConnectionCount(float64(connection))
	l.Info("complete put connection metrics ")
}
