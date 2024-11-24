package utils

import (
	"encoding/json"
	"errors"
	"os"
)

func ConfigFromJson[T interface{}](path string, output *T) error {
	// read the config file contents
	config_contents, err := os.ReadFile(path)

	if err != nil {
		return errors.New("could not read file")
	}

	// unmarshal the config into a struct
	err = json.Unmarshal(config_contents, output)
	if err != nil {
		return errors.New("could not unmarshal config")
	}

	return nil
}
