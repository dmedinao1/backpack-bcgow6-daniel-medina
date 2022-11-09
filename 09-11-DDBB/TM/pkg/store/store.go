package store

import (
	"encoding/json"
	"os"
)

type Type string

const (
	FileType Type = "file"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type fileStore struct {
	filePath string
}

func New(storeType Type, filePath string) Store {
	switch storeType {
	case FileType:
		return &fileStore{filePath}
	}

	return nil
}

func (f fileStore) Read(data interface{}) error {
	file, err := os.ReadFile(f.filePath)

	if err != nil {
		return err
	}

	err = json.Unmarshal(file, data)

	if err != nil {
		return err
	}

	return nil
}

func (f fileStore) Write(data interface{}) error {
	rawData, err := json.MarshalIndent(data, "", "	")

	if err != nil {
		return err
	}

	err = os.WriteFile(f.filePath, rawData, 0644)

	if err != nil {
		return err
	}

	return nil
}
