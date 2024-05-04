package repository

import (
	"github.com/GustavoFreitas22/git-switch/config"
	"github.com/GustavoFreitas22/git-switch/model"
)

func FindProfile(name string) model.Profile {
	var profile model.Profile
	config.Datasource.Where("name = ?", name).Find(&profile)
	return profile
}

func InsertProfile(profile model.Profile) {
	config.Datasource.Create(&profile)
}
