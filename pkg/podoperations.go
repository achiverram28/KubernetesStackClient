//Function to list the pods in a particular namespace along with the count

package pkg

import (
	"context"
	"fmt"
	"log"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

func ListPods(namespace string, clientset *kubernetes.Clientset) ([]string, int32, error) {

	ans := []string{}

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{})

	if err != nil {
		log.Panicln("Failed to get pods")
		return nil, int32(-1), err
	}

	for _, pod := range pods.Items {
		ans = append(ans, fmt.Sprintf("Name: %s, IP: %s", pod.ObjectMeta.Name, pod.Status.PodIP))
	}

	return ans, int32(len(pods.Items)), nil
}

func PodHealthCheck(clientset *kubernetes.Clientset, namespace string, podName string) {

	pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, v1.GetOptions{})
	if err != nil {
		log.Printf("Failed to get pod: %s", podName)
	}

	if pod.Status.Phase != corev1.PodRunning {
		log.Printf("Pod %s is not running", podName)
		return
	}

	podConditions := pod.Status.Conditions

	for _, condition := range podConditions {
		if condition.Type == corev1.PodReady {
			if condition.Status == corev1.ConditionTrue {
				log.Printf("Pod %s is ready", podName)
			} else {
				log.Printf("Pod %s is not ready", podName)
			}
		}
	}

	for _, condition := range podConditions {
		fmt.Println(condition.Type, " : ", condition.Status)
	}

}

func PodMetrics(metricset *metricsv.Clientset, namespace string, podName string) {

	podMetrics, err := metricset.MetricsV1beta1().PodMetricses(namespace).Get(context.TODO(), podName, v1.GetOptions{})
	if err != nil {
		log.Printf("Failed to get metrics for pod %s: %v", podName, err)
		return
	}

	log.Printf("Metrics for the pod %s: ", podName)
	for _, container := range podMetrics.Containers {
		log.Printf("Container name : %s,\n CPU usage: %v,\nMemory usage: %v\n", container.Name, container.Usage["cpu"], container.Usage["memory"])
	}

}
