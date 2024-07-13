package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func SetTestEnv() error {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Errorf("error determining current file path")
	}

	configPath := filepath.Join(filepath.Dir(currentFile), "../../.env")
	err := godotenv.Load(configPath)

	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	err = os.Setenv("ENV", "test")
	if err != nil {
		return fmt.Errorf("error setting ENV variable: %w", err)
	}

	return nil
}
