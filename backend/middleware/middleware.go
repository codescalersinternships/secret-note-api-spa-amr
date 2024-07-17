package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// RateLimiterConfig holds the configuration for the rate limiter
type RateLimiterConfig struct {
	Rate  rate.Limit // rate of requests per second
	Burst int        // burst size
}

// NewRateLimiter creates a rate limiting middleware
func NewRateLimiter(config RateLimiterConfig) gin.HandlerFunc {
	limiter := rate.NewLimiter(config.Rate, config.Burst)

	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
