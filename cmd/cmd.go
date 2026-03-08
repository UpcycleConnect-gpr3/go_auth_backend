package cmd

import (
	"authentication_backend/cmd/server"
	"flag"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3/log"
)

func Cmd() {
	startCmd := flag.NewFlagSet("start", flag.ExitOnError)
	serveCmd := flag.NewFlagSet("server", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Commande manquante. Utilisation : monexecutable [start|serve]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "serve":
		serveCmd.Parse(os.Args[2:])
		if err := server.Start("dev"); err != nil {
			log.Fatalf("Erreur lors du démarrage du serveur : %v", err)
		}
	case "start":
		startCmd.Parse(os.Args[2:])
		fmt.Println("Commande 'start' exécutée")

	default:
		fmt.Println("Commande inconnue. Utilisation : go main [start|serve]")
		os.Exit(1)
	}
}
