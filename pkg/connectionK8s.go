// Function to connect to Kubernetes cluster

package pkg

import (
	"log"
	"os"
	"path/filepath"

	kubernetes "k8s.io/client-go/kubernetes"
	clientcmd "k8s.io/client-go/tools/clientcmd"
)

func ConnectTok8s() *kubernetes.Clientset {
	home, exists := os.LookupEnv("HOME")
	if !exists {
		home = "/root"
	}
	configPath := filepath.Join(home,".kube","config")


	config , err := clientcmd.BuildConfigFromFlags("",configPath)
	if err != nil {
		log.Panicln("failed to create k8s config")
	}

	clientset, err:= kubernetes.NewForConfig(config)
	if err != nil {
		log.Panicln("Failed to create k8s clientset")
	}

	return clientset
}