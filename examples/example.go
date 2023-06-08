// This is an example implementation of the pkg package
//everthing is commented here , you can uncomment the respective ones which you want and run it to see the output
// You can take this as a reference to implement the pkg in your projects

package examples

import (
	"fmt"

	"github.coom/k8sjobs/pkg"
)

func examples() {
	// clientset, _ := pkg.ConnectTok8s()
	// metricset, _ := pkg.MetricsConnection()

	// node, nolen, _ := pkg.ListNodes(clientset, metricset)

	// fmt.Println(node, nolen)

	// nam,nlen,_ := pkg.ListNamespaces(clientset)

	// fmt.Println(nam,nlen)

	// ans, len, _ := pkg.ListPods("default", clientset)

	// fmt.Println(ans, len)
	// pkg.CreatePod(clientset, "default", "testpod", "testcontainer", "nginx:latest", 80)

	// pkg.CreateDeployment(clientset, "default", "testdeployment", 1, "app", "test", "testcontainer", "nginx:latest", 80)

	// pkg.CreateService(clientset, "default", "testservice", "app", "test", 80,9000)
	// pkg.PodHealthCheck(clientset, "default", "testpod")
	// pkg.PodMetrics(metricset, "default", "testpod")
}
