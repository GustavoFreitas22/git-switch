package config

import (
	"flag"
	"os"
	"strings"

	"github.com/GustavoFreitas22/git-switch/model"
)

var Profile string
var rawProfile string
var AddProfile model.Profile

func Init() {
	if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-") {
		flag.StringVar(&Profile, "profile", "", "Define qual profile será utilizado")
		flag.StringVar(&rawProfile, "add", "", "Cria novo profile")
		flag.Parse()

		if rawProfile != "" {
			AddProfile = initInsertProfile()
		}
	}
}

func initInsertProfile() model.Profile {

	var newProfile model.Profile

	data := strings.Split(rawProfile, ";")

	newProfile.Name = data[0]
	newProfile.UserName = data[1]
	newProfile.UserEmail = data[2]

	return newProfile
}
