package fonctions

import (
	"fmt"
	"strings"
)

// Fonction pour afficher l'état du jeu
func afficherEtatJeu(vies int, motCache []string) {
	// Affichage des vies restantes et du mot caché
	fmt.Printf("\n❤️ Vies : %d, Mot : %s\n", vies, strings.Join(motCache, " "))
	// Affichage de l'état actuel du pendu
	fmt.Println(RenvoieBonhomme(vies))
}
