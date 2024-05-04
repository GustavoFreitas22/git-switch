package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/GustavoFreitas22/git-switch/config"
	"github.com/GustavoFreitas22/git-switch/repository"
	"github.com/GustavoFreitas22/git-switch/service"
)

func main() {

	config.Init()
	config.InitDatabase()

	if len(os.Args) == 1 {
		fmt.Print("\n\tO Git-switch precisa de argumentos para saber como seguir! Veja nosso Help abaixo:\n")
		helper()
	} else if len(os.Args) > 1 && !strings.HasPrefix(os.Args[1], "-") {
		//Verifica se existem argumentos e se o primeiro argumento na verdade não é uma flag
		switch os.Args[1] {
		case "help":
			helper()
		default:
			log.Fatalf("Opção \"%s\" não existe, digite git-switch help para obter ajuda", os.Args[1])
		}
	} else {
		if config.Profile == "" {
			insertProfile()
		} else {
			changeProfile()
		}
	}
}

func changeProfile() {
	profile := repository.FindProfile(config.Profile)
	fmt.Print("Olha o retorno: ", profile)
	err := service.SwithProfile(profile)
	if err != nil {
		fmt.Println(err)
	}
}

func insertProfile() {
	repository.InsertProfile(config.AddProfile)
}

func helper() {
	fmt.Print("\n\tPara realizar a troca de perfil, basta passar um argumento que esteja salvo no arquivo .env, por exemplo:\n\n\tgit-switch -profile='personal'\n\n")
}
