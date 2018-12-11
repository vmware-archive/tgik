package kinbd

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/kris-nova/charlie/network"
	"github.com/kubicorn/kubicorn/pkg/logger"
	"sigs.k8s.io/kind/pkg/cluster"
	"sigs.k8s.io/kind/pkg/cluster/config"
)

const (
	APISleepSeconds   = 5
	APISocketAttempts = 40
	APIEndpoint       = "localhost"
	APIPort           = "58318"
	ClusterName       = "tgik"
)

func TestMain(m *testing.M) {
	logger.TestMode = true
	logger.Level = 4
	exitCode := 0

	// Set up cluster using KIND
	ctx := cluster.NewContext(ClusterName)
	cfg := &config.Config{
		Image: "kindest/node:v1.12.2",
	}
	retain := false
	wait := time.Duration(0)
	err := ctx.Create(cfg, retain, wait)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Defer tear town cluster
	defer func() {
		ctx.Delete()
	}()

	// Run tests
	status := m.Run()

	if status != 0 {
		fmt.Printf("-----------------------------------------------------------------------\n")
		fmt.Printf("[FAILURE]\n")
		fmt.Printf("-----------------------------------------------------------------------\n")
		exitCode = 1
	}

	os.Exit(exitCode)
}

func TestApiListen(t *testing.T) {
	success := false
	for i := 0; i < APISocketAttempts; i++ {
		_, err := network.AssertTcpSocketAcceptsConnection(fmt.Sprintf("%s:%s", APIEndpoint, APIPort), "opening a new socket connection against the Kubernetes API")
		if err != nil {
			logger.Info("Attempting to open a socket to the Kubernetes API: %v...\n", err)
			time.Sleep(time.Duration(APISleepSeconds) * time.Second)
			continue
		}
		success = true
	}
	if !success {
		t.Fatalf("Unable to connect to Kubernetes API")
	}
}
