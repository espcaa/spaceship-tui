package main

import (
	"encoding/json"
	"os"

	"github.com/zalando/go-keyring"
)

type Credentials struct {
	APIKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
}

func SaveCredentials(creds Credentials) error {

	err := saveCredentialsToKeyring(creds)
	if err != nil {
		// keyring failed, fallback to disk
		return saveCredentialsToDisk(creds)
	}
	return nil
}

func LoadCredentials() (Credentials, error) {

	creds, err := loadCredentialsFromKeyring()
	if err != nil {
		// keyring failed, fallback to disk
		return loadCredentialsFromDisk()
	}
	return creds, nil
}

func loadCredentialsFromKeyring() (Credentials, error) {

	apiKey, err := keyring.Get("spaceship-tui", "api_key")
	if err != nil {
		return Credentials{}, err
	}

	apiSecret, err := keyring.Get("spaceship-tui", "api_secret")
	if err != nil {
		return Credentials{}, err
	}

	return Credentials{
		APIKey:    apiKey,
		APISecret: apiSecret,
	}, nil
}

func loadCredentialsFromDisk() (Credentials, error) {

	// load from json file
	homedir, err := os.UserHomeDir()
	path := homedir + "/.config/spaceship-tui/secrets.json"
	file, err := os.ReadFile(path)
	if err != nil {
		return Credentials{}, err
	}

	var creds struct {
		APIKey    string `json:"api_key"`
		APISecret string `json:"api_secret"`
	}

	err = json.Unmarshal(file, &creds)
	if err != nil {
		return Credentials{}, err
	}

	return Credentials{
		APIKey:    creds.APIKey,
		APISecret: creds.APISecret,
	}, nil
}

func saveCredentialsToDisk(creds Credentials) error {

	data, err := json.MarshalIndent(creds, "", "  ")
	if err != nil {
		return err
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	path := homedir + "/.config/spaceship-tui/secrets.json"

	err = os.MkdirAll(homedir+"/.config/spaceship-tui", 0700)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0600)
	if err != nil {
		return err
	}

	return nil
}

func saveCredentialsToKeyring(creds Credentials) error {

	err := keyring.Set("spaceship-tui", "api_key", creds.APIKey)
	if err != nil {
		return err
	}

	err = keyring.Set("spaceship-tui", "api_secret", creds.APISecret)
	if err != nil {
		return err
	}

	return nil
}
