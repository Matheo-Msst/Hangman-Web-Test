package fonctions

import (
	"bufio"
	"fmt"
	"os"
)

// Fonction pour obtenir l'entrée de l'utilisateur
func obtenirEntree() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// Fonction qui demande à l'utilisateur d'entrer une lettre ou un mot
func DemanderLettre(vies int, motCache []string) string {
	fmt.Print("\nEntrez une lettre ou un mot : ")
	return obtenirEntree()
}
