package request

import (
	"HangmanWeb/hangman-Web/variables"
	"fmt"
	"html/template"
	"net/http"
)

// Variable globale pour garder l'état du jeu
var GameVar variables.GameState

// SoumettreFormulaire traite les formulaires de lettres et de mots
func SubmitForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Initialiser une nouvelle partie avec 10 vies (ou un nombre que vous préférez)
		GameVar = variables.InitialiserGameState(10)
		// Redirige vers la page de jeu
		http.Redirect(w, r, "/game", http.StatusSeeOther)
	}
}

func RenderTemplate(w http.ResponseWriter, tmplName string) {
	// Chargement du template
	tmplPath := fmt.Sprintf("./hangman-Web/templates/%s.page.tmpl", tmplName)
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors du chargement du template: %s", err), http.StatusInternalServerError)
		return
	}

	// Rendre le template avec l'état du jeu
	err = t.Execute(w, GameVar)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du template: %s", err), http.StatusInternalServerError)
		return
	}
}
