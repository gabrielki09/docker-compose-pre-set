package pkg

import (
	"errors"
	"fmt"
)

const (
	FILE_NAME = "compose.yaml"

	MIN_CONTAINER_NAME = 3
	MAX_CONTAINER_NAME = 255

	MIN_DATABASE_NAME = 3
	MAX_DATABASE_NAME = 255

	MIN_DATABASE_USER = 3
	MAX_DATABASE_USER = 255

	MIN_DATABASE_PASSWORD = 1
	MAX_DATABASE_PASSWORD = 255

	MIN_DATABASE_PORT = 4
	MAX_DATABASE_PORT = 16
)

var ErrNoInformedApplicationName = errors.New("O nome da aplicação é obrigatório.")

var ErrNoInformedContainerName = errors.New("o nome do container da aplicação é obrigatório.")
var ErrMinContainerName = fmt.Errorf("o nome do container precisa possuir no mínimo %d caracteres.", MIN_CONTAINER_NAME)
var ErrMaxContainerName = fmt.Errorf("o nome do container precisa possuir no máximo %d caracteres.", MAX_CONTAINER_NAME)

var ErrNoInformedDatabaseName = errors.New("o nome do banco de dados da aplicação é obrigatório.")
var ErrMinDatabaseName = fmt.Errorf("o nome do banco de dados precisa possuir no mínimo %d caracteres.", MIN_DATABASE_NAME)
var ErrMaxDatabaseName = fmt.Errorf("o nome do banco de dados precisa possuir no máximo %d caracteres.", MAX_DATABASE_NAME)

var ErrNoInformedDatabaseUser = errors.New("o usuário do banco é obrigatório.")
var ErrMinDatabaseUser = fmt.Errorf("o usuário do banco de dados precisa possuir no mínimo %d caracteres.", MIN_DATABASE_USER)
var ErrMaxDatabaseUser = fmt.Errorf("o usuário do banco de dados precisa possuir no máximo %d caracteres.", MAX_DATABASE_USER)

var ErrNoInformedDatabasePassword = errors.New("a senha do banco de dados da aplicação é obrigatório.")
var ErrMinDatabasePassword = fmt.Errorf("a senha do banco de dados precisa possuir no mínimo %d caracteres.", MIN_DATABASE_PASSWORD)
var ErrMaxDatabasePassword = fmt.Errorf("a senha do banco de dados precisa possuir no máximo %d caracteres.", MAX_DATABASE_PASSWORD)

var ErrNoInformedDatabasePort = errors.New("a porta do banco de dados da aplicação é obrigatório.")
var ErrMinDatabasePort = fmt.Errorf("a porta do banco de dados precisa possuir no mínimo %d caracteres.", MIN_DATABASE_PORT)
var ErrMaxDatabasePort = fmt.Errorf("a porta do banco de dados precisa possuir no máximo %d caracteres.", MAX_DATABASE_PORT)
var ErrInvalidFormatPort = errors.New("a porta do banco de dados da aplicação precisa ser um número inteiro.")
