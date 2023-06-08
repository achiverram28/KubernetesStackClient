package pkg

import (
	"log"
	"os"
	"path/filepath"

	clientcmd "k8s.io/client-go/tools/clientcmd"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

func MetricsConnection() (*metricsv.Clientset, error) {
	home, exists := os.LookupEnv("HOME")
	if !exists {
		home = "/root"
	}
	configPath := filepath.Join(home, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		log.Panicln("failed to create k8s config")
		return nil, err
	}

	metricsClientset, err := metricsv.NewForConfig(config)
	if err != nil {
		log.Panicln("Failed to create metrics clientset")
		return nil, err
	}

	return metricsClientset, nil

}
