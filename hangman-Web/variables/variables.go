package variables

import (
	"HangmanWeb/hangman-classic/fonctions"
	"strings"
)

type EtatJeu struct {
	Mot            string
	MotCacherStr   string
	MotCache       []string
	Vies           int
	LettresTentees []string
	MotsTentés     []string
	EtatPendu      string
}

// Une carte pour conserver l'état du jeu de chaque joueur
var Etatjeu *EtatJeu

// Initialiser une nouvelle partie
func NouvellePartie() *EtatJeu {
	mot := fonctions.ObtenirMotAleatoire()
	// Initialiser MotCache
	motCache := fonctions.GenererTirets(mot)
	// Joindre les tirets avec un espace pour MotCacherStr
	motCacherStr := strings.Join(motCache, " ")

	return &EtatJeu{
		Mot:            mot,
		MotCacherStr:   motCacherStr,
		MotCache:       motCache,
		Vies:           10,
		EtatPendu:      fonctions.RenvoieBonhomme(10),
		LettresTentees: []string{}, // Liste vide pour les lettres tentées
		MotsTentés:     []string{}, // Liste vide pour les mots tentés
	}
}
