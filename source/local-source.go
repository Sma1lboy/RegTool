package source

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func readBackupFile() (map[string]string, error) {
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

func ReadBackup() map[string]string {
	m, err := readBackupFile()

	if err != nil {
		fmt.Println("Error reading backup file:", err)
		return nil
	}
	return m
}

// Convert map[Name]Source
func convertLocalSources(sources map[string]string) map[string]Source {
	if SOURCES == nil {
		SOURCES, _ = GetRemoteSourcesMap()
		//TODO: handle if networking error
	}

	result := make(map[string]Source)
	for name, url := range sources {
		if SOURCES[url] != (Source{}) {
			result[name] = Source{
				Region: SOURCES[url].Region,
				Url:    SOURCES[url].Url,
				Name:   SOURCES[url].Name,
			}
		} else {
			result[name] = Source{
				Region: "local",
				Url:    url,
				Name:   name,
			}
		}
	}
	return result
}

func GetLocalSourcesMap() (map[string]Source, error) {
	sources := ReadBackup()
	if sources == nil {
		return nil, fmt.Errorf("failed to read backup file")
	}
	return convertLocalSources(sources), nil
}
