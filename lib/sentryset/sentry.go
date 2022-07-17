package sentryset

import (
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
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
	sentry.Flush(2 * time.Second)
}

type APIGatewayFunc func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func WithCatchErrInAPIGateway(fn APIGatewayFunc) APIGatewayFunc {
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
