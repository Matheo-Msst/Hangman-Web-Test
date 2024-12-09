package fonctions

// EtatsBonhomme contient les différentes étapes du pendu sous forme de tableau de chaînes
var EtatsBonhomme = []string{
	"",
	"\n\n\n\n\n\n===========",
	"|\n|\n|\n|\n|\n=========",
	"=========\n|\n|\n|\n|\n|\n===========",
	"=========\n|/\n|\n|\n|\n|\n===========",
	"=========\n|/  |\n|\n|\n|\n|\n===========",
	"=========\n|/  |\n|   O\n|\n|\n|\n===========",
	"=========\n|/  |\n|   O\n|  /|\\\n|\n|\n===========",
	"=========\n|/  |\n|   O\n|  /|\\\n|\n|\n===========",
	"=========\n|/  |\n|   O\n|  /|\\\n|  /\n|\n===========",
}

// Fonction qui renvoie l'état du pendu selon les vies restantes
func RenvoieBonhomme(vies int) string {
	IndexEtats := 10 - vies
	if IndexEtats < 0 {
		IndexEtats = 0
	}
	return EtatsBonhomme[IndexEtats]
}
