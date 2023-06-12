package helmpkg

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

type releaseFormat struct {
	Name string
	Namespace string
	ChartName string
	Status string
}
