package test

import (
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

func LoadTestEnv() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	projectDir := cwd
	for {
		envPath := filepath.Join(projectDir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			err = godotenv.Load(envPath)
			if err != nil {
				return err
			}
		}

		parentDir := filepath.Dir(projectDir)
		if parentDir == projectDir {
			break
		}
		projectDir = parentDir
	}
	return nil
}
