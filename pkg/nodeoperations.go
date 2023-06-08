package pkg

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

func ListNodes(clientset *kubernetes.Clientset, metricset *metricsv.Clientset) ([]string, int32, error) {

	ans := []string{}

	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{})
	if err != nil {
		log.Panicln("Failed to get nodes")
		return nil, int32(-1), err
	}

	for _, node := range nodes.Items {
		ans = append(ans, fmt.Sprintf("%v", node.Name))
		nodeMetrics, err := metricset.MetricsV1beta1().NodeMetricses().List(context.TODO(), v1.ListOptions{})
		if err != nil {
			log.Printf("Failed to get metrics for node %s: %v", node.Name, err)
			continue
		}

		log.Printf("Metrics for the node %s: ", node.Name)
		for _, metric := range nodeMetrics.Items {
			log.Printf("CPU usage: %v, Memory usage: %v", metric.Usage["cpu"], metric.Usage["memory"])
		}

	}

	return ans, int32(len(nodes.Items)), nil

}
