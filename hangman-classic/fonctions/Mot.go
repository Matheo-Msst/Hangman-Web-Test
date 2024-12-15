package fonctions

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

// Obtenir un mot aléatoire à partir du fichier
func ObtenirMotAleatoire() string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		mot := TraiterChaine(scanner.Text(), 10, []string{}) // On ne veut pas afficher d'état ici
		if mot != "" {
			Mots = append(Mots, mot)
		}
	}

	if len(Mots) == 0 {
		log.Fatal("Le fichier 'words.txt' est vide ou ne contient pas de mots valides.")
	}

	rand.Seed(time.Now().UnixNano())
	return Mots[rand.Intn(len(Mots))]
}
