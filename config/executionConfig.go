package config

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type ConfigData struct {
	Profile string
}

var Data ConfigData
var profile string

func Init() {
	if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-") {
		flag.StringVar(&profile, "profile", "", "Define qual profile será utilizado")
		flag.Parse()

		initValidation(profile)

		Data = ConfigData{profile}
	}
}

func initValidation(value string) {
	if value == "" {
		fmt.Println("A flag de profile é necessaria")
		os.Exit(500)
	}
}
