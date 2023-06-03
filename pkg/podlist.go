//Function to list the pods in a particular namespace along with the count

package pkg

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)
 

func ListPods(namespace string,clientset *kubernetes.Clientset) ([]string,int32,error) {

	ans := []string{}

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(),v1.ListOptions{})

	if err != nil {
		log.Panicln("Failed to get pods")
		return nil,int32(-1),err
	}

	for _, pod := range pods.Items {
        ans = append(ans,fmt.Sprintf("%v",pod.Name))
    }

	return ans,int32(len(pods.Items)),nil
}