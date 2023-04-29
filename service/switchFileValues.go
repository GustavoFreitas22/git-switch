package service

import (
	"fmt"
	"log"
	"os"

	"github.com/GustavoFreitas22/git-switch/config"
)

var (
	ProfileValue string
	UserName     string
	Email        string
)

func SwithProfile(profile config.ConfigData) {

	switch profile.Profile {
	case "personal":
		err := editFileToPersonal()
		if err != nil {
			log.Fatal("Erro ao mudar para o profile Work")
		}
	case "work":
		err := editFileToWork()
		if err != nil {
			log.Fatal("Erro ao mudar para o profile Work")
		}
	default:
		log.Fatal("O profile desejado não existe nas configuraçoes")
	}

}

func editFileToWork() error {

	ProfileValue = "WORK"
	valuesToProfile := config.GetEnv(ProfileValue)

	dirName, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	UserName = valuesToProfile[0]
	Email = valuesToProfile[1]

	file, err := os.OpenFile(dirName+"/.gitconfig", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	file.WriteString(fmt.Sprintf("[user]\n\t name= %s \n\t email= %s", UserName, Email))

	errToSave := file.Sync()
	if errToSave != nil {
		return errToSave
	}
	return nil
}
func editFileToPersonal() error {

	ProfileValue = "PERSONAL"
	valuesToProfile := config.GetEnv(ProfileValue)

	dirName, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	UserName = valuesToProfile[0]
	Email = valuesToProfile[1]

	file, err := os.OpenFile(dirName+"/.gitconfig", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	file.WriteString(fmt.Sprintf("[user]\n\t name= %s \n\t email= %s", UserName, Email))

	errToSave := file.Sync()
	if errToSave != nil {
		return errToSave
	}
	return nil
}
