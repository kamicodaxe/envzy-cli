// utils/env_writer.go

package utils

import (
	"fmt"
	"os"
)

func writeToEnvFile(filename string, key, value string) error {
	// Open the .env file in append mode
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the key-value pair to the .env file
	_, err = fmt.Fprintf(file, "%s=%s\n", key, value)
	if err != nil {
		return err
	}

	return nil
}
