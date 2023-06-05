// Functionality to create the deployment
package pkg

import (
	"context"
	"fmt"
	"log"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

func CreateDeployment(clientset *kubernetes.Clientset, namespace string, deploymentName string,replicas int32,labelKey string, labelValue string, containerName string, image string, port int32) {

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: func(i int32) *int32 { return &i }(replicas),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					labelKey: labelValue,
				},
			},
		Template: v1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
                    labelKey: labelValue,
				},
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
                     Name: containerName,
					 Image: image,
					 Ports: []v1.ContainerPort{
						{
						ContainerPort: port,
						},
					 },
					},
				},
			},
		},
},
	}

	result, err := clientset.AppsV1().Deployments(namespace).Create(context.TODO(),deployment,metav1.CreateOptions{})
    if err != nil {
		log.Panicln("Failed to create deployment")
	}

	fmt.Printf("Successfully created deployment: %s\n",result.ObjectMeta.Name)
}

func UpdateDEployment() {
	//TODO
}

func DeleteDeployment() {
	//TODO
}
