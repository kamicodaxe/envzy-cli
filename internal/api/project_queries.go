package api

import (
	"errors"
	"fmt"
	"log"

	"github.com/kamicodaxe/envzy-cli/internal/app"
	"github.com/kamicodaxe/envzy-cli/internal/models"
)

// GetProjectByName retrieves a project by its name.
func GetProjectByName(projectName string) (*models.Project, error) {
	db := app.GetDB()
	var project models.Project
	if err := db.Where("name = ?", projectName).First(&project).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

// CreateProject creates a new project in the database.
func CreateProject(project *models.Project) error {
	db := app.GetDB()
	return db.Create(project).Error
}

// GetProjects retrieves all projects from the database.
func GetProjects() ([]models.Project, error) {
	db := app.GetDB()
	var projects []models.Project
	if err := db.Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

// UpdateProject updates an existing project in the database.
func UpdateProject(project *models.Project) error {
	db := app.GetDB()
	return db.Save(project).Error
}

// DeleteProject deletes a project and its associated secrets from the database.
func DeleteProjectByID(projectID uint) error {
	db := app.GetDB()
	if db == nil {
		return errors.New("failed to connect to the database")
	}

	// Start a database transaction
	tx := db.Begin()

	// Delete the project
	if err := tx.Delete(&models.Project{}, projectID).Error; err != nil {
		tx.Rollback() // Rollback the transaction on error
		log.Printf("Failed to delete project (ID: %d): %v", projectID, err)
		return fmt.Errorf("failed to delete project: %v", err)
	}

	// Delete the associated secrets
	if err := tx.Where("project_id = ?", projectID).Delete(&models.Secret{}).Error; err != nil {
		tx.Rollback() // Rollback the transaction on error
		log.Printf("Failed to delete associated secrets for project (ID: %d): %v", projectID, err)
		return fmt.Errorf("failed to delete associated secrets: %v", err)
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback() // Rollback the transaction on error
		log.Printf("Failed to commit transaction: %v", err)
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}
