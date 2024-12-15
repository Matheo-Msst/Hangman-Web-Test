package requetes

import (
	"HangmanWeb/hangman-Web/variables"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

// Fonction pour afficher un template
func AfficherTemplate(w http.ResponseWriter, tmpl string, etatJeu *variables.EtatJeu) {
	// Obtenir le chemin absolu du répertoire de travail actuel
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Erreur lors de la récupération du répertoire de travail: %v\n", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Construire le chemin absolu vers le fichier du template
	tmplPath := filepath.Join(cwd, "hangman-Web", "templates", fmt.Sprintf("%s.page.tmpl", tmpl))
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		fmt.Printf("Erreur de chargement du template: %v\n", err)
		http.Error(w, fmt.Sprintf("Erreur lors du chargement du template: %s", err), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, etatJeu)
	if err != nil {
		fmt.Printf("Erreur d'exécution du template: %v\n", err)
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du template: %s", err), http.StatusInternalServerError)
	}
}
