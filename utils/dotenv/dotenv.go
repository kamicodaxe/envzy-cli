package dotenv

import (
	"fmt"
	"os"

	"github.com/kamicodaxe/envzy-cli/internal/app"
	"github.com/kamicodaxe/envzy-cli/internal/models"
)

// Variable represents a dotenv variable and its associated comment.
type Variable struct {
	Key     string
	Value   string
	Comment string
}

// ReadDotenvFile reads the content of a dotenv file.
func ReadDotenvFile(dotenvFilePath string) (string, error) {
	content, err := os.ReadFile(dotenvFilePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// StoreVariablesInDatabase stores the extracted variables in the secrets table.
func StoreVariablesInDatabase(variables []Variable, project models.Project, branch models.Branch) {
	db := app.GetDB()
	// Iterate through the extracted variables
	for _, variable := range variables {
		// Check if the secret already exists in the database
		var existingSecret models.Secret
		result := db.Where("project_id = ? AND branch_id = ? AND name = ?", project.ID, branch.ID, variable.Key).First(&existingSecret)
		if result.Error != nil {
			// Secret doesn't exist, so create a new one
			newSecret := models.Secret{
				ProjectID:   project.ID,
				BranchID:    branch.ID,
				Name:        variable.Key,
				Value:       variable.Value,
				Description: variable.Comment,
			}
			db.Create(&newSecret)
			fmt.Printf("Added new secret: Key=%s, Value=%s, Comment=%s\n", variable.Key, variable.Value, variable.Comment)
		} else {
			// Secret already exists, update it
			existingSecret.Value = variable.Value
			existingSecret.Description = variable.Comment
			db.Save(&existingSecret)
			fmt.Printf("Updated secret: Key=%s, Value=%s, Comment=%s\n", variable.Key, variable.Value, variable.Comment)
		}
	}
}

// AddDotenvFileToSecrets orchestrates the process.
func AddDotenvFileToSecrets(dotenvFilePath string, project models.Project, branch models.Branch) error {
	// Step 1: Read the dotenv file
	content, err := ReadDotenvFile(dotenvFilePath)
	if err != nil {
		return err
	}

	// Step 2: Parse the dotenv file
	variables := ParseDotenvFile(content) // From env_parser.go

	// Step 3: Store variables in the secrets table
	StoreVariablesInDatabase(variables, project, branch)

	return nil
}
