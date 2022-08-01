package authset

import (
	"os"
	"portfolio-backend/lib"

	"github.com/aws/aws-lambda-go/events"
)

func WithApiKeyAuth(fn lib.APIGatewayFunc) lib.APIGatewayFunc {
	return func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		apiKey := os.Getenv("API_KEY")
		s := request.Headers["x-portfolio-backend-api-key"]
		if s == "" || s != apiKey {
			return events.APIGatewayProxyResponse{
				StatusCode: 401,
				Body:       "Unauthorized",
			}, nil
		}

		res, _ := fn(request)
		return res, nil
	}
}
