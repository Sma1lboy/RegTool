package source

import (
	"fmt"
	"regtool/console"
	"regtool/source/localdata"
	"regtool/source/structs"
)

//here is the command implementation of the source

// Check if there is support registry
func Update(updateChan chan string) error {

	managers := GetAllRegisteredApp()
	res := make(map[string]string)
	for name, manager := range managers {
		res[name], _ = manager.GetCurrRegistry()
		updateChan <- name + " is updated"
	}
	localdata.SaveToBackup(res)
	return nil
}

func ChangeAllRegistry(region string, updateChan chan string) error {
	rs, err := GetRemoteRegistrySources()
	if err != nil {
		console.Error("Failed to fetch remote sources:", err.Error())
		return err
	}

	localAppsMap, err2 := localdata.ReadBackupFile()
	if err2 != nil {
		console.Error("Failed to read backup file:", err2.Error())
		return err2
	}

	appManagers := GetAllRegisteredApp()
	//TODO do backup if changed
	//lets do a git log-like backup for chang every time
	for name, _ := range localAppsMap {
		if manager, ok := appManagers[name]; ok {
			manager.SetRegistry(structs.StringToRegion(region), rs)

		} else {
			console.Error("Manager not found for:", name)
		}

	}
	return nil
}

func ListAllRegistry(ch chan<- string) {
	rs, err := GetRemoteRegistrySources()
	if err != nil {
		ch <- fmt.Sprintf("ERROR: Failed to get remote registry sources: %s", err.Error())
		return
	}

	res := make(map[string][]Source)
	for region, regionSources := range *rs {
		for appName, urls := range regionSources {
			for _, url := range urls {
				res[appName] = append(res[appName], Source{
					Region: string(region),
					Url:    url,
					Name:   appName,
				})
			}
		}
	}

	for appName, sources := range res {
		ch <- fmt.Sprintf("APP: %s", appName)
		for _, source := range sources {
			ch <- fmt.Sprintf("  REGION: %s, URL: %s", source.Region, source.Url)
		}
		ch <- "\n"
	}
}
