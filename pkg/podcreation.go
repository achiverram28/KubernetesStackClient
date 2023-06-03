//Functionality to create pods

package pkg

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

func CreatePod(clientset *kubernetes.Clientset,Namespace string,PodName string,ContainerName string,Image string, Port int32) {

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: PodName,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name: ContainerName,
					Image: Image,
					Ports: []v1.ContainerPort{
						{
							ContainerPort: Port,
						},
					},
				},
			

			},
		},
	}

	createdPod, err := clientset.CoreV1().Pods(Namespace).Create(context.TODO(),pod,metav1.CreateOptions{})
	if err != nil {
		log.Panicln("Failed to create pod")
	}

	fmt.Printf("Successfully created pod: %s\n",createdPod.ObjectMeta.Name)

}