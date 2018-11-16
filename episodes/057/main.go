package main

import (
	"github.com/hashicorp/vault/api"
	"github.com/kubicorn/kubicorn/pkg/logger"
	"os"
	"fmt"
	"io/ioutil"
)

func main() {



	config := &api.Config{
		Address: "https://127.0.0.1:8200",
	}
	client, err := api.NewClient(config)
	if err != nil {
		logger.Critical("Unable to create client: %v", err)
		os.Exit(1)
	}


	// ---------------------------------------------------
	// Kubernetes Auth Method
	//
	//
	// [WARNING] Passing this file to another system will give away total access to the service account if compromised
	//
	bytes, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount")
	if err != nil {
		logger.Critical("Unable to read service account: %v", err)
		os.Exit(1)
	}
	//options := map[string]interface{}{
	//	"role": "vault-operator-role",
	//  "jwt": string(bytes),
	//}
	//path := "auth/kubernetes/login"
	//secret, err := client.Logical().Write(path, options)
	//if err != nil {
	//	logger.Critical("Unable to login with Kubernetes auth method: %v", err)
	//	os.Exit(1)
	//}
	//token, err := secret.TokenID()
	//if err != nil {
	//	logger.Critical("Unable to get token with Kubernetes auth method: %v", err)
	//	os.Exit(1)
	//}
	//
	// ---------------------------------------------------



	// ---------------------------------------------------
	// Token
	//
	//
	// Define the token here to authenticate with Vault
	token := ""
	client.SetToken(token)
	//
	// ---------------------------------------------------


	client.Logical().Write("/secret/tgik-secret", map[string]interface{}{
		"value": "tgik is awesome",
		"foo": "bar",
	})

	secretValues, err := client.Logical().Read("/secret/tgik-secret")
	if err != nil {
		logger.Critical("Unable to read secret: %v", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", secretValues)

	client.Auth().Token().

}
