package fonctions

import (
	"fmt"
	"unicode"
)

// Vérifie si le caractère est une lettre valide (A-Z ou a-z)
func EstRuneValide(r rune) bool {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
		return true
	}
	// Vérifie si le caractère est un chiffre (0-9), retourne false dans ce cas
	if (r >= '0' && r <= '9') || (r == 'à' || r == 'ù' || r == 'ç' || r == '_' || r == '-' || r == 'é') {
		return false
	}
	return false
}

// Traite la chaîne en acceptant uniquement les lettres de A à Z
func TraiterChaine(input string, vies int, motCache []string) string {
	var resultat []rune
	for _, r := range input {
		if !EstRuneValide(r) {
			fmt.Println("⛔ Veuillez saisir uniquement des lettres de A à Z (pas d'accents, chiffres ou symboles).")
			return DemanderLettre() // Demande une nouvelle entrée
		}
		resultat = append(resultat, unicode.ToLower(r))
	}
	return string(resultat)
}
