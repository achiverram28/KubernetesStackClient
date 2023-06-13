package helmpkg

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/gofrs/flock"
	"github.com/pkg/errors"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/downloader"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)



func RepoAdd(name string, url string){
	settings = cli.New()
    repoFile := settings.RepositoryConfig

	err := os.MkdirAll(filepath.Dir(repoFile), os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	fileLock := flock.New(strings.Replace(repoFile, filepath.Ext(repoFile), ".lock", 1))
	lockContext, cancel := context.WithTimeout(context.Background(),30*time.Second)
	defer cancel()
	locked, err := fileLock.TryLockContext(lockContext, time.Second)
	if err == nil && locked {
		defer fileLock.Unlock()
	}
	if err != nil {
		log.Fatal(err)
	}

	k, err := ioutil.ReadFile(repoFile)
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}

	var f repo.File
	if err := yaml.Unmarshal(k, &f); err != nil{
		log.Fatal(err)
	}

	if f.Has(name){
		log.Println("Repo already exists")
		return
	}

	ent := repo.Entry{
		Name: name,
		URL: url,
	}

	r, err := repo.NewChartRepository(&ent,getter.All(settings))
	if err != nil {
		log.Fatal(err)
	}

	if _, err := r.DownloadIndexFile(); err != nil {
		err := errors.Wrapf(err, "looks like %s is not a valid chart repository or cannot be reached", url)
		log.Fatal(err)
	}

	f.Update(&ent)

	if err := f.WriteFile(repoFile, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s has been added to your repositories\n", name)
}

func RepoUpdate() {
	settings = cli.New()

	repoConfig := settings.RepositoryConfig
	repoFile,err := repo.LoadFile(repoConfig)
	if os.IsNotExist(errors.Cause(err)) || len(repoFile.Repositories) == 0{
		log.Fatal(errors.New("No repositories found. You must add one before updating."))

	}

	var repos []*repo.ChartRepository
	for _, rep := range repoFile.Repositories {
		r, err := repo.NewChartRepository(rep, getter.All(settings))
		
		if err != nil {
			log.Fatal(err)
		}
		repos = append(repos, r)

	}

	log.Printf("Hang tight while we grab the latest from your chart repositories...")

	var wg sync.WaitGroup

	for _, re := range repos {
		wg.Add(1)
		go func(re *repo.ChartRepository) {
			defer wg.Done()
			if _, err := re.DownloadIndexFile(); err != nil {
				log.Printf("...Unable to get an update from the %q chart repository (%s):\n\t%s", re.Config.Name, re.Config.URL, err)
			} else {
				log.Printf("...Successfully got an update from the %q chart repository", re.Config.Name)
			}
		}(re)
	}
	wg.Wait()
	log.Printf("Update Complete. ⎈Happy Helming!⎈\n")
}

func InstallChart(name string,repo string,chart string,namespace string){

	settings = cli.New()

	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), namespace, "secret",log.Printf); err != nil {
		log.Fatal(err)
	}

	client := action.NewInstall(actionConfig)
	client.ReleaseName = name
	client.Namespace = namespace

	chartPath, err := client.ChartPathOptions.LocateChart(fmt.Sprintf("%s/%s",repo,chart),settings)
	if err != nil {
		log.Fatal(err)
	}

	chartRequested, err := loader.Load(chartPath)
	if err != nil {
		log.Fatal(err)
	}
	
	validInstallableChart, err := isChartInstallable(chartRequested)
	if !validInstallableChart {
		log.Fatal(err)
	}

	if req:= chartRequested.Metadata.Dependencies; req != nil{
		if err := action.CheckDependencies(chartRequested, req); err!= nil {
			if client.DependencyUpdate {
				 man := &downloader.Manager{
					 Out:			  os.Stdout,
					 ChartPath:		  chartPath,
					 Keyring:		  client.ChartPathOptions.Keyring,
					 SkipUpdate:	  false,
					 Getters:		  getter.All(settings),
					 RepositoryConfig:settings.RepositoryConfig,
					 RepositoryCache:  settings.RepositoryCache,
				 }
				 if err := man.Update(); err != nil {
					log.Fatal(err)
				 }
			} else{
				log.Fatal(err)
			 }
		}
	}


	release, err := client.Run(chartRequested, nil)
	if err != nil {
		log.Fatal(err)
	}

	
	fmt.Println(release.Manifest)
	fmt.Println(release.Name)
}

func isChartInstallable(ch *chart.Chart) (bool, error) {
	switch ch.Metadata.Type {
	case "", "application":
		return true, nil
	}
	return false, errors.Errorf("%s charts are not installable", ch.Metadata.Type)
}