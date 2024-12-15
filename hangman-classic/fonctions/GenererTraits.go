package fonctions

import "strings"

// Fonction pour générer les tirets pour le mot caché
func GenererTirets(mot string) []string {
	motCache := make([]string, len(mot))
	for i := range motCache {
		motCache[i] = "_"
	}
	return motCache
}

// Fonction pour mettre à jour le mot caché
func MettreAJourMotCache(motCache []string) []string {
	// Utilise strings.Join pour assembler les éléments du tableau en une seule chaîne
	motCacheString := strings.Join(motCache, " ")
	return strings.Split(motCacheString, " ")
}
