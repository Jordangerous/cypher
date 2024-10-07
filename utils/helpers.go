package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetFileExtension returns the file extension for a given file path
func GetFileExtension(filePath string) string {
	return filepath.Ext(filePath)
}

// CreateDirIfNotExists creates a directory if it does not already exist
func CreateDirIfNotExists(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
		}
	}
	return nil
}
