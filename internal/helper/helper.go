package helper

import (
	"encoding/json"
	"os"
)

func ParseManifestJSON(filename string) (map[string]map[string]any, error) {
	jsonFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var manifest map[string]map[string]any
	err = json.Unmarshal(jsonFile, &manifest)
	if err != nil {
		return nil, err
	}

	return manifest, nil
}
