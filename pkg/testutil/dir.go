package testutil

import (
	"github.com/joho/godotenv"
	"github.com/otiai10/copy"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func WorkOnTempDir(t *testing.T, cleanup bool) string {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	td, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}

	testDataDir := filepath.Join(wd, "testdata")
	if _, err := os.Stat(testDataDir); err == nil {
		err = copy.Copy(testDataDir, filepath.Join(td, "testdata"))
		if err != nil {
			t.Fatal(err)
		}
	}

	err = os.Chdir(td)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Working on temp dir: %s", td)
	t.Cleanup(func() {
		err := os.Chdir(wd)
		if err != nil {
			t.Fatal(err)
		}
		if cleanup {
			err = os.RemoveAll(td)
			if err != nil {
				t.Fatal(err)
			}
		}
	})
	return td
}

func MainOnTempDir() func() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	td, err := os.MkdirTemp("", "test")
	if err != nil {
		panic(err)
	}

	testDataDir := filepath.Join(wd, "testdata")
	if _, err := os.Stat(testDataDir); err == nil {
		err = copy.Copy(testDataDir, filepath.Join(td, "testdata"))
		if err != nil {
			panic(err)
		}
	}

	err = os.Chdir(td)
	if err != nil {
		panic(err)
	}
	log.Printf("Working on temp dir: %s", td)
	return func() {
		err := os.Chdir(wd)
		if err != nil {
			panic(err)
		}
		err = os.RemoveAll(td)
		if err != nil {
			panic(err)
		}
	}
}

func LoadTestEnv() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	projectDir := cwd
	for {
		envPath := filepath.Join(projectDir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			err = godotenv.Load(envPath)
			if err != nil {
				panic(err)
			}
		}
		testDataEnvPath := filepath.Join(projectDir, "testdata", ".env")
		if _, err := os.Stat(testDataEnvPath); err == nil {
			err = godotenv.Load(testDataEnvPath)
			if err != nil {
				panic(err)
			}
		}

		parentDir := filepath.Dir(projectDir)
		if parentDir == projectDir {
			break
		}
		projectDir = parentDir
	}
}
