package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/webhippie/medialize/pkg/command"
)

func main() {
	if env := os.Getenv("MEDIALIZE_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}
