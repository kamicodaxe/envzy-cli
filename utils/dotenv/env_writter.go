package dotenv

import (
	"fmt"
	"os"

	"github.com/kamicodaxe/envzy-cli/internal/models"
)

func WriteToEnvFile(filename string, secrets []models.Secret) error {
	// Open the .env file in append mode
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// For each secret, call WriteToEnvFile to write them to the env file
	for _, secret := range secrets {
		_, err = fmt.Fprintf(file, "%s%s=%s\n", secret.Comment, secret.Name, secret.Value)
	}

	// Write the key-value pair to the .env file
	if err != nil {
		return err
	}

	return nil
}
