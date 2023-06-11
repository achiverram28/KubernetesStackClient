// Functionality to create service
package pkg

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	kubernetes "k8s.io/client-go/kubernetes"
)

func CreateService(clientset *kubernetes.Clientset, namespace string,
	serviceName string, labelKey string, labelValue string, containerPort int32, targetPort int32) {

	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: serviceName,
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				labelKey: labelValue,
			},
			Ports: []v1.ServicePort{
				{
					Port:       containerPort,
					TargetPort: intstr.FromInt(int(targetPort)),
					Protocol:   v1.ProtocolTCP,
				},
			},
			Type: v1.ServiceTypeLoadBalancer,
		},
	}

	result, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		log.Panicln("Failed to create service")
	}

	fmt.Printf("Successfully created service: %s\n", result.ObjectMeta.Name)
}

func DeleteService(clientset *kubernetes.Clientset, namespace string, serviceName string) {

	err := clientset.CoreV1().Services(namespace).Delete(context.TODO(), serviceName, metav1.DeleteOptions{})
	if err != nil {
		log.Panicln("Failed to delete service")
	}

	log.Printf("The service %s in namespace %s has been deleted", serviceName, namespace)

}
