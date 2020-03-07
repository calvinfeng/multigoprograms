package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configurationName string

func initViperConfig() {
	logrus.Infof("initializing viper with %s configurations", configurationName)
	viper.AddConfigPath("./conf")
	viper.SetConfigName(configurationName)
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		logrus.WithError(err).Fatalf("failed to load configuration from /conf/%s", configurationName)
	}
}

func Execute() error {
	cobra.OnInitialize(initViperConfig)

	rootCmd := &cobra.Command{
		Use: "multigoprograms",
		Short: "A program that has multiple programs",
	}

	rootCmd.PersistentFlags().
		StringVarP(&configurationName, "conf", "c", "development", "configuration file name" )

	runServerCmd := &cobra.Command{
		Use:                        "server",
		Short:                      "Run a HTTP server",
		PreRunE:                    nil,
		RunE:                       runServer,
	}

	runClientCmd := &cobra.Command{
		Use:                        "client",
		Short:                      "Run a HTTP client",
		PreRunE:                    nil,
		RunE:                       runServer,
	}

	rootCmd.AddCommand(runServerCmd, runClientCmd)

	return rootCmd.Execute()
}

