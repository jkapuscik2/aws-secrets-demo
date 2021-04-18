package sm

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"log"
)

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
	svc := secretsmanager.New(sess)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(name),
	}
	result, err := svc.GetSecretValue(input)

	if err != nil {
		return "", err
	}

	return *result.SecretString, nil
}
