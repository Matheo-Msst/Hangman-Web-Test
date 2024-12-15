package fonctions

import (
	"fmt"
	"strings"
)

// Liste des lettres et mots déjà tentés
var lettresTentees []string
var motsTentes []string

// Fonction qui vérifie si le mot est correct
func VerifierMot(mot string, motCache []string, utilisateur string, vies *int) {
	// Vérifier si le mot a déjà été tenté
	if Doublons(motsTentes, utilisateur) {
		fmt.Println("⛔ Vous avez déjà tenté ce mot. Essayez un autre.")
		return
	}

	// Ajouter le mot aux mots tentés
	motsTentes = append(motsTentes, utilisateur)

	// Si le mot proposé n'est pas correct
	if utilisateur != mot {
		perteVies := len(utilisateur)
		*vies -= perteVies
		if *vies < 0 {
			*vies = 0 // Empêche les vies de devenir négatives
		}
		fmt.Printf("❌ Le mot '%s' n'est pas correct ! Vous perdez %d vies.\n", utilisateur, perteVies)
	} else {
		// Si le mot est correct, remplacer motCache par le mot complet
		copy(motCache, strings.Split(mot, ""))
		fmt.Printf("✅ Le mot '%s' est correct !\n", utilisateur)
	}
}

// Fonction qui vérifie si la lettre a déjà été tentée et la met à jour si elle est correcte
func VerifierUneLettre(mot string, motCache []string, lettre string, vies *int) {
	// Vérifier si la lettre a déjà été tentée
	if Doublons(lettresTentees, lettre) {
		fmt.Printf("⛔ Vous avez déjà tenté la lettre '%s'. Essayez une autre lettre.\n", lettre)
		return
	}
	// Ajouter la lettre aux lettres déjà tentées
	lettresTentees = append(lettresTentees, lettre)

	// Vérification si la lettre est présente dans le mot
	if strings.Contains(mot, lettre) {
		fmt.Printf("✅ La lettre '%s' est dans le mot.\n", lettre)
		// Mise à jour du mot caché
		for i, char := range mot {
			if string(char) == lettre {
				motCache[i] = lettre // Remplace le tiret par la lettre trouvée
			}
		}
	} else {
		fmt.Printf("❌ La lettre '%s' n'est pas dans le mot.\n", lettre)
		*vies-- // Réduit le nombre de vies
	}
}

// Fonction pour vérifier si une valeur est présente dans une liste
func Doublons(liste []string, element string) bool {
	for _, item := range liste {
		if item == element {
			return true
		}
	}
	return false
}
