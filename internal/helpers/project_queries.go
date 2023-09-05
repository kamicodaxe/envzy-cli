package helpers

import (
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

// DeleteProjectByID deletes a project by its ID.
func DeleteProjectByID(projectID uint) error {
	db := app.GetDB()
	return db.Delete(&models.Project{}, projectID).Error
}
