package main

import (
	"HangmanWeb/requetes"
	"fmt"
	"net/http"
)

const port = ":1608"

// Fonction main pour démarrer le serveur
func main() {
	// Servir le dossier statique
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./hangman-Web/static"))))

	// Routes du serveur
	http.HandleFunc("/", requetes.AccueilHandler)
	http.HandleFunc("/jeu", requetes.JeuHandler)
	http.HandleFunc("/perdu", requetes.PerduHandler)       // Page de défaite
	http.HandleFunc("/victoire", requetes.VictoireHandler) // Page de victoire
	http.HandleFunc("/relancer", requetes.RelancerHandler) // Relancer la partie

	fmt.Println("Serveur démarré sur http://localhost:1608")
	http.ListenAndServe(port, nil)
}
