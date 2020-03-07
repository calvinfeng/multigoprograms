package main

import (
	"github.com/calvinfeng/multigoprograms/cmd"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:             false,
		FullTimestamp:             true,
	})
}

func main() {
	if err := cmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
