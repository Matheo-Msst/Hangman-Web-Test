package fonctions

// EtatsBonhomme contient les différentes étapes du pendu sous forme de tableau de chaînes
var EtatsBonhomme = []string{
	`

`,
	`

===========`,
	`

|
|
|
|
|
===========`,
	`
=========
|
|
|
|
|
===========`,
	`
=========
|/
|
|
|
|
===========`,
	`
=========
|/  |
|
|
|
|
===========`,
	`
=========
|/  |
|   O
|
|
|
===========`,
	`
=========
|/  |
|   O
|   |
|
|
===========`,
	`
=========
|/  |
|   O
|  /|
|
|
===========`,
	`
=========
|/  |
|   O
|  /|\
|  /
|
===========`,
	`
=========
|/  |
|   O
|  /|\
|  / \
|
===========`,
}

// Fonction qui renvoie l'état du pendu selon les vies restantes
func RenvoieBonhomme(vies int) string {
	// Calcul de l'index basé sur le nombre de vies restantes
	IndexEtats := 10 - vies
	if IndexEtats > 10 {
		IndexEtats = 10 // Dernier état (le pendu complet)
	}
	if IndexEtats < 0 {
		IndexEtats = 0 // Assurez-vous que l'index ne soit pas inférieur à 0
	}
	return EtatsBonhomme[IndexEtats]
}
