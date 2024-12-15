package fonctions

import (
	"fmt"
	"strings"
)

// (666 = perdu, 777 = gagné, 0 = continue)
func PartieTerminee(mot string, motCache []string, vies int) int {
	if vies <= 0 {
		fmt.Printf("\n💀 Vous avez perdu ! Le mot était : %s\n", mot)
		return 666 // Partie perdue
	} else if strings.Join(motCache, "") == mot {
		fmt.Printf("\n🎉 Bravo ! Vous avez trouvé le mot : %s\n", mot)
		return 777 // Partie gagnée
	}
	return 0 // La partie continue
}
