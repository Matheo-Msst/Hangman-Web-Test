package fonctions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Obtient l'entrée de l'utilisateur
func obtenirEntree() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// Demander une lettre ou un mot à l'utilisateur
func DemanderLettre() string {
	fmt.Print("\nEntrez une lettre ou un mot : ")
	return obtenirEntree()
}

func VerifierUneLettre(mot string, motCache []string, lettre string, vies *int) {
	if len(mot) == 0 {
		log.Fatal("Erreur: le mot à deviner est vide !")
	}

	if len(motCache) == 0 {
		log.Fatal("Erreur: le mot caché est vide !")
	}

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
