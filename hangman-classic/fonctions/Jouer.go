package fonctions

import (
	"fmt"
	"strings"
)

// Fonction principale qui orchestre toutes les étapes du jeu
func Jouer() {
	vies := 10
	mot := ObtenirMotAleatoire()
	motCache := GenererTirets(mot)

	// Boucle principale du jeu
	for {
		// Affichage de l'état du jeu avant chaque tour
		afficherEtatJeu(vies, motCache)

		// Affichage des lettres et mots déjà tentés
		fmt.Println("Lettres déjà tentées :", strings.Join(lettresTentees, ", "))
		fmt.Println("Mots déjà tentés :", strings.Join(motsTentes, ", "))

		// Demander une lettre ou un mot (en passant vies et motCache)
		utilisateur := DemanderLettre(vies, motCache)

		// Traiter la chaîne de l'utilisateur
		utilisateur, valide := TraiterChaine(utilisateur)

		// Si l'entrée est invalide, redemander l'entrée
		if !valide {
			fmt.Println("⛔ Veuillez entrer une lettre valide (A-Z, pas de symboles, pas de chiffres).")
			continue
		}

		// Convertir l'entrée en minuscules
		utilisateur = Minuscule(utilisateur)

		// Si l'utilisateur entre un mot complet
		if len(utilisateur) > 1 {
			VerifierMot(mot, motCache, utilisateur, &vies)
		} else {
			// Si l'utilisateur entre une lettre
			VerifierUneLettre(mot, motCache, utilisateur, &vies)
		}

		// Mise à jour et affichage du mot caché à chaque tour
		MettreAJourMotCache(motCache)

		// Vérification de l'état de la partie
		resultat := PartieTerminee(mot, motCache, vies)
		if resultat == 666 {
			// Partie perdue
			fmt.Println("💀 Vous avez perdu !")
			break
		} else if resultat == 777 {
			// Partie gagnée
			fmt.Println("🎉 Vous avez gagné !")
			break
		}
	}
}
