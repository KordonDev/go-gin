package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s: %s [%s - %v] %s %s \n",
			param.TimeStamp.Format(time.RFC822),
			param.ClientIP,
			param.Method,
			param.StatusCode,
			param.Path,
			param.Latency,
		)
	})
}
