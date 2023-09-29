package models

import (
	"time"

	"gorm.io/gorm"
)

type Config struct {
	gorm.Model
	Name  string
	Value string
}

// Project represents a project or environment.
type Project struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	Description string
	Branches    []Branch
	CreatedAt   time.Time // createdAt field
	UpdatedAt   time.Time
}

type Branch struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	ProjectID   uint `gorm:"not null"`
	IsDefault   bool `gorm:"default:false"` // Indicates if it's the default branch
	Secrets     []Secret
}

// Secret represents a secret associated with a project.
type Secret struct {
	gorm.Model
	ProjectID   uint   `gorm:"not null"`
	BranchID    uint   `gorm:"not null"`
	Name        string `gorm:"not null"`
	Value       string `gorm:"not null"`
	Comment     string
	Description string
	Branch      Branch
	Project     Project
	History     []SecretHistory
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// SecretHistory represents a historical record of changes to a secret.
type SecretHistory struct {
	gorm.Model
	SecretID  uint   `gorm:"not null"`
	Name      string `gorm:"not null"`
	Value     string `gorm:"not null"`
	Action    string `gorm:"not null"`
	Timestamp time.Time
	Secret    Secret
}

// StagedSecret represents a secret that is staged for commit.
type StagedSecret struct {
	gorm.Model
	ProjectID uint
	Branch    string
	Name      string
	Value     string
}
