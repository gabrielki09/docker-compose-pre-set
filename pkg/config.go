package pkg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gabrielki09/docker-compose-pre-set/helpers"
)

type DockerFile struct {
	ContainerName    string
	DataBaseDB       string
	DataBaseUser     string
	DataBasePassword string
	DataBasePort     string
	Volume           string
}

func validateData(
	aplicationName,
	containerName,
	dataBaseName,
	dataBaseUser,
	dataBasePassword,
	dataBaseport string,
	fullInformed bool,
) map[string]string {
	validation := make(map[string]string)

	dataBasePassword = helpers.GlobalTrimSpace(dataBasePassword)

	if helpers.GlobalTrimSpace(aplicationName) == "" {
		validation["application_name"] = ErrNoInformedApplicationName.Error()
		return validation
	}

	if dataBasePassword == "" {
		validation["postgres_password"] = ErrNoInformedDatabasePassword.Error()
	} else if len(dataBasePassword) < MIN_DATABASE_PASSWORD {
		validation["postgres_password"] = ErrMinDatabasePassword.Error()
	} else if len(dataBasePassword) > MAX_DATABASE_PASSWORD {
		validation["postgres_password"] = ErrMaxDatabasePassword.Error()
	}

	if fullInformed {
		containerName = helpers.GlobalTrimSpace(containerName)
		dataBaseName = helpers.GlobalTrimSpace(dataBaseName)
		dataBaseUser = helpers.GlobalTrimSpace(dataBaseUser)
		dataBaseport = helpers.GlobalTrimSpace(dataBaseport)

		if containerName == "" {
			validation["container_name"] = ErrNoInformedContainerName.Error()
		} else if len(containerName) < MIN_CONTAINER_NAME {
			validation["container_name"] = ErrMinContainerName.Error()
		} else if len(containerName) > MAX_CONTAINER_NAME {
			validation["container_name"] = ErrMaxContainerName.Error()
		}

		if dataBaseName == "" {
			validation["postgres_db"] = ErrNoInformedDatabaseName.Error()
		} else if len(dataBaseName) < MIN_DATABASE_NAME {
			validation["postgres_db"] = ErrMinDatabaseName.Error()
		} else if len(dataBaseName) > MAX_DATABASE_NAME {
			validation["postgres_db"] = ErrMaxDatabaseName.Error()
		}

		if dataBaseUser == "" {
			validation["postgres_user"] = ErrNoInformedDatabaseUser.Error()
		} else if len(dataBaseUser) < MIN_DATABASE_USER {
			validation["postgres_user"] = ErrMinDatabaseUser.Error()
		} else if len(dataBaseUser) > MAX_DATABASE_USER {
			validation["postgres_user"] = ErrMaxDatabaseUser.Error()
		}

		if dataBaseport == "" {
			validation["postgres_port"] = ErrNoInformedDatabasePort.Error()
		} else if len(dataBaseport) < MIN_DATABASE_PORT {
			validation["postgres_port"] = ErrMinDatabasePort.Error()
		} else if len(dataBaseport) > MAX_DATABASE_PORT {
			validation["postgres_port"] = ErrMaxDatabasePort.Error()
		} else if _, err := strconv.Atoi(dataBaseport); err != nil {
			validation["postgres_port"] = ErrInvalidFormatPort.Error()
		}

		return validation
	}

	return nil
}

func buildDockerComposeFile(
	aplicationName,
	containerName,
	dataBaseName,
	dataBaseUser,
	dataBasePassword,
	dataBasePort string,
	fullInformed bool,
) (DockerFile, error) {
	containerName = helpers.GlobalTrimSpace(toSnakeCase(containerName))
	dataBaseName = helpers.GlobalTrimSpace(toSnakeCase(dataBaseName))
	dataBaseUser = helpers.GlobalTrimSpace(toSnakeCase(dataBaseUser))
	dataBasePassword = helpers.GlobalTrimSpace(toSnakeCase(dataBasePassword))
	dataBasePort = helpers.GlobalTrimSpace(toSnakeCase(dataBasePort))

	if fullInformed {
		return DockerFile{
			ContainerName:    containerName,
			DataBaseDB:       dataBaseName,
			DataBaseUser:     dataBaseUser,
			DataBasePassword: dataBasePassword,
			DataBasePort:     buildDatabasePort(dataBasePort),
			Volume:           buildVolume(containerName),
		}, nil
	}

	return DockerFile{
		ContainerName:    buildContainerName(aplicationName),
		DataBaseDB:       buildDatabaseName(aplicationName),
		DataBaseUser:     aplicationName,
		DataBasePassword: dataBasePassword,
		DataBasePort:     buildDatabasePort(dataBasePort),
		Volume:           buildVolume(aplicationName),
	}, nil
}

func toSnakeCase(s string) string {
	s = helpers.GlobalTrimSpace(s)
	s = strings.ReplaceAll(s, "-", "_")
	s = strings.ReplaceAll(s, " ", "_")

	return s
}

func buildContainerName(s string) string {
	return fmt.Sprintf("%s_postgres", toSnakeCase(s))
}

func buildDatabaseName(s string) string {
	return fmt.Sprintf("%s_db", toSnakeCase(s))
}

func buildDatabasePort(s string) string {
	return fmt.Sprintf("%s:5432", s)
}

func buildVolume(s string) string {
	return fmt.Sprintf("%s_postgres_data:", s)
}
