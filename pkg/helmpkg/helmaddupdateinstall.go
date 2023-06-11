package helmpkg

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/gofrs/flock"
	"github.com/pkg/errors"

	"helm.sh/helm/v3/pkg/cli"
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

	fmt.Println(f)

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