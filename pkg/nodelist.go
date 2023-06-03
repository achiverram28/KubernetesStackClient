package pkg

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

func ListNodes(clientset *kubernetes.Clientset)([]string , int32, error){
	
	ans := []string{}

	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(),v1.ListOptions{})
	if err != nil {
		log.Panicln("Failed to get nodes")
		return nil,int32(-1),err
	}
     
	for _, node := range nodes.Items {
		ans = append(ans,fmt.Sprintf("%v",node.Name))
	}

	return ans,int32(len(nodes.Items)),nil

}