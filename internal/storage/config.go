package storage

import (
	"fmt"
	"github.com/vrischmann/envconfig"
)

type Settings struct {
	AutoMigrate   bool   `envconfig:"default=true"`
	LogMode       bool   `envconfig:"default=true"`
	SeedDB        bool   `envconfig:"default=true"`
	Dialect       string `envconfig:"default=postgres"`
	ConnectionDSN string `envconfig:"default="`
}

func (s *Settings) Validate() error {
	if len(s.Dialect) == 0 {
		return fmt.Errorf("ORM dialect missing")
	}
	if len(s.ConnectionDSN) == 0 {
		return fmt.Errorf("DB connection DSN missing")
	}
	return nil
}

func (s *Settings) Init() error {
	return envconfig.InitWithPrefix(s, "STORAGE")
}
