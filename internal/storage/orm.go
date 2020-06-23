package storage

import (
	"fmt"
	"github.com/DimitarPetrov/learn-graphql/internal/storage/migration"
	"github.com/jinzhu/gorm"
	"log"

	//Imports the database dialect of choice
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ORM struct to holds the gorm pointer to db
type ORM struct {
	DB *gorm.DB
}

// Factory creates a db connection with the selected dialect and connection string
func New(settings Settings) (*ORM, error) {
	db, err := gorm.Open(settings.Dialect, settings.ConnectionDSN)
	if err != nil {
		return nil, fmt.Errorf("error opening storage: %s", err)
	}
	orm := &ORM{
		DB: db,
	}
	db.LogMode(settings.LogMode)
	if settings.AutoMigrate {
		err = migration.ServiceAutoMigration(orm.DB)
	}
	log.Println("[ORM] Database connection initialized.")
	return orm, err
}
