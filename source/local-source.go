package source

import (
	"fmt"
	"regtool/source/localdata"
)

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
	sources, err := localdata.ReadBackupFile()
	if err != nil {
		return nil, fmt.Errorf("failed to read backup file: %v", err)
	}
	return convertLocalSources(sources), nil
}
