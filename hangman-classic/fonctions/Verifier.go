package fonctions

import (
	"fmt"
	"strings"
)

// Vérifier si le mot proposé est correct
func VerifierMot(mot string, motCache []string, utilisateur string, vies *int) {
	if Doublons(motsTentes, utilisateur) {
		fmt.Println("⛔ Vous avez déjà tenté ce mot. Essayez un autre.")
		return
	}

	motsTentes = append(motsTentes, utilisateur)

	if utilisateur != mot {
		perteVies := len(utilisateur)
		*vies -= perteVies
		if *vies < 0 {
			*vies = 0 // Empêche les vies de devenir négatives
		}
		fmt.Printf("❌ %s n'est pas le bon mot ! Vous perdez %d vies.\n", utilisateur, perteVies)
	} else {
		// Si le mot est trouvé, on remplace motCache par le mot correct
		copy(motCache, strings.Split(mot, ""))
		fmt.Printf("✅ Le mot '%s' est correct !\n", utilisateur)
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
