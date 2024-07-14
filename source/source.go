package source

import (
	"encoding/json"
	"os/exec"
	"registryhub/common/alias"
	"registryhub/console"
	"registryhub/source/structs"
	"strings"
)

var aliasManager = alias.NewAliasManager()

// GetRemoteRegistrySources fetches the remote sources and returns them
func GetRemoteRegistrySources() (*structs.RegistrySources, error) {
	cmd := exec.Command("curl", "-L", "https://gitee.com/Sma1lboyyy/registry-hub/raw/main/sources.json")
	output, err := cmd.Output()
	if err != nil {
		console.Error("Failed to fetch remote sources:", err.Error())
		return &structs.RegistrySources{}, err
	}
	var sources structs.RegistrySources
	err = json.Unmarshal(output, &sources)
	return &sources, err
}

// ConvertSources converts sources to a map of package managers to sources
func ConvertSources(sources *structs.RegistrySources) map[string]Source {
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

var SOURCES map[string]Source

func GetRemoteSourcesMap() (map[string]Source, error) {
	sources, err := GetRemoteRegistrySources()
	if err != nil {
		return nil, err
	}
	SOURCES = ConvertSources(sources)
	return SOURCES, nil
}

type Source struct {
	Region string
	Url    string
	Name   string
}

// PrintSources prints all registered sources
func PrintSources(m map[string]Source) {
	printSourceTitle("====================Current Sources=====================")
	for _, v := range m {
		printSource(v.Name, v.Url, v.Region)
	}
}

func printRegionSources(sources structs.RegistryRegionSources, region string) {
	for k, v := range sources {
		// TODO: only concern first url for now
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

func printSuccessMessage(app string, registry string, region string) {
	console.Printf(console.Color.Reset, "Setting ")
	console.Printf(console.Color.Red, "%s", app)
	console.Print(console.Color.Reset, " to")
	console.Printf(console.Color.Green, " %s ", strings.ToUpper(region))
	console.Print(console.Color.Reset, "registry")
	console.Printf(console.Color.Green, " %s", registry)
	console.Print(console.Color.Reset, " ✅\n")
}

func printChangeRegistryHeader(region string) {
	console.Println(console.Color.Purple, "=========Changing all sources to the", strings.ToUpper(region), "registry=========")
	console.Println("", "")
}

func printChangeRegistryFooter() {
	console.Println("", "")
	console.Println(console.Color.Purple, "=================================================================")
}

func ChangeAllRegistry(region string) bool {
	printChangeRegistryHeader(region)
	rs, err := GetRemoteRegistrySources()
	if err != nil {
		console.Error("Failed to fetch remote sources:", err.Error())
		return false
	}

	// Init source manager

	for app, manager := range registryManagers {
		registry, _ := manager.SetRegistry(structs.StringToRegion(region), rs)
		printSuccessMessage(app, registry, region)
	}

	printChangeRegistryFooter()
	return true
}

var registryManagers = map[string]RegistryManager{}

// RegisterManager registers a manager for the given names
func RegisterManager(names []string, manager RegistryManager) {
	for _, name := range names {
		registryManagers[name] = manager
	}
}

func UpdateRegistry(region string, app string) error {
	rs, err := GetRemoteRegistrySources()
	if err != nil {
		console.Error("Failed to fetch remote sources:", err.Error())
		return &exec.Error{Name: "Failed to fetch remote sources", Err: err}
	}

	primaryApp := aliasManager.GetPrimary(app)
	aliases := aliasManager.GetAllAliases(primaryApp)

	for _, alias := range aliases {
		if registryManager, ok := registryManagers[alias]; ok {
			registry, _ := registryManager.SetRegistry(structs.StringToRegion(region), rs)
			printSuccessMessage(alias, registry, region)
		} else {
			return &exec.Error{Name: "Key does not exist", Err: nil}
		}
	}
	return nil
}
