package pkg

import (
	"bytes"
	"text/template"
)

func processTemplate(templateForParse string, structForExecute DockerFile) (string, error) {
	tmpl, err := template.New("compose.yaml").Parse(templateForParse)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer

	if err := tmpl.Execute(&buffer, structForExecute); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func CreateContent(dockerFile DockerFile) (string, error) {
	const defaultTemplate = `
services:
    {{.ServiceName}}:
        image: postgres:18-alpine
        container_name: {{.ContainerName}}
        restart: always
        environment:
            POSTGRES_DB: {{.DataBaseDB}}
            POSTGRES_USER: {{.DataBaseUser}}
            POSTGRES_PASSWORD: {{.DataBasePassword}}
        ports:
            - "{{.DataBasePort}}"
        volumes:
            - {{.Volume}}:/var/lib/postgresql/data

volumes:
    {{.Volume}}:`

	return processTemplate(defaultTemplate, dockerFile)
}
