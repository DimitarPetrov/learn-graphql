package migration

import (
	"fmt"
	"github.com/DimitarPetrov/learn-graphql/internal/storage/migration/jobs"
	"github.com/DimitarPetrov/learn-graphql/internal/storage/models"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
	"log"
)

// ServiceAutoMigration migrates all the tables and modifications to the connected source
func ServiceAutoMigration(db *gorm.DB) error {
	// Keep a list of migrations here
	m := gormigrate.New(db, gormigrate.DefaultOptions, nil)
	m.InitSchema(func(db *gorm.DB) error {
		log.Println("[Migration.InitSchema] Initializing database schema")
		switch db.Dialect().GetName() {
		case "postgres":
			// Let's create the UUID extension, the user has to have superuser
			// permission for now
			db.Exec("create extension \"uuid-ossp\";")
		}
		if err := updateMigration(db); err != nil {
			return fmt.Errorf("[Migration.InitSchema]: %v", err)
		}
		// Add more jobs, etc here
		return nil
	})
	if err := m.Migrate(); err != nil {
		return fmt.Errorf("failed to initialize DB schema: %s", err)
	}

	if err := updateMigration(db); err != nil {
		return err
	}
	m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		jobs.SeedUsers,
	})
	return m.Migrate()
}

func updateMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
	).Error
}
