package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

type settings struct {
	Key       string
	Encrypted bool
}

type Store interface {
	Store(apiKey string) error
	Retrieve() (string, error)
}

func NewStore(pwd string) Store {
	return &plainStringStore{}
}

type plainStringStore struct{}

func getConfigFileName() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(usr.HomeDir, ".iftttclient")
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.json"), nil
}

func (p *plainStringStore) Store(apiKey string) error {
	s := &settings{}
	s.Key = apiKey
	fileName, err := getConfigFileName()
	if err != nil {
		return err
	}
	str, err := json.Marshal(s)
	if err != nil {
		return err
	}
	ioutil.WriteFile(fileName, str, 0600)
	return nil
}

func (p *plainStringStore) Retrieve() (string, error) {
	fileName, err := getConfigFileName()
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	s := &settings{}
	err = json.Unmarshal(data, s)
	if err != nil {
		return "", err
	}
	return s.Key, nil
}
