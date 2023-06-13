package configuration

import (
	"encoding/json"
	"os"
)

type JsonConfig struct {
	Port int
}

func OpenFile(path string) (JsonConfig, error) {
	obj := JsonConfig{Port: 3000}

	_, err := os.Stat(path)
	if err != nil {
		file, err := os.Create(path)
		if err != nil {
			return obj, err
		}
		defer file.Close()

		err = json.NewEncoder(file).Encode(&obj)

		if err != nil {
			return obj, err
		}

		return obj, nil
	}

	data, err := os.ReadFile(path)

	if err != nil {
		return obj, err
	}

	err = json.Unmarshal(data, &obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
