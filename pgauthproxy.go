package main

import (
	"github.com/luqasn/pgAuthProxy/cmd"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	if err := cmd.RootCommand(); err != nil {
		log.WithError(err).Fatal("Application start failed")
		os.Exit(1)
	}
}
