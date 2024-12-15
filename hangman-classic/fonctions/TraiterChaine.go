package fonctions

import "strings"

// Fonction pour vérifier si la chaîne est une lettre valide (A-Z)
func EstRuneValide(r rune) bool {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
		return true
	}
	return false
}

// Fonction qui traite l'entrée utilisateur et retourne un booléen indiquant si l'entrée est valide
func TraiterChaine(input string) (string, bool) {
	// Vérification si l'entrée est valide
	if !EstRuneValide(rune(input[0])) {
		// Entrée invalide
		return "", false
	}
	// Sinon, retourner l'entrée convertie en minuscule et valider
	return strings.ToLower(input), true
}

// Fonction pour convertir l'entrée en minuscules si valide
func Minuscule(input string) string {
	return strings.ToLower(input)
}
