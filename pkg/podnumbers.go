package pkg

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

)
 

func NumberOfPods(namespace string) {
	clientset := ConnectTok8s()

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(),v1.ListOptions{})

	if err != nil {
		log.Panicln("Failed to get pods")
	}
    
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
}