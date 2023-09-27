package api

import (
	"github.com/kamicodaxe/envzy-cli/internal/app"
	"github.com/kamicodaxe/envzy-cli/internal/models"
)

// GetBranchByName retrieves a branch by its name.
func GetBranchByName(branchName string) (*models.Branch, error) {
	db := app.GetDB()
	var branch models.Branch
	if err := db.Where("name = ?", branchName).First(&branch).Error; err != nil {
		return nil, err
	}
	return &branch, nil
}

// CreateBranch creates a new branch within a project in the database.
func CreateBranch(branch *models.Branch) error {
	db := app.GetDB()
	return db.Create(branch).Error
}

// GetBranchesByProjectID retrieves all branches within a specific project.
func GetBranchesByProjectID(projectID uint) ([]models.Branch, error) {
	db := app.GetDB()
	var branches []models.Branch
	if err := db.Where("project_id = ?", projectID).Find(&branches).Error; err != nil {
		return nil, err
	}
	return branches, nil
}

// UpdateBranch updates an existing branch in the database.
func UpdateBranch(branch *models.Branch) error {
	db := app.GetDB()
	return db.Save(branch).Error
}

// DeleteBranchByID deletes a branch by its ID.
func DeleteBranchByID(branchID uint) error {
	db := app.GetDB()
	return db.Delete(&models.Branch{}, branchID).Error
}
