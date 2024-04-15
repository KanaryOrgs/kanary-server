package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func LoggerConfig() gin.HandlerFunc {
	location, _ := time.LoadLocation("Asia/Seoul")

	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		timeStamp := param.TimeStamp.In(location).Format(time.RFC3339)
		errorMessage := ""
		if param.ErrorMessage != "" {
			errorMessage = fmt.Sprintf(" Error: %s", param.ErrorMessage)
		}
		return fmt.Sprintf("%s - [%s] \"%s %s\" %d %s %d%s\n",
			param.ClientIP,
			timeStamp,
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
			errorMessage,
		)
	})
}
