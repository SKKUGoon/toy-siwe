package util

import "github.com/gin-gonic/gin"

// Access control
const (
	ALLOW_ORIGIN          = "Access-Control-Allow-Origin"
	ALLOW_CREDENTIALS     = "Access-Control-Allow-Credentials"
	ALLOW_CONTROL_HEADERS = "Access-Control-Allow-Headers"
	ALLOW_CONTROL_METHODS = "Access-Control-Allow-Methods"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set(ALLOW_ORIGIN, "*")
		c.Writer.Header().Set(ALLOW_CREDENTIALS, "true")
		c.Writer.Header().Set(ALLOW_CONTROL_HEADERS, "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set(ALLOW_CONTROL_METHODS, "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
