package cmd

import (
	"authentication_backend/cmd/database"
	"authentication_backend/cmd/server"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3/log"
)

func Cmd() {

	if len(os.Args) < 2 {
		fmt.Println("Commande manquante. Utilisation : monexecutable [start|serve]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "serve":
		if err := server.Start(); err != nil {
			log.Fatalf("Erreur lors du démarrage du serveur : %v", err)
		}

	case "migrate":
		database.Migrate()

	default:
		fmt.Println("Commande inconnue. Utilisation : go main serve")
		os.Exit(1)
	}
}
