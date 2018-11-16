package main

import (
	"github.com/hashicorp/vault/api"
	"github.com/kubicorn/kubicorn/pkg/logger"
	"os"
	"fmt"
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
	client.SetToken("59ca3b88-b142-c39a-059c-34d999c13a36")
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

}
