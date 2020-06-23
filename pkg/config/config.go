package config

import (
	"github.com/DimitarPetrov/learn-graphql/internal/storage"
	"github.com/DimitarPetrov/learn-graphql/pkg/server"
)

type Config interface {
	Validate() error
	Init() error
}

type Settings struct {
	Server  server.Settings
	Storage storage.Settings
}

func (s *Settings) Init() error {
	configs := []Config{&s.Server, &s.Storage}
	for _, v := range configs {
		if err := v.Init(); err != nil {
			return err
		}
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func New() (*Settings, error) {
	s := &Settings{}
	if err := s.Init(); err != nil {
		return nil, err
	}
	return s, nil
}
