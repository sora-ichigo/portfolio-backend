package sentryset

import (
	"log"
	"os"
	"portfolio-backend/lib"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/getsentry/sentry-go"
)

func init() {
	dsn := os.Getenv("SENTRY_DSN")
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		log.Println("APP_ENV is not set")
	} else if dsn == "" {
		log.Fatalf("SENTRY_DSN is not set")
	}
	err := sentry.Init(sentry.ClientOptions{
		Environment:      appEnv,
		Dsn:              dsn,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Printf("sentry.Init: %v", err)
	}

	sentry.CaptureMessage("sentry.Init success!")
}

func CleanUp() {
	sentry.Flush(2 * time.Second)
}

func WithCatchErr(fn lib.APIGatewayFunc) lib.APIGatewayFunc {
	return func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		defer func() {
			r := recover()
			if r != nil {
				sentry.CaptureException(r.(error))
			}
		}()

		res, err := fn(request)
		if err != nil {
			sentry.CaptureException(err)
		}

		return res, nil
	}
}
