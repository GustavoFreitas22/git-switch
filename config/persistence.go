package config

import (
	"github.com/GustavoFreitas22/git-switch/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Datasource *gorm.DB

func InitDatabase() {
	result, err := gorm.Open(sqlite.Open(".git-profiles.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = result.AutoMigrate(&model.Profile{})
	if err != nil {
		panic("failed to start database")
	}
	Datasource = result
}
