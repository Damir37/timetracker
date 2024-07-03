package service

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"timertracker/internel/config"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	Valider *validator.Validate
}

func NewServiceContext(cfg config.Config, db *gorm.DB, valider *validator.Validate) *ServiceContext {
	return &ServiceContext{
		Config:  cfg,
		DB:      db,
		Valider: valider,
	}
}
