package variables

import (
	"HangmanWeb/hangman-classic/fonctions"
	"log"
)

// GameState représente l'état actuel d'une partie du jeu
type GameState struct {
	Mot             string
	MotCacher       []string
	MotCacherString string
	Vies            int
	LettresTentees  []string
	MotsTentes      []string
	PartieTerminee  int
	DisplayHangman  string
}

func InitialiserGameState(vies int) GameState {
	mot := fonctions.ObtenirMotAleatoire()
	if mot == "" {
		log.Fatal("Erreur: le mot obtenu est vide !")
	}

	motCache := fonctions.GenererTirets(mot)
	if len(motCache) == 0 {
		log.Fatal("Erreur: le mot caché est vide !")
	}

	moCacheString := fonctions.MettreAJourMotCache(motCache)
	displayHangman := fonctions.RenvoieBonhomme(vies)

	return GameState{
		Mot:             mot,
		MotCacher:       motCache,
		MotCacherString: moCacheString,
		Vies:            vies,
		LettresTentees:  []string{},
		MotsTentes:      []string{},
		PartieTerminee:  0,
		DisplayHangman:  displayHangman,
	}
}
