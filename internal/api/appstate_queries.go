package api

import (
	"log"

	"github.com/kamicodaxe/envzy-cli/internal/app"
	"github.com/kamicodaxe/envzy-cli/internal/constants"
	"github.com/kamicodaxe/envzy-cli/internal/models"
)

func GetCurrentProject() *models.Project {
	var selectedProject *models.Project
	var projectName string
	var err error

	kvstore := app.GetKVStore()
	hasSelectedProject := kvstore.HasKey(constants.PROJECT_NAME)
	if hasSelectedProject {
		projectName = kvstore.String(constants.PROJECT_NAME)
	}

	selectedProject, err = GetProjectByName(projectName)
	if err != nil {
		log.Fatal("GetProjectByName error")
		return nil
	}

	return selectedProject
}

func GetCurrentBranch() *models.Branch {
	var selectedBranch *models.Branch
	var branchName string
	var err error

	kvstore := app.GetKVStore()
	hasSelectedBranch := kvstore.HasKey(constants.CURRENT_BRANCH_NAME)
	if hasSelectedBranch {
		branchName = kvstore.String(constants.CURRENT_BRANCH_NAME)
	}

	selectedBranch, err = GetBranchByName(branchName)
	if err != nil {
		return nil
	}

	return selectedBranch
}
