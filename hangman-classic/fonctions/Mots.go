package fonctions

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

// Liste globale des mots
var Mots []string

// Chemin vers le fichier de mots
const filePath = "hangman-classic/txt/words.txt"

// Obtenir un mot aléatoire à partir du fichier
func ObtenirMotAleatoire() string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		mot, valide := TraiterChaine(scanner.Text())
		if valide {
			Mots = append(Mots, mot)
		}
	}

	if len(Mots) == 0 {
		log.Fatal("Le fichier 'words.txt' est vide ou ne contient pas de mots valides.")
	}

	rand.Seed(time.Now().UnixNano())
	return Mots[rand.Intn(len(Mots))]
}
