package fonctions

import (
	"fmt"
	"strings"
)

// Liste globale des mots
var Mots []string

// Chemin vers le txt
const filePath = "hangman-classic/txt/words.txt"

// Liste des lettres et mots déjà tentés
var lettresTentees []string
var motsTentes []string

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
		utilisateur := DemanderLettre()

		// La fonction "TraiterChaine" redemande l'entrée si elle n'est pas correcte
		utilisateur = TraiterChaine(utilisateur, vies, motCache)

		// Si l'entrée est valide, on vérifie la lettre ou le mot
		if len(utilisateur) > 1 {
			VerifierMot(mot, motCache, utilisateur, &vies)
		} else {
			VerifierUneLettre(mot, motCache, utilisateur, &vies)
		}

		// Mise à jour et affichage du mot caché à chaque tour
		MettreAJourMotCache(motCache) // Afficher le mot caché mis à jour

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
