// kube-prometheus-stack installation using helm
package helmpkg

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/repo"
)

var settings *cli.EnvSettings

func HelmRepoList() {

	settings = cli.New()

	repoFile := settings.RepositoryConfig

    err := os.MkdirAll(filepath.Dir(repoFile), os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	
	r, err := repo.LoadFile(repoFile)
	if err != nil {
		log.Fatalf("Failed to load repository file: %v", err)
	}

	for _, re := range r.Repositories {
		fmt.Println(re.Name," ",re.URL)
	}

}