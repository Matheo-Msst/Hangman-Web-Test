package request

import (
	"HangmanWeb/hangman-classic/fonctions"
	"net/http"
)

func UpdateGame(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	utilisateur := r.FormValue("input-user")

	// Vérifier si l'entrée est valide (lettre ou mot)
	if len(utilisateur) == 0 {
		http.Error(w, "Aucune entrée reçue", http.StatusBadRequest)
		return
	}

	utilisateur = fonctions.TraiterChaine(utilisateur, GameVar.Vies, GameVar.MotCacher)

	// Si l'utilisateur a entré une lettre
	if len(utilisateur) == 1 {
		fonctions.VerifierUneLettre(GameVar.Mot, GameVar.MotCacher, utilisateur, &GameVar.Vies)
	} else { // Si l'utilisateur a entré un mot
		fonctions.VerifierMot(GameVar.Mot, GameVar.MotCacher, utilisateur, &GameVar.Vies)
	}

	// Vérification si la partie est terminée
	resultat := fonctions.PartieTerminee(GameVar.Mot, GameVar.MotCacher, GameVar.Vies)
	if resultat == 666 {
		// Partie perdue
		http.Redirect(w, r, "/gameover", http.StatusSeeOther)
		return
	} else if resultat == 777 {
		// Partie gagnée
		http.Redirect(w, r, "/victoire", http.StatusSeeOther)
		return
	}

	// Afficher l'état du jeu avec le mot caché mis à jour
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}
