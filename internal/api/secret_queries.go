package api

import (
	"github.com/kamicodaxe/envzy-cli/internal/app"
	"github.com/kamicodaxe/envzy-cli/internal/models"
	"gorm.io/gorm"
)

// GetSecretByID retrieves a secret by its ID.
func GetSecretByID(db *gorm.DB, secretID uint) (*models.Secret, error) {
	var secret models.Secret
	if err := db.Where("ID = ?", secretID).First(&secret).Error; err != nil {
		return nil, err
	}
	return &secret, nil
}

// GetSecretByName retrieves a secret by its name within a project.
func GetSecretByName(projectID uint, secretName string) (*models.Secret, error) {
	db := app.GetDB()
	var secret models.Secret
	if err := db.Where("project_id = ? AND name = ?", projectID, secretName).First(&secret).Error; err != nil {
		return nil, err
	}
	return &secret, nil
}

// CreateSecret creates a new secret in the database.
func CreateSecret(secret *models.Secret) error {
	db := app.GetDB()
	return db.Create(secret).Error
}

// UpdateSecret updates an existing secret in the database.
func UpdateSecret(secret *models.Secret) error {
	db := app.GetDB()
	return db.Save(secret).Error
}

// DeleteSecretByID deletes a secret by its ID.
func DeleteSecretByID(secretID uint) error {
	db := app.GetDB()
	return db.Delete(&models.Secret{}, secretID).Error
}

// GetSecretsByProjectAndBranch retrieves all secrets within a specific branch of a project.
func GetSecretsByProjectAndBranch(projectID uint, branchID uint) ([]models.Secret, error) {
	db := app.GetDB()
	var secrets []models.Secret

	if err := db.Where("project_id = ? AND branch_id = ?", projectID, branchID).Find(&secrets).Error; err != nil {
		return nil, err
	}

	return secrets, nil
}
