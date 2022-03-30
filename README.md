# Golang-Fundamentos-Linguagem
[Udemy - Aprenda Golang do Zero](https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa/)

Curso Aprenda Golang do Zero ! - Módulo Fundamentos da Linguagem

# Tecnologias Utilizadas

- [Go](https://golang.org/) versão 1.16.3

# Executando arquivo

- Para iniciar a compilação do arquivo modulo usar o comando:

`go mod init name-your-module`

- Para criar o arquivo build com a compilação da aplicação :

`go build`

- Executar o seguinte comando via terminal dentro da pasta `1 - Pacotes` para executar a aplicação :

`go run main.go`

- Exemplo de um comando para instalação de dependencias, dentro da pasta `1 - Pacotes`:

`go get github.com/badoux/checkmail`

- Para excluir dependencias nao utilizadas do `go.mod`, usar o comando :

`go mod tidy`

- Para instalar pacotes


    Código exemplo da aula 17 


    go get your-link-package

    go get github.com/urfave/cli


- Executar a app por linha de comando , aula 17 (exemplo):


    Código exemplo da aula 17 


    go run main.go ip --host mercadolivre.com.br

    go run main.go servidores --host mercadolivre.com.br

    ./linha-de-comando ip --host mercadolivre.com.br


# Autor
Jéssika Fernandes 
