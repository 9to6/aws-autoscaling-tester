package log

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var Logger *log.Logger

func init() {
	Logger = log.New()
	Logger.Level = log.InfoLevel
	Logger.Formatter = &log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "02/Jan/2006:15:04:05 +0900",
	}
	Logger.Out = os.Stdout
}

func GetLoggerFromContext(c *gin.Context) (*log.Logger, error) {
	logger, ok := c.Get("logger")
	if !ok {
		return nil, errors.New("Gin context has no logger")
	}
	return logger.(*log.Logger), nil
}
