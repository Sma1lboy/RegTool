package source

import (
	"encoding/json"
	"os/exec"
	"regtool/common/alias"
	"regtool/console"
	"regtool/source/structs"
)

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

func ChangeAllRegistry(region string) bool {
	rs, err := GetRemoteRegistrySources()
	if err != nil {
		console.Error("Failed to fetch remote sources:", err.Error())
		return false
	}

	// Init source manager

	for _, manager := range registryManagers {
		_, _ = manager.SetRegistry(structs.StringToRegion(region), rs)
	}

	return true
}

var registryManagers = map[string]AppManager{}

// RegisterManager registers a manager for the given names
func RegisterManager(names []string, manager AppManager) {
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

	primaryApp := alias.GetPrimary(app)
	aliases := alias.GetAllAliases(primaryApp)

	if registryManager, ok := registryManagers[primaryApp]; ok {
		_, _ = registryManager.SetRegistry(structs.StringToRegion(region), rs)
	} else {
		return &exec.Error{Name: "Key does not exist", Err: nil}
	}

	for _, alias := range aliases {
		if registryManager, ok := registryManagers[alias]; ok {
			_, _ = registryManager.SetRegistry(structs.StringToRegion(region), rs)
		} else {
			return &exec.Error{Name: "Key does not exist", Err: nil}
		}
	}
	return nil
}

// Get All Registered
func GetAllExistLocalApp() map[string]AppManager {

	res := make(map[string]AppManager)
	for _, appName := range alias.GetAllPrimary() {
		if manager, ok := registryManagers[appName]; ok {
			if !manager.IsExists() {
				continue
			}
			res[appName] = manager
		}
	}
	return res
}
