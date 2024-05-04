package service

import (
	"errors"
	"fmt"
	"os"

	"github.com/GustavoFreitas22/git-switch/model"
)

func SwithProfile(profile model.Profile) error {

	if profile.UserName == "" || profile.UserEmail == "" {
		return errors.New("Dados de perfil inválidos")
	}

	dirName, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(dirName+"/.gitconfig", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	file.WriteString(fmt.Sprintf("[user]\n\t name= %s \n\t email= %s", profile.UserName, profile.UserEmail))

	errToSave := file.Sync()
	if errToSave != nil {
		return errToSave
	}
	return nil
}
