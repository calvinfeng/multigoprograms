package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func runServer(_ *cobra.Command, _ []string) error {
	logrus.Info("running server")
	return nil
}

func runClient() error {
	logrus.Info("running client")
	return nil
}
