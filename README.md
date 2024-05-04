# git-switch
 Esse é um projeto que tem a função de trocar os dados de acesso do arquivo `.gitconfig`, com base nos profiles adicionados no app.

 - (Sim, existem algumas formas de fazer isso sem usar uma aplicação em CLI para isso... Mas eu tive problemas com essas outras maneira e então criei a minha.)

## Como instalar
 - É preciso ter o go instalado corretamente no seu ambiente
 - Para instalar basta fazer o download do projeto e na raiz executar o comando `go install git-switch.go`, isso ira gerar um executavel compativel com seu ambiente. 

## Como utilizar
### Para adicionar um profile
 - Basta executar o comando da seguinte maneira: `git-switch -add="<NOME_DO_PROFILE>;<GIT_USERNAME>;<GIT_EMAIL>"`

### Para mudar entre os profiles cadastrados
 - Basta executar o comando da seguinte maneira: `git-switch -profile="<NOME_DO_PROFILE>"`





