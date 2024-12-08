package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imniynaiy/ticket-system/internal/log"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Stop timer
		timeStamp := time.Now()
		latency := timeStamp.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		bodySize := c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}

		log.Info("request:", log.String("path", path), log.String("t", timeStamp.Format("2006-01-02 15:04:05.000")), log.String("ip", clientIP),
			log.Int("lat_ms", int(latency.Milliseconds())), log.String("m", method), log.Int("ret", statusCode),
			log.Int("size", bodySize), log.String("err", errorMessage),
		)
	}
}
