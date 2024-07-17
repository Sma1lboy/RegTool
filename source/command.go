package source

import (
	"regtool/source/localdata"
)

//here is the command implementation of the source

// Check if there is support registry
func Update(updateChan chan string) error {

	managers := GetAllExistLocalApp()
	res := make(map[string]string)
	for name, manager := range managers {
		res[name], _ = manager.GetCurrRegistry()
		updateChan <- name + " is updated"
	}
	localdata.SaveToBackup(res)
	return nil
}
