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

var Etatjeu *EtatJeu

// Initialiser une nouvelle partie
func NouvellePartie() *EtatJeu {
	mot := fonctions.ObtenirMotAleatoire()
	motCache := fonctions.GenererTirets(mot)
	motCacherStr := strings.Join(motCache, " ")

	return &EtatJeu{
		Mot:            mot,
		MotCacherStr:   motCacherStr,
		MotCache:       motCache,
		Vies:           10,
		EtatPendu:      fonctions.RenvoieBonhomme(10),
		LettresTentees: []string{},
		MotsTentés:     []string{},
	}
}
