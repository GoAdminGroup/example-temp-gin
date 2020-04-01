package middleware

import (
	"github.com/GoAdminGroup/example-temp-gin/pkg/zlog"
	"github.com/gin-gonic/gin"
	"time"
)

// just use logger to record
// filter at status code
// less than 400 use debug
// less than 500 use Info
// other use Warning
// use as
//	g.Use(middleware.LoggerMiddleware())
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// start time
		startTime := time.Now()
		// to next
		c.Next()
		// end time
		endTime := time.Now()

		// latency time
		latencyTime := endTime.Sub(startTime)

		// request IP
		clientIP := c.ClientIP()

		// Method
		reqMethod := c.Request.Method

		// request router
		reqUri := c.Request.RequestURI

		// status code
		statusCode := c.Writer.Status()
		if statusCode < 400 {
			zlog.S().Debugf(
				"=> %15s %13v | %s < %3d -> %s",
				clientIP,
				latencyTime,
				reqMethod,
				statusCode,
				reqUri)
		} else if statusCode < 500 {
			zlog.S().Infof(
				"=> %15s %13v | %s < %3d -> %s",
				clientIP,
				latencyTime,
				reqMethod,
				statusCode,
				reqUri)
		} else {
			zlog.S().Warnf(
				"=> %15s %13v | %s < %3d -> %s",
				clientIP,
				latencyTime,
				reqMethod,
				statusCode,
				reqUri)
		}
	}
}

// logger to mongo
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// logger to MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
