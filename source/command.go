package source

import (
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
