package source

import (
	"encoding/json"
	"os/exec"
	"registryhub/console"
)

type Region string

const (
	CN Region = "cn"
	US Region = "us"
	EU Region = "eu"
)

// RegistrySources is a map of regions to registry regions
type RegistrySources map[Region]RegistryRegion

// RegistryRegion is a map of package managers to urls
type RegistryRegion map[string][]string

// Run fetches the remote sources and returns them
func GetRemoteRegistrySources() (*RegistrySources, error) {

	cmd := exec.Command("curl", "-L", "https://gitee.com/Sma1lboyyy/registry-hub/raw/main/sources.json")
	output, err := cmd.Output()
	if err != nil {
		console.Error("Failed to fetch remote sources:", err.Error())
		return &RegistrySources{}, err

	}
	var sources RegistrySources
	err = json.Unmarshal(output, &sources)

	return &sources, err
}
func GetRemoteSourcesMap() (map[string]Source, error) {
	sources, err := GetRemoteRegistrySources()
	if err != nil {
		return nil, err
	}
	return ConvertSources(sources), nil
}

type Source struct {
	Region string
	Url    string
	Name   string
}

/*
Convert sources to a map of package managers to sources
*/
func ConvertSources(sources *RegistrySources) map[string]Source {
	result := make(map[string]Source)
	for region, registryRegion := range *sources {
		for packageManager, urls := range registryRegion {
			result[packageManager] = Source{
				Region: string(region),
				Url:    urls[0],
				Name:   packageManager,
			}
		}
	}
	return result
}

/*
Custom print functions with colors
*/

func PrintSources(sources *RegistrySources) {
	m := ConvertSources(sources)
	printSourceTitle("====================Current Sources=====================")
	for _, v := range m {
		printSource(v.Name, v.Url, v.Region)
	}
}
func printRegionSources(sources RegistryRegion, region string) {
	for k, v := range sources {
		//TODO: only concern first url for now
		printSource(k, v[0], region)
	}
}
func printSourceTitle(title string) {
	console.Println(console.Color.Purple, title)
}
func printSource(source string, url string, region string) {
	console.Print(source)
	console.Print(" ")
	console.Print(console.Color.Cyan, url)
	console.Print(" ")
	console.Print(console.Color.Purple, region)
	console.Println("")
}
