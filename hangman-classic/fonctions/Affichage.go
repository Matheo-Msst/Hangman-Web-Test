package fonctions

import (
	"fmt"
	"strings"
)

// Affiche l'état du jeu (vies, mot caché et l'état du pendu)
func afficherEtatJeu(vies int, motCache []string) {
	// Affichage des vies restantes et du mot caché
	fmt.Printf("\n❤️ Vies : %d, Mot : %s\n", vies, strings.Join(motCache, " "))
	// Affichage de l'état actuel du pendu
	fmt.Println(RenvoieBonhomme(vies))
}
