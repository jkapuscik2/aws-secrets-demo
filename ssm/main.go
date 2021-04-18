package ssm

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"log"
)

// Provide your region name
var region = "eu-west-1"
var sess *session.Session

func init() {
	s, err := session.NewSession(aws.NewConfig().WithRegion(region))

	if err != nil {
		log.Fatalln(err)
	}
	sess = s
}

func Get(name string) (string, error) {
	svc := ssm.New(sess)

	output, err := svc.GetParameter(
		&ssm.GetParameterInput{
			Name:           aws.String(name),
			WithDecryption: aws.Bool(true),
		},
	)

	if err != nil {
		return "", err
	}

	return *output.Parameter.Value, nil
}
