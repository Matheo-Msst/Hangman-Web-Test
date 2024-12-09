package fonctions

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"
)

// Liste globale des mots
var Mots []string

// Liste des lettres et mots déjà tentés
var lettresTentees []string
var motsTentes []string

// Fonction principale qui orchestre toutes les étapes du jeu
func Jouer() {
	vies := 10
	mot := ObtenirMotAleatoire()
	motCache := GenererTirets(mot)

	// Boucle principale du jeu
	for {
		// Affichage de l'état du jeu avant chaque tour
		afficherEtatJeu(vies, motCache)

		// Affichage des lettres et mots déjà tentés
		fmt.Println("Lettres déjà tentées :", strings.Join(lettresTentees, ", "))
		fmt.Println("Mots déjà tentés :", strings.Join(motsTentes, ", "))

		// Demander une lettre ou un mot (en passant vies et motCache)
		utilisateur := DemanderLettre()

		// La fonction "TraiterChaine" redemande l'entrée si elle n'est pas correcte
		utilisateur = TraiterChaine(utilisateur, vies, motCache)

		// Si l'entrée est valide, on vérifie la lettre ou le mot
		if len(utilisateur) > 1 {
			VerifierMot(mot, motCache, utilisateur, &vies)
		} else {
			VerifierUneLettre(mot, motCache, utilisateur, &vies)
		}

		// Mise à jour et affichage du mot caché à chaque tour
		MettreAJourMotCache(motCache) // Afficher le mot caché mis à jour

		// Vérification de l'état de la partie
		resultat := PartieTerminee(mot, motCache, vies)
		if resultat == 666 {
			// Partie perdue
			fmt.Println("💀 Vous avez perdu !")
			break
		} else if resultat == 777 {
			// Partie gagnée
			fmt.Println("🎉 Vous avez gagné !")
			break
		}
	}
}

// Affiche l'état du jeu (vies, mot caché et l'état du pendu)
func afficherEtatJeu(vies int, motCache []string) {
	// Affichage des vies restantes et du mot caché
	fmt.Printf("\n❤️ Vies : %d, Mot : %s\n", vies, MettreAJourMotCache(motCache))
	// Affichage de l'état actuel du pendu
	fmt.Println(RenvoieBonhomme(vies))
}

// Obtenir un mot aléatoire à partir du fichier
func ObtenirMotAleatoire() string {
	// Obtient le répertoire de travail actuel
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Concatène le répertoire relatif de words.txt
	filePath := filepath.Join(workingDir, "hangman-classic", "fonctions", "words.txt")

	// Ouvre le fichier
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		mot := scanner.Text()
		if mot != "" {
			// Si le mot n'est pas vide, on l'ajoute à Mots
			Mots = append(Mots, mot)
		}
	}

	if len(Mots) == 0 {
		log.Fatal("Le fichier 'words.txt' est vide ou ne contient pas de mots valides.")
	}

	// Choisir un mot aléatoire
	rand.Seed(time.Now().UnixNano())
	selectedWord := Mots[rand.Intn(len(Mots))]
	return selectedWord
}

// Générer les tirets pour le mot caché
func GenererTirets(mot string) []string {
	motCache := make([]string, len(mot))
	for i := range motCache {
		motCache[i] = "_"
	}
	return motCache
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

	// Exemple de traitement
	if strings.Contains(mot, lettre) {
		fmt.Printf("✅ La lettre '%s' est dans le mot.\n", lettre)
	} else {
		fmt.Printf("❌ La lettre '%s' n'est pas dans le mot.\n", lettre)
		*vies--
	}
}

// Vérifier si le mot proposé est correct
func VerifierMot(mot string, motCache []string, utilisateur string, vies *int) {
	if Doublons(motsTentes, utilisateur) {
		fmt.Println("⛔ Vous avez déjà tenté ce mot. Essayez un autre.")
		return
	}

	motsTentes = append(motsTentes, utilisateur)

	if utilisateur != mot {
		perteVies := len(utilisateur)
		*vies -= perteVies
		if *vies < 0 {
			*vies = 0 // Empêche les vies de devenir négatives
		}
		fmt.Printf("❌ %s n'est pas le bon mot ! Vous perdez %d vies.\n", utilisateur, perteVies)
	} else {
		copy(motCache, strings.Split(mot, ""))
		fmt.Printf("✅ Le mot '%s' est correct !\n", utilisateur)
	}
}

// Vérifie si la partie est terminée et renvoie un code de fin de partie (666 = perdu, 777 = gagné, 0 = continue)
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

// Vérifie si le caractère est une lettre valide (A-Z ou a-z)
func EstRuneValide(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
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

// Fonction pour vérifier si une valeur est présente dans une liste
func Doublons(liste []string, element string) bool {
	for _, item := range liste {
		if item == element {
			return true
		}
	}
	return false
}

// Obtient l'entrée de l'utilisateur
func obtenirEntree() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// MettreAJourMotCache prend un tableau de string et renvoie une version string
func MettreAJourMotCache(motCache []string) string {
	// Utilise strings.Join pour assembler les éléments du tableau en une seule chaîne
	motCacheString := strings.Join(motCache, " ")
	return motCacheString
}
