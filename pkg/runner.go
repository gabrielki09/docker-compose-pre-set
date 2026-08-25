package pkg

func Runner(
	serviceName,
	containerName,
	dataBaseDB,
	dataBaseUser,
	dataBasePassword,
	dataBasePort,
	volume string,
) error {
	dockerFile, err := buildDockerComposeFile(
		serviceName,
		containerName,
		dataBaseDB,
		dataBaseUser,
		dataBasePassword,
		dataBasePort,
		volume,
	)
	if err != nil {
		return err
	}

	if err := CreateFile(dockerFile); err != nil {
		return err
	}

	return nil
}
