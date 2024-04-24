package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func LoggerFormatter(param gin.LogFormatterParams) string {
	return fmt.Sprintf("[kanary] - %s - [%s] \"%s %s %s %d %s\" %s\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.ErrorMessage,
	)
}
