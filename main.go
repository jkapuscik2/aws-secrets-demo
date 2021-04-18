package main

import (
	"aws-secrets-demo/sm"
	"aws-secrets-demo/ssm"
	"fmt"
	"log"
)

func connectSSM() {
	pass, err := ssm.Get("/SampleApp/uat/db/postgres/password")
	if err != nil {
		log.Fatalln(err)
	}
	endpoint, err := ssm.Get("/SampleApp/uat/db/postgres/endpoint")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Connecting using Simple System Manager Paramater Store with password %s to %s \n", pass, endpoint)
}

func connectSM() {
	pass, err := sm.Get("/SampleApp/uat/db/postgres/password")
	if err != nil {
		log.Fatalln(err)
	}
	endpoint, err := sm.Get("/SampleApp/uat/db/postgres/endpoint")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Connecting using Secret Manager with password %s to %s \n", pass, endpoint)
}

func main() {
	connectSSM()
	connectSM()
}
