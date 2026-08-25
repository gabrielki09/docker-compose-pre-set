package main

import (
	"fmt"

	"charm.land/log/v2"
	"github.com/gabrielki09/docker-compose-pre-set/pkg"
	"github.com/spf13/cobra"
)

func main() {
	var rootCommand = &cobra.Command{}

	var (
		serviceName      string
		containerName    string
		databaseName     string
		databaseUser     string
		databasePassword string
		databasePort     string
		volume           string
		debug            bool
	)

	var cmd = &cobra.Command{
		Use:          "docker",
		Short:        "Pre set for database docker compose file.",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if debug {
				log.SetLevel(log.DebugLevel)
			}

			if err := pkg.Runner(
				serviceName,
				containerName,
				databaseName,
				databaseUser,
				databasePassword,
				databasePort,
				volume,
			); err != nil {
				return fmt.Errorf("Erro ao criar o arquivo: %s", err)
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&serviceName, "name", "n", "", "Application name")
	cmd.Flags().StringVarP(&containerName, "container", "c", "", "Container name")
	cmd.Flags().StringVarP(&databaseName, "database", "b", "", "Database name")
	cmd.Flags().StringVarP(&databaseUser, "user", "u", "", "Database user")
	cmd.Flags().StringVarP(&databasePassword, "password", "w", "", "Database password")
	cmd.Flags().StringVarP(&databasePort, "port", "p", "", "Database port")
	cmd.Flags().StringVarP(&volume, "volume", "m", "", "Volume container")
	cmd.Flags().BoolVarP(&debug, "debug", "d", false, "Enable debug mode")

	rootCommand.AddCommand(cmd)
	rootCommand.Execute()
}
