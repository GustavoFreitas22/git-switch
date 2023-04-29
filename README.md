# git-switch
 Esse é um projeto que tem a função de trocar os dados de acesso do arquivo git, conforme os dados informado no arquivo config.env.

 - (Sim, existem algumas formas de fazer isso sem usar uma aplicação em CLI para isso... Mas eu tive problemas com essas outras maneira e então criei a minha.)

## Como instalar

A instalação é bem simples, basta ter o Golang configurado em seu ambiente e seguir os passos:

 - Fazer o clone ou download do projeto
 - Abrir o diretório do projeto e executar `go install git-switch`
 - criar um arquivo `config.env` com o modelo que se encontra no arquivo `configExempli.env`

#### Aviso
Lembrando que nessa versão só é possível trocar entre os perfis `personal` e `work`


