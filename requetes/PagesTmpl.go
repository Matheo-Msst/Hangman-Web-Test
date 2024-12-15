package requetes

import (
	"HangmanWeb/hangman-Web/variables"
	"HangmanWeb/hangman-classic/fonctions"

	"net/http"
	"strings"
)

// Handler pour la page d'accueil
func AccueilHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Crée une nouvelle partie et redirige vers la page de jeu
		variables.Etatjeu = variables.NouvellePartie()
		http.Redirect(w, r, "/jeu", http.StatusFound)
	} else {
		AfficherTemplate(w, "Accueil", nil)
	}
}

func JeuHandler(w http.ResponseWriter, r *http.Request) {
	// Si l'état du jeu n'existe pas, initialiser une nouvelle partie
	if variables.Etatjeu == nil {
		variables.Etatjeu = variables.NouvellePartie()
	}

	if r.Method == http.MethodPost {
		// Demander de saisir une lettre ou un mot
		Input := r.FormValue("input")

		// Vérifier si la lettre ou le mot a déjà été tenté
		if fonctions.Doublons(variables.Etatjeu.LettresTentees, Input) || fonctions.Doublons(variables.Etatjeu.MotsTentés, Input) {
			// Si c'est un doublon, renvoyer l'état actuel
			AfficherTemplate(w, "Jeu", variables.Etatjeu)
			return
		}

		// Traiter l'entrée (lettre ou mot)
		utilisateur := fonctions.TraiterChaine(Input, variables.Etatjeu.Vies, variables.Etatjeu.MotCache)

		if len(utilisateur) > 1 {
			// Si l'utilisateur a entré un mot
			fonctions.VerifierMot(variables.Etatjeu.Mot, variables.Etatjeu.MotCache, Input, &variables.Etatjeu.Vies)
			variables.Etatjeu.MotsTentés = append(variables.Etatjeu.MotsTentés, Input)
		} else {
			// Si l'utilisateur a entré une lettre
			fonctions.VerifierUneLettre(variables.Etatjeu.Mot, variables.Etatjeu.MotCache, Input, &variables.Etatjeu.Vies)
			variables.Etatjeu.LettresTentees = append(variables.Etatjeu.LettresTentees, Input)
		}

		// Mettre à jour les infos
		variables.Etatjeu.EtatPendu = fonctions.RenvoieBonhomme(variables.Etatjeu.Vies)
		variables.Etatjeu.MotCache = fonctions.MettreAJourMotCache(variables.Etatjeu.MotCache)
		variables.Etatjeu.MotCacherStr = strings.Join(variables.Etatjeu.MotCache, " ")
	}

	// Vérification de la fin de la partie
	if variables.Etatjeu.Vies <= 0 {
		// Redirige vers la page GameOver
		http.Redirect(w, r, "/perdu", http.StatusFound)
	} else if strings.Join(variables.Etatjeu.MotCache, "") == variables.Etatjeu.Mot {
		// Redirige vers la page Victoire
		http.Redirect(w, r, "/victoire", http.StatusFound)
	} else {
		AfficherTemplate(w, "Jeu", variables.Etatjeu)
	}
}

// Handler pour la page de défaite
func PerduHandler(w http.ResponseWriter, r *http.Request) {
	AfficherTemplate(w, "Perdu", variables.Etatjeu)
}

// Handler pour la page de victoire
func VictoireHandler(w http.ResponseWriter, r *http.Request) {
	AfficherTemplate(w, "Victoire", variables.Etatjeu)
}

// Fonction pour relancer une nouvelle partie
func RelancerHandler(w http.ResponseWriter, r *http.Request) {
	// Réinitialiser l'état du jeu
	variables.Etatjeu = variables.NouvellePartie()
	http.Redirect(w, r, "/", http.StatusFound)
}
