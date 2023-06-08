// keep the kubectl proxy running and then run this to get the response body of the nginx using the localhost url to access it

package examples

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func proxynginxurloutput() {

	url := "http://127.0.0.1:<port>/api/v1/namespaces/<namespace-name>/pods/<pod-name>/proxy/" //replace the port, namespace-name and pod-name with the appropriate values , for eg. by default the port will be 8081 as the kube proxy runs on that port, if you defined a different one , you can use that ; namespace-name will be the namespace in which your nginx pod is there; pod-name will be the name of the nginx pod you created

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Failed to get response")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read response body")
	}

	fmt.Println(string(body))

}
