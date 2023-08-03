package pkg

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

)


func ListNodes(clientset *kubernetes.Clientset) ([]string, error) {

	nodeNames := []string{}

	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{})
	if err != nil {
		log.Panicln("Failed to get nodes")
		return nil, err
	}

	for _, node := range nodes.Items {
		nodeNames = append(nodeNames, node.Name)

	}

	return nodeNames, nil

}

func GetNodeCapacity(clientset *kubernetes.Clientset , nodeName string) ([]NodeCapacity,error) {

	capacity := []NodeCapacity{}

	node, err := clientset.CoreV1().Nodes().Get(context.TODO(),nodeName,v1.GetOptions{})
	if err != nil {
		log.Panicln("Failed to get the node")
		return nil, err
	}

	cpuCapacity, cpuFound := node.Status.Capacity["cpu"]
	memCapacity, memFound := node.Status.Capacity["memory"]

   if !(cpuFound && memFound){

		log.Panicln("Failed to get the capacity")
		return nil, err
   }

   capacity = append(capacity,NodeCapacity{CPU: cpuCapacity.String(),Memory: memCapacity.String()})
   
   return capacity,nil

}

func GetNodeStatus(clientset *kubernetes.Clientset , nodeName string) ([]NodeStatus,error) {

	statusList := []NodeStatus{}

	node, err := clientset.CoreV1().Nodes().Get(context.TODO(),nodeName,v1.GetOptions{})
	if err != nil {
		log.Panicln("Failed to get the node")
		return nil, err
	}

	for _, condition := range node.Status.Conditions {
			statusList = append(statusList,NodeStatus{Type: fmt.Sprintf("%s",condition.Type),Status: fmt.Sprintf("%s",condition.Status),Reason:fmt.Sprintf("%s",condition.Reason)})
		}

	return statusList,nil


}
