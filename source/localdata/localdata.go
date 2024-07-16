package localdata

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
	REGISTRY_HUB_FOLDER_NAME = "regtool"
	SOURCE_BACKUP_FILE_NAME  = "backup.json"
	READ_PERMISSION          = 0755
)

var (
	HOME_DIR, _        = os.UserHomeDir()
	DOT_CONFIG_DIR     = filepath.Join(HOME_DIR, DOT_CONFIG_NAME)
	REGISTRY_HUB_DIR   = filepath.Join(DOT_CONFIG_DIR, REGISTRY_HUB_FOLDER_NAME)
	SOURCE_BACKUP_FILE = filepath.Join(REGISTRY_HUB_DIR, SOURCE_BACKUP_FILE_NAME)
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

// ReadBackupFile reads the backup file and returns the data
func ReadBackupFile() (map[string]string, error) {
	var data map[string]string

	// Check if the backup file exists
	if _, err := os.Stat(SOURCE_BACKUP_FILE); err == nil {
		// File exists, read the existing data
		file, err := os.Open(SOURCE_BACKUP_FILE)
		if err != nil {
			return nil, fmt.Errorf("failed to open backup file: %v", err)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&data)
		if err != nil {
			return nil, fmt.Errorf("failed to decode backup file: %v", err)
		}
		return data, nil
	} else if errors.Is(err, os.ErrNotExist) {
		// File does not exist, return empty map
		return make(map[string]string), nil
	} else {
		return nil, fmt.Errorf("failed to check backup file: %v", err)
	}
}

// SaveToBackup saves the provided data to the backup file
func SaveToBackup(source map[string]string) error {
	// Ensure directories exist
	if err := ensureDirectory(DOT_CONFIG_DIR); err != nil {
		return err
	}

	if err := ensureDirectory(REGISTRY_HUB_DIR); err != nil {
		return err
	}

	// Read existing data
	existingData, err := ReadBackupFile()
	if err != nil {
		return err
	}

	// Merge new data into existing data
	for key, value := range source {
		existingData[key] = value
	}

	// Write the merged data back to the backup file
	file, err := os.Create(SOURCE_BACKUP_FILE)
	if err != nil {
		return fmt.Errorf("failed to create backup file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(existingData)
	if err != nil {
		return fmt.Errorf("failed to encode data to backup file: %v", err)
	}

	return nil
}
