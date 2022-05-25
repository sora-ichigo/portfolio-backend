package sentryset

import (
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

func init() {
	dsn := os.Getenv("SENTRY_DSN")
	appEnv := os.Getenv("APP_ENV")
	err := sentry.Init(sentry.ClientOptions{
		Environment: appEnv,
		Dsn:         dsn,
	})
	if err != nil {
		log.Printf("sentry.Init: %v", err)
	}
}

func CleanUp() {
	defer sentry.Flush(2 * time.Second)
}
