package config

import (
	"os"
	"path/filepath"
)

func SaveAPIKey(key string) error {
	dir, err := configDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(dir, "api_key"))
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(key); err != nil {
		return err
	}

	return nil
}

func LoadAPIKey() (string, error) {
	if key := os.Getenv("KANSAI_API_KEY"); key != "" {
		return key, nil
	}

	dir, err := configDir()
	if err != nil {
		return "", err
	}

	b, err := os.ReadFile(filepath.Join(dir, "api_key"))
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func configDir() (string, error) {
	h, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(h, ".kansai"), nil
}
