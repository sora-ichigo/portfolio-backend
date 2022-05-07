package config

import (
	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func DSN(env string) (string, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String("ap-northeast-1")},
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := ssm.New(sess)

	key := ""
	switch env {
	case "production":
		key = "/portfolio/dsn"
	case "qa":
		key = "/portfolio/dsn/qa"
	}

	if key != "" {
		output, err := svc.GetParameter(&ssm.GetParameterInput{
			Name:           aws.String(key),
			WithDecryption: aws.Bool(true),
		})
		if err != nil {
			return "", errors.Wrap(err, "failed to ssm GetParameter")
		}

		return *output.Parameter.Value, nil
	}

	return "root:root@tcp(db:3306)/portfolio", nil
}
