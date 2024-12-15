package fonctions

// EtatsBonhomme contient les différentes étapes du pendu sous forme de tableau de chaînes
var EtatsBonhomme = []string{
	// 10 vies (aucun élément)
	`







`,
	// 9 vies (sol uniquement)
	`






===========
`,
	// 8 vies (base de la potence)
	`

|          
|          
|          
|          
|          
===========
`,
	// 7 vies (potence vide)
	`
=========  
|          
|          
|          
|          
|          
===========
`,
	// 6 vies (juste la potence)
	`
=========  
|/         
|          
|          
|          
|          
===========
`,
	// 5 vies
	`
=========  
|/  |      
|          
|          
|          
|          
===========
`,
	// 4 vies
	`
=========
|/  |
|   O
|   
|   
|
===========
`,
	// 3 vies
	`
=========
|/  |
|   O
|   |
|   
|
===========
`,
	// 2 vies
	`
=========
|/  |
|   O
|  /|
|   
|
===========
`,
	// 1 vie
	`
=========
|/  |
|   O
|  /|\
|  /
|
===========
`,
	// 0 vies (le pendu complet)
	`
=========
|/  |
|   O
|  /|\
|  / \
|
===========
`,
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
