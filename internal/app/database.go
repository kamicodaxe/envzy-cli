package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kamicodaxe/envzy-cli/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitializeDatabase() (*gorm.DB, error) {
	// Determine the user-specific configuration directory path
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	// Define the name of your CLI tool's configuration directory
	toolConfigDir := "envzy"

	// Create the full path to the tool's configuration directory
	toolDir := filepath.Join(configDir, toolConfigDir)

	// Ensure that the directory exists; create it if it doesn't
	if _, err := os.Stat(toolDir); os.IsNotExist(err) {
		if err := os.MkdirAll(toolDir, 0700); err != nil {
			return nil, err
		}
	}

	// Define the name of your SQLite database file
	dbName := "envzy.db"

	// Create the full path to the SQLite database file
	dbPath := filepath.Join(toolDir, dbName)

	// Initialize the SQLite database
	sqliteDB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db = sqliteDB

	// Auto-migrate the schema
	db.AutoMigrate(
		&models.Config{},
		&models.Project{},
		&models.Branch{},
		&models.Secret{},
		&models.SecretHistory{},
	)

	fmt.Println("Database path:", dbPath)

	return db, nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return db
}
