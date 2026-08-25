package pkg

import (
	"fmt"
	"strconv"

	"charm.land/log/v2"
)

type DockerFile struct {
	ServiceName      string
	ContainerName    string
	DataBaseDB       string
	DataBaseUser     string
	DataBasePassword string
	DataBasePort     string
	Volume           string
}

func validateData(
	serviceName,
	containerName,
	dataBaseName,
	dataBaseUser,
	dataBasePassword,
	dataBasePort,
	volume string,
) error {
	if serviceName == "" {
		log.Infof("O nome do serviço não foi definido, nome padrão definido como: %s", DEFAULT_SERVICE_NAME)
		serviceName = DEFAULT_SERVICE_NAME

	} else if len(serviceName) < MIN_SERVICE_NAME {
		return ErrMinsServiceName

	} else if len(serviceName) > MAX_SERVICE_NAME {
		return ErrMaxsServiceName

	}

	if containerName == "" {
		return ErrNoInformedContainerName

	} else if len(containerName) < MIN_CONTAINER_NAME {
		return ErrMinContainerName

	} else if len(containerName) > MAX_CONTAINER_NAME {
		return ErrMaxContainerName

	}

	if dataBaseName == "" {
		return ErrNoInformedDatabaseName

	} else if len(dataBaseName) < MIN_DATABASE_NAME {
		return ErrMinDatabaseName

	} else if len(dataBaseName) > MAX_DATABASE_NAME {
		return ErrMaxDatabaseName

	}

	if dataBaseUser == "" {
		return ErrNoInformedDatabaseUser
	} else if len(dataBaseUser) < MIN_DATABASE_USER {
		return ErrMinDatabaseUser
	} else if len(dataBaseUser) > MAX_DATABASE_USER {
		return ErrMaxDatabaseUser
	}

	port, err := strconv.Atoi(dataBasePort)

	if err != nil {
		return ErrInvalidFormatPort
	} else if port < 1 || port > 65536 {
		return ErrOutRangePort
	} else if len(dataBasePort) < MIN_DATABASE_PORT {
		return ErrMinDatabasePort
	} else if len(dataBasePort) > MAX_DATABASE_PORT {
		return ErrMaxDatabasePort
	}

	if dataBasePassword == "" {
		return ErrNoInformedDatabasePassword
	} else if len(dataBasePassword) < MIN_DATABASE_PASSWORD {
		return ErrMinDatabasePassword
	} else if len(dataBasePassword) > MAX_DATABASE_PASSWORD {
		return ErrMaxDatabasePassword
	}

	if volume == "" {
		return ErrNoInformedVolume
	} else if len(volume) < MIN_VOLUME {
		return ErrMinVolume
	} else if len(volume) > MAX_VOLUME {
		return ErrMaxVolume
	}

	return nil
}

func buildDockerComposeFile(
	serviceName,
	containerName,
	dataBaseName,
	dataBaseUser,
	dataBasePassword,
	dataBasePort,
	volume string,
) (DockerFile, error) {
	serviceName = globalTrimSpace(serviceName)
	containerName = globalTrimSpace(containerName)
	dataBaseName = globalTrimSpace(dataBaseName)
	dataBaseUser = globalTrimSpace(dataBaseUser)
	dataBasePassword = globalTrimSpace(dataBasePassword)
	dataBasePort = globalTrimSpace(dataBasePort)

	if err := validateData(
		serviceName,
		containerName,
		dataBaseName,
		dataBaseUser,
		dataBasePassword,
		dataBasePort,
		volume,
	); err != nil {
		return DockerFile{}, err
	}

	return DockerFile{
		ServiceName:      serviceName,
		ContainerName:    buildContainerName(containerName),
		DataBaseDB:       buildDatabaseName(dataBaseName),
		DataBaseUser:     dataBaseUser,
		DataBasePassword: dataBasePassword,
		DataBasePort:     buildDatabasePort(dataBasePort),
		Volume:           buildVolume(volume),
	}, nil
}

func buildContainerName(s string) string {
	return fmt.Sprintf("%s_postgres", s)
}

func buildDatabaseName(s string) string {
	return fmt.Sprintf("%s_db", s)
}

func buildDatabasePort(s string) string {
	if s == "" {
		s = DEFAULT_DATABASE_PORT
	}

	return fmt.Sprintf("%s:5432", s)
}

func buildVolume(s string) string {
	return fmt.Sprintf("%s_postgres_data", s)
}
