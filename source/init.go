package source

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Constants for directory and file paths
const (
	DOT_CONFIG_NAME          = ".config"
	REGISTRY_HUB_FOLDER_NAME = "registry-hub"
	SOURCE_BACKUP_FILE_NAME  = "backup.json"
	SOURCE_FILE_NAME         = "sources.json"
	READ_PERMISSION          = 0755
)

var (
	HOME_DIR, _        = os.UserHomeDir()
	DOT_CONFIG_DIR     = filepath.Join(HOME_DIR, DOT_CONFIG_NAME)
	REGISTRY_HUB_DIR   = filepath.Join(DOT_CONFIG_DIR, REGISTRY_HUB_FOLDER_NAME)
	SOURCE_BACKUP_FILE = filepath.Join(REGISTRY_HUB_DIR, SOURCE_BACKUP_FILE_NAME)
	SOURCES_FILE       = filepath.Join(REGISTRY_HUB_DIR, SOURCE_FILE_NAME)
)

// EnsureDirectory checks if the directory exists and creates it if it doesn't
func ensureDirectory(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, READ_PERMISSION)
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
	err := ensureDirectory(DOT_CONFIG_DIR)
	if err != nil {
		fmt.Println("Error ensuring .config directory:", err)
		return
	}

	err = ensureDirectory(REGISTRY_HUB_DIR)
	if err != nil {
		fmt.Println("Error ensuring registry-hub directory:", err)
		return
	}
}

func saveToBackup(source map[string]string) {
	var existingData map[string]string

	// Check if the backup file exists
	if _, err := os.Stat(SOURCE_BACKUP_FILE); err == nil {
		// File exists, read the existing data
		file, err := os.Open(SOURCE_BACKUP_FILE)
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

	// Merge new data into existing data
	for key, value := range source {
		existingData[key] = value
	}

	// Write the merged data back to the backup file
	file, err := os.Create(SOURCE_BACKUP_FILE)
	if err != nil {
		fmt.Println("Error: failed to create backup file", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(existingData)
	if err != nil {
		fmt.Println("Error: failed to encode data to backup file", err)
	}
}

func Run() map[string]string {
	backup()
	sources := map[string]string{
		"npm": getNpmRegistry(),
	}
	saveToBackup(sources)
	return sources
}
