package dotenv

import (
	"fmt"
	"os"
)

func WriteToEnvFile(filename string, key, value string, comment string) error {
	// Delete the file if it exists
	if _, err := os.Stat(filename); err == nil {
		err := os.Remove(filename)
		if err != nil {
			return err
		}
	}

	// Open the .env file in append mode
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the key-value pair to the .env file
	_, err = fmt.Fprintf(file, "%s%s=%s\n", comment, key, value)
	if err != nil {
		return err
	}

	return nil
}
