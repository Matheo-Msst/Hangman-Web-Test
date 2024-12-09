package main

import (
	"HangmanWeb/request"
	"fmt"
	"net/http"
)

const port = ":1608"
const staticDir = "hangman-Web/static/"

func main() {
	// Servir les fichiers statiques à partir du dossier "static"
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	// Définir les différentes routes pour afficher les templates
	http.HandleFunc("/", request.AcceuilHandler)
	http.HandleFunc("/gameover", request.GameOverHandler)
	http.HandleFunc("/game", request.GameHandler)
	http.HandleFunc("/victory", request.VictoryHandler)

	// Démarrer le serveur HTTP sur le port 1608
	fmt.Println("(http://localhost:1608) - Serveur démarré sur le port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
}
