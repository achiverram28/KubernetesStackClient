// kube-prometheus-stack installation using helm
package helmpkg

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"helm.sh/helm/v3/pkg/repo"
)

type EnvVariable struct {
	Key   string
	Value string
}


func HelmRepoList() {

	cmd := exec.Command("helm", "env")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to run `helm env` command: %v", err)
	}

	envOutput := string(output)

	lines := strings.Split(envOutput, "\n")

	var envVars []EnvVariable
	var repoFile string
	

	for _, line := range lines {
		if line == "" || !strings.Contains(line, "=") {
			continue
		}
		keyValue := strings.SplitN(line, "=", 2)
		key := keyValue[0]
		value := keyValue[1][1 : len(keyValue[1])-1]
		envVar := EnvVariable{
			Key:   key,
			Value: value,
		}
		envVars = append(envVars, envVar)
	}

	for _, envVar := range envVars {
		if  envVar.Key == "HELM_REPOSITORY_CONFIG" {
			repoFile = string(envVar.Value)
		}
	}
	if repoFile == "" {
		log.Fatalf("Failed to find helm repository config file")
	}

	r, err := repo.LoadFile(repoFile)
	if err != nil {
		log.Fatalf("Failed to load repository file: %v", err)
	}

	for _, re := range r.Repositories {
		fmt.Println(re.Name," ",re.URL)
	}

}