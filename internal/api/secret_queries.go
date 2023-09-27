package api

import (
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
func GetSecretByName(db *gorm.DB, projectID uint, secretName string) (*models.Secret, error) {
	var secret models.Secret
	if err := db.Where("project_id = ? AND name = ?", projectID, secretName).First(&secret).Error; err != nil {
		return nil, err
	}
	return &secret, nil
}

// CreateSecret creates a new secret in the database.
func CreateSecret(db *gorm.DB, secret *models.Secret) error {
	return db.Create(secret).Error
}

// UpdateSecret updates an existing secret in the database.
func UpdateSecret(db *gorm.DB, secret *models.Secret) error {
	return db.Save(secret).Error
}

// DeleteSecretByID deletes a secret by its ID.
func DeleteSecretByID(db *gorm.DB, secretID uint) error {
	return db.Delete(&models.Secret{}, secretID).Error
}

// GetAllSecretsByBranch retrieves all secrets within a specific branch of a project.
func GetAllSecretsByBranch(db *gorm.DB, projectID uint, branchName string) ([]models.Secret, error) {
	var secrets []models.Secret

	if err := db.Where("project_id = ? AND branch = ?", projectID, branchName).Find(&secrets).Error; err != nil {
		return nil, err
	}

	return secrets, nil
}
