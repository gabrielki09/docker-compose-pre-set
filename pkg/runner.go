package pkg

// Valida com len(err) > 0
func Runner(
	aplicationName,
	containerName,
	dataBaseDB,
	dataBaseUser,
	dataBasePassword,
	dataBaseport string,
	fullInformed bool,
) map[string]string {
	runnerErrors := make(map[string]string)

	if err := validateData(
		aplicationName,
		containerName,
		dataBaseDB,
		dataBaseUser,
		dataBasePassword,
		dataBaseport,
		fullInformed,
	); err != nil {
		return err
	}

	dockerFile, err := buildDockerComposeFile(
		aplicationName,
		containerName,
		dataBaseDB,
		dataBaseUser,
		dataBasePassword,
		dataBaseport,
		fullInformed,
	)
	if err != nil {
		runnerErrors["build_docker_compose_file"] = err.Error()
		return runnerErrors
	}

	if err := CreateFile(dockerFile); err != nil {
		runnerErrors["create_file"] = err.Error()

		return runnerErrors
	}

	return nil
}
