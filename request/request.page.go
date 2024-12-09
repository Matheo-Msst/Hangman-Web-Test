package request

import (
	"HangmanWeb/hangman-Web/variables"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Déclare une variable globale pour charger les templates
var tmpl *template.Template

func init() {
	templateDir, err := filepath.Abs("./hangman-Web/templates")
	if err != nil {
		log.Fatal("Erreur lors de la résolution du chemin des templates: ", err)
	}

	tmpl, err = template.ParseFiles(
		filepath.Join(templateDir, "Home.page.tmpl"),
		filepath.Join(templateDir, "GameOver.page.tmpl"),
		filepath.Join(templateDir, "Game.page.tmpl"),
		filepath.Join(templateDir, "Victory.page.tmpl"),
	)
	if err != nil {
		log.Fatalf("Erreur lors du chargement des templates : %v", err)
	}
	fmt.Println("Templates chargés avec succès")
}

func AcceuilHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Démarrage d'une nouvelle partie
		fmt.Println("Démarrage d'une nouvelle partie...")
		GameVar = variables.InitialiserGameState(10)
		fmt.Println("Partie initialisée :", GameVar)
		http.Redirect(w, r, "/game", http.StatusSeeOther) // Redirige vers la page de jeu
	} else {
		// Affiche la page d'accueil
		err := tmpl.ExecuteTemplate(w, "Home.page.tmpl", nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du template: %s", err), http.StatusInternalServerError)
		}
	}
}

func GameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Récupérer l'entrée de l'utilisateur
		inputUser := r.FormValue("input-user")
		fmt.Println("Input reçu :", inputUser)

		if len(inputUser) == 0 {
			fmt.Println("Aucune entrée reçue")
		} else {
			fmt.Println("Entrée utilisateur:", inputUser)
			// Appeler la fonction pour mettre à jour le jeu avec la lettre ou le mot
			UpdateGame(w, r)
		}
	}

	// Afficher l'état du jeu après le traitement
	RenderTemplate(w, "Game")
}

// Afficher la page de Game Over
func GameOverHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "GameOver.page.tmpl", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Afficher la page de victoire
func VictoryHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "Victory.page.tmpl", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
