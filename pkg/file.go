package pkg

import (
	"errors"
	"os"
	"path/filepath"

	"charm.land/log/v2"
)

func CreateFile(dockerFile DockerFile) error {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Errorf("Erro ao localizar o caminho de execução: %s", err)
		return err
	}

	fullPath := filepath.Join(currentPath, FILE_NAME)

	content, err := CreateContent(dockerFile)
	if err != nil {
		return err
	}

	osFile, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			log.Errorf("O arquivo %s já existe.", FILE_NAME)
			return err
		}

		return err
	}

	defer osFile.Close()

	if _, err := osFile.WriteString(content); err != nil {
		log.Errorf("Erro ao escrever no arquivo: %s", err)
		return err
	}

	log.Info("Arquivo criado com sucesso!")
	return nil
}
