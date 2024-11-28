package middlewares

import (
	"CMS/config"
	"fmt"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func SentryMiddleware() gin.HandlerFunc {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.ApplicationConfig.Sentry.Dns,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
	return sentrygin.New(sentrygin.Options{})
}
