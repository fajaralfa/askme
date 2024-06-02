package helper

import (
	"encoding/json"
	"io"
	"os"
)

func ParseManifestJSON(filename string) map[string]map[string]any {
	jsonFile, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var manifest map[string]map[string]any
	err = json.Unmarshal(jsonFile, &manifest)
	if err != nil {
		panic(err)
	}

	return manifest
}

func ReadJSON(r io.Reader, result any) error {
	err := json.NewDecoder(r).Decode(result)
	return err
}

func WriteJSON(w io.Writer, o any) error {
	err := json.NewEncoder(w).Encode(o)
	return err
}
