package fonctions

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"os"
// 	"strings"
// 	"time"
// )

// // Liste globale des mots
// var Mots []string

// // Chemin vers le fichier de mots
// const filePath = "hangman-classic/txt/words.txt"

// // Liste des lettres et mots déjà tentés
// var lettresTentees []string
// var motsTentes []string

// // EtatsBonhomme contient les différentes étapes du pendu sous forme de tableau de chaînes
// var EtatsBonhomme = []string{
// 	`

// `,
// 	`

// ===========`,
// 	`

// |
// |
// |
// |
// |
// ===========`,
// 	`
// =========
// |
// |
// |
// |
// |
// ===========`,
// 	`
// =========
// |/
// |
// |
// |
// |
// ===========`,
// 	`
// =========
// |/  |
// |
// |
// |
// |
// ===========`,
// 	`
// =========
// |/  |
// |   O
// |
// |
// |
// ===========`,
// 	`
// =========
// |/  |
// |   O
// |   |
// |
// |
// ===========`,
// 	`
// =========
// |/  |
// |   O
// |  /|
// |
// |
// ===========`,
// 	`
// =========
// |/  |
// |   O
// |  /|\
// |  /
// |
// ===========`,
// 	`
// =========
// |/  |
// |   O
// |  /|\
// |  / \
// |
// ===========`,
// }

// // Fonction qui renvoie l'état du pendu selon les vies restantes
// func RenvoieBonhomme(vies int) string {
// 	// Calcul de l'index basé sur le nombre de vies restantes
// 	IndexEtats := 10 - vies
// 	if IndexEtats > 10 {
// 		IndexEtats = 10 // Dernier état (le pendu complet)
// 	}
// 	if IndexEtats < 0 {
// 		IndexEtats = 0 // Assurez-vous que l'index ne soit pas inférieur à 0
// 	}
// 	return EtatsBonhomme[IndexEtats]
// }

// // Fonction pour afficher l'état du jeu
// func afficherEtatJeu(vies int, motCache []string) {
// 	// Affichage des vies restantes et du mot caché
// 	fmt.Printf("\n❤️ Vies : %d, Mot : %s\n", vies, strings.Join(motCache, " "))
// 	// Affichage de l'état actuel du pendu
// 	fmt.Println(RenvoieBonhomme(vies))
// }

// // Obtenir un mot aléatoire à partir du fichier
// func ObtenirMotAleatoire() string {
// 	f, err := os.Open(filePath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	scanner := bufio.NewScanner(f)
// 	for scanner.Scan() {
// 		mot := TraiterChaine(scanner.Text(), 10, []string{}) // On ne veut pas afficher d'état ici
// 		if mot != "" {
// 			Mots = append(Mots, mot)
// 		}
// 	}

// 	if len(Mots) == 0 {
// 		log.Fatal("Le fichier 'words.txt' est vide ou ne contient pas de mots valides.")
// 	}

// 	rand.Seed(time.Now().UnixNano())
// 	return Mots[rand.Intn(len(Mots))]
// }

// // Fonction pour générer les tirets pour le mot caché
// func GenererTirets(mot string) []string {
// 	motCache := make([]string, len(mot))
// 	for i := range motCache {
// 		motCache[i] = "_"
// 	}
// 	return motCache
// }

// // Fonction pour mettre à jour le mot caché
// func MettreAJourMotCache(motCache []string) []string {
// 	// Utilise strings.Join pour assembler les éléments du tableau en une seule chaîne
// 	motCacheString := strings.Join(motCache, " ")
// 	return strings.Split(motCacheString, " ")
// }

// // Fonction qui vérifie si le mot est correct
// func VerifierMot(mot string, motCache []string, utilisateur string, vies *int) {
// 	// Vérifier si le mot a déjà été tenté
// 	if Doublons(motsTentes, utilisateur) {
// 		fmt.Println("⛔ Vous avez déjà tenté ce mot. Essayez un autre.")
// 		return
// 	}

// 	// Ajouter le mot aux mots tentés
// 	motsTentes = append(motsTentes, utilisateur)

// 	// Si le mot proposé n'est pas correct
// 	if utilisateur != mot {
// 		perteVies := len(utilisateur)
// 		*vies -= perteVies
// 		if *vies < 0 {
// 			*vies = 0 // Empêche les vies de devenir négatives
// 		}
// 		fmt.Printf("❌ Le mot '%s' n'est pas correct ! Vous perdez %d vies.\n", utilisateur, perteVies)
// 	} else {
// 		// Si le mot est correct, remplacer motCache par le mot complet
// 		copy(motCache, strings.Split(mot, ""))
// 		fmt.Printf("✅ Le mot '%s' est correct !\n", utilisateur)
// 	}
// }

// // Fonction pour vérifier si une valeur est présente dans une liste
// func Doublons(liste []string, element string) bool {
// 	for _, item := range liste {
// 		if item == element {
// 			return true
// 		}
// 	}
// 	return false
// }

// // Fonction qui vérifie si la lettre a déjà été tentée et la met à jour si elle est correcte
// func VerifierUneLettre(mot string, motCache []string, lettre string, vies *int) {
// 	// Vérifier si la lettre a déjà été tentée
// 	if Doublons(lettresTentees, lettre) {
// 		fmt.Printf("⛔ Vous avez déjà tenté la lettre '%s'. Essayez une autre lettre.\n", lettre)
// 		return
// 	}
// 	// Ajouter la lettre aux lettres déjà tentées
// 	lettresTentees = append(lettresTentees, lettre)

// 	// Vérification si la lettre est présente dans le mot
// 	if strings.Contains(mot, lettre) {
// 		fmt.Printf("✅ La lettre '%s' est dans le mot.\n", lettre)
// 		// Mise à jour du mot caché
// 		for i, char := range mot {
// 			if string(char) == lettre {
// 				motCache[i] = lettre // Remplace le tiret par la lettre trouvée
// 			}
// 		}
// 	} else {
// 		fmt.Printf("❌ La lettre '%s' n'est pas dans le mot.\n", lettre)
// 		*vies-- // Réduit le nombre de vies
// 	}
// }

// // Fonction qui traite l'entrée utilisateur
// func TraiterChaine(input string, vies int, motCache []string) string {
// 	// Vérification si l'entrée est valide
// 	if !EstRuneValide(rune(input[0])) {
// 		fmt.Println("⛔ Veuillez entrer une lettre valide (A-Z, pas de symboles, pas de chiffres).")
// 		return DemanderLettre(vies, motCache) // Redemander l'entrée si elle est invalide
// 	}
// 	// Sinon, retourner l'entrée convertie en minuscule
// 	return strings.ToLower(input)
// }

// // Fonction pour vérifier la validité de la chaîne (lettres uniquement)
// func EstRuneValide(r rune) bool {
// 	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
// 		return true
// 	}
// 	return false
// }

// // Fonction pour obtenir l'entrée de l'utilisateur
// func obtenirEntree() string {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	return scanner.Text()
// }

// // Fonction qui demande à l'utilisateur d'entrer une lettre ou un mot
// func DemanderLettre(vies int, motCache []string) string {
// 	fmt.Print("\nEntrez une lettre ou un mot : ")
// 	return obtenirEntree()
// }

// // Fonction pour vérifier si la partie est terminée
// func PartieTerminee(mot string, motCache []string, vies int) int {
// 	if vies <= 0 {
// 		fmt.Printf("\n💀 Vous avez perdu ! Le mot était : %s\n", mot)
// 		return 666 // Partie perdue
// 	} else if strings.Join(motCache, "") == mot {
// 		fmt.Printf("\n🎉 Bravo ! Vous avez trouvé le mot : %s\n", mot)
// 		return 777 // Partie gagnée
// 	}
// 	return 0 // La partie continue
// }

// // Fonction principale qui orchestre toutes les étapes du jeu
// func Jouer() {
// 	vies := 10
// 	mot := ObtenirMotAleatoire()
// 	motCache := GenererTirets(mot)

// 	// Boucle principale du jeu
// 	for {
// 		// Affichage de l'état du jeu avant chaque tour
// 		afficherEtatJeu(vies, motCache)

// 		// Affichage des lettres et mots déjà tentés
// 		fmt.Println("Lettres déjà tentées :", strings.Join(lettresTentees, ", "))
// 		fmt.Println("Mots déjà tentés :", strings.Join(motsTentes, ", "))

// 		// Demander une lettre ou un mot (en passant vies et motCache)
// 		utilisateur := DemanderLettre(vies, motCache)

// 		// Traiter la chaîne de l'utilisateur
// 		utilisateur = TraiterChaine(utilisateur, vies, motCache)

// 		// Si l'utilisateur entre un mot complet
// 		if len(utilisateur) > 1 {
// 			VerifierMot(mot, motCache, utilisateur, &vies)
// 		} else {
// 			// Si l'utilisateur entre une lettre
// 			VerifierUneLettre(mot, motCache, utilisateur, &vies)
// 		}

// 		// Mise à jour et affichage du mot caché à chaque tour
// 		MettreAJourMotCache(motCache)

// 		// Vérification de l'état de la partie
// 		resultat := PartieTerminee(mot, motCache, vies)
// 		if resultat == 666 {
// 			// Partie perdue
// 			fmt.Println("💀 Vous avez perdu !")
// 			break
// 		} else if resultat == 777 {
// 			// Partie gagnée
// 			fmt.Println("🎉 Vous avez gagné !")
// 			break
// 		}
// 	}
// }
