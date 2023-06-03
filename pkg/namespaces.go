package pkg

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)


func ListNamespaces(clientset *kubernetes.Clientset) ([]string,int32,error) {

	ans := []string{}

	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(),v1.ListOptions{})
	if err != nil {
		log.Panicln("Failed to get namespaces")
		return nil,int32(-1),err
	}

	for _, namespace := range namespaces.Items {
		ans = append(ans,fmt.Sprintf("%v",namespace.Name))
	}
    
	return ans,int32(len(namespaces.Items)),nil

}