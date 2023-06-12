// kube-prometheus-stack installation using helm
package helmpkg

import (
	"log"
	"os"
	"path/filepath"

	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

var settings *cli.EnvSettings

type repoFormat struct {
	Name string
	URL string
}

type chartFormat struct {
	Name string
	chartVersion string
	appVersion string
	Description string
}

func HelmRepoList() []repoFormat{

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

	var repo []repoFormat
	for _, re := range r.Repositories {
		repo = append(repo, repoFormat{Name: re.Name, URL: re.URL})
	}

	return repo

}

func HelmRepoChartList(repoName string) []chartFormat{

	settings = cli.New()

	repoConfig := settings.RepositoryConfig

    err := os.MkdirAll(filepath.Dir(repoConfig), os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	repoFile, err := repo.LoadFile(repoConfig)
	if err != nil {
		log.Fatalf("Failed to load repository file: %v", err)
	}

	var chartNames []chartFormat

	flag := false

	for _, rep := range repoFile.Repositories {
		if rep.Name == repoName {
		flag = true
		r, err := repo.NewChartRepository(rep, getter.All(settings))
		if err != nil {
			log.Fatal(err)
		}

		if _, err := r.DownloadIndexFile(); err != nil {
			log.Fatal(err)
		}

		indexFilePath := r.CachePath+"/"+string(repoName)+"-index.yaml"
		indexFile, err := repo.LoadIndexFile(indexFilePath)
		if err != nil {
			log.Fatal(err)
		}
        
		for name := range indexFile.Entries {
			chartNames = append(chartNames, chartFormat{Name: name, chartVersion: indexFile.Entries[name][0].Version, appVersion: indexFile.Entries[name][0].AppVersion, Description: indexFile.Entries[name][0].Description})
		}

}
}
	if !flag{
		log.Println("Repo does not exist")
		return nil
	}

	return chartNames
	
}