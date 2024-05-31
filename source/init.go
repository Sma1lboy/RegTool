package source

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const DOT_CONFIG_NAME = ".config"
const REGISTRY_HUB_FOLDER_NAME = "registry-hub"
const SOURCE_BACKUP_FILE_NAME = "backup.json"
const READ_PERMISSION = 0755

// check if the directory exists and create it if it doesn't
func ensureDirectory(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
		fmt.Println("Directory created:", path)
	} else if err != nil {
		return fmt.Errorf("failed to check directory: %v", err)
	}
	return nil
}
func backup() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: error while find user home directory")
	}

	configDir := filepath.Join(homeDir, ".config")

	err = ensureDirectory(configDir)
	if err != nil {
		fmt.Println("Error: err")
	}

	registryHubDir := filepath.Join(configDir, REGISTRY_HUB_FOLDER_NAME)
	err = ensureDirectory(registryHubDir)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func saveToBackup(source map[string]string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: error while find user home directory")
	}

	backupPath := filepath.Join(homeDir, DOT_CONFIG_NAME, REGISTRY_HUB_FOLDER_NAME, SOURCE_BACKUP_FILE_NAME)

	var existingData map[string]string

	// Check if the backup file exists
	if _, err := os.Stat(backupPath); err == nil {
		// File exists, read the existing data
		file, err := os.Open(backupPath)
		if err != nil {
			fmt.Println("Error: failed to open backup file", err)
			return
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&existingData)
		if err != nil {
			fmt.Println("Error: failed to decode backup file", err)
			return
		}
	} else if errors.Is(err, os.ErrNotExist) {
		// File does not exist, initialize empty map
		existingData = make(map[string]string)
	} else {
		fmt.Println("Error: failed to check backup file", err)
		return
	}
}
func Run() {
	backup()
	sources := map[string]string{
		"npm": getNpmRegistry(),
	}
	saveToBackup(sources)

}
