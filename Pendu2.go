package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func mot_random4() string {
	file, _ := os.Open("langue.txt") //ouvre fichier
	var i, numLigne int
	fileScanner := bufio.NewScanner(file)
	var mot string
	numLigne = rand.Intn(24) + 1
	for fileScanner.Scan() && i < numLigne {
		mot = fileScanner.Text()
		i++
	}
	fmt.Printf("Le nombre de caractères du mot à deviner est: %d\n", len(mot)) //affichage nombre de caractères
	file.Close()
	return mot
}

func mot_random2() string {
	file, _ := os.Open("environnement.txt") //ouvre fichier
	var i, numLigne int
	fileScanner := bufio.NewScanner(file)
	var mot string
	numLigne = rand.Intn(24) + 1
	for fileScanner.Scan() && i < numLigne {
		mot = fileScanner.Text()
		i++
	}
	fmt.Printf("Le nombre de caractères du mot à deviner est: %d\n", len(mot)) //affichage nombre de caractères
	file.Close()
	return mot
}
func mot_random3() string {
	file, _ := os.Open("animaux.txt") //ouvre fichier
	var i, numLigne int
	fileScanner := bufio.NewScanner(file)
	var mot string
	numLigne = rand.Intn(24) + 1
	for fileScanner.Scan() && i < numLigne {
		mot = fileScanner.Text()
		i++
	}
	fmt.Printf("Le nombre de caractères du mot à deviner est: %d\n", len(mot)) //affichage nombre de caractères
	file.Close()
	return mot
}
func main() {
	choix()
}

func choix() {
	var choix string
	var verifier bool
	for !verifier {
		fmt.Print("Choisissez entre ces fichiers pour le pendu : animaux, environnement, test ou langue? ")
		fmt.Scanln(&choix)
		if choix == "animaux" {
			mot_random3()
			affichage_jeu3()
			verifier = true
		} else if choix == "environnement" {
			mot_random2()
			affichage_jeu2()
			verifier = true
		} else if choix == "langue" {
			mot_random4()
			affichage_jeu4()
			verifier = true
		} else if choix == "test" {
			mot_random()
			affichage_jeu()
			verifier = true
		} else {
			fmt.Println("Choisissez animaux, environnement, test ou langue !")
		}
	}
}

// dessine le hangman dans sa potence
func hangman_dessin(nombre_essai int) {
	file, _ := os.Open("hangman.txt")
	defer file.Close()

	// Lire les lignes du fichier et les stocker dans un tableau
	var etapes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		etapes = append(etapes, scanner.Text())
	}

	// Chaque étape du pendu
	lignesParEtape := 7
	totalEtapes := len(etapes) / lignesParEtape

	// Calculer quelle étape afficher selon nombre essai
	etapeActuelle := totalEtapes - nombre_essai
	if etapeActuelle < 0 {
		etapeActuelle = 0
	} else if etapeActuelle >= totalEtapes {
		etapeActuelle = totalEtapes
	}

	// Afficher l'étape correspondante
	debut := etapeActuelle * lignesParEtape
	fin := debut + lignesParEtape
	for i := debut; i < fin && i < len(etapes); i++ {
		fmt.Println(etapes[i])
	}
}

// choisi un mot aléatoire dans un fichier texte
func mot_random() string {
	file, _ := os.Open("test.txt") //ouvre fichier
	var i, numLigne int
	fileScanner := bufio.NewScanner(file)
	var mot string
	numLigne = rand.Intn(24) + 1
	for fileScanner.Scan() && i < numLigne {
		mot = fileScanner.Text()
		i++
	}
	fmt.Printf("Le nombre de caractères du mot à deviner est: %d\n", len(mot)) //affichage nombre de caractères
	file.Close()
	return mot
}

func affichage_jeu4() {
	var mot_caché string
	fmt.Println("Bienvenue dans le pendu/hangman !")
	mot := mot_random4()
	for i := 0; i < len(mot); i++ {
		mot_caché += "_ " //mot caché _ séparé par des espaces
	}
	fmt.Println("Mot actuel :", mot_caché)
	word(mot, mot_caché)
}

func affichage_jeu3() {
	var mot_caché string
	fmt.Println("Bienvenue dans le pendu/hangman !")
	mot := mot_random3()
	for i := 0; i < len(mot); i++ {
		mot_caché += "_ " //mot caché _ séparé par des espaces
	}
	fmt.Println("Mot actuel :", mot_caché)
	word(mot, mot_caché)
}

func affichage_jeu2() {
	var mot_caché string
	fmt.Println("Bienvenue dans le pendu/hangman !")
	mot := mot_random2()
	for i := 0; i < len(mot); i++ {
		mot_caché += "_ " //mot caché _ séparé par des espaces
	}
	fmt.Println("Mot actuel :", mot_caché)
	word(mot, mot_caché)
}

// fonction pour afficher le début du jeu
func affichage_jeu() {
	var mot_caché string
	fmt.Println("Bienvenue dans le pendu/hangman !")
	mot := mot_random()
	for i := 0; i < len(mot); i++ {
		mot_caché += "_ " //mot caché _ séparé par des espaces
	}
	fmt.Println("Mot actuel :", mot_caché)
	word(mot, mot_caché)
}

// fonction pour les essais, se soustrait si on l'appelle (mauvaise lettre tapé)
func nombre_essais(essais_restants int) int {
	essais_restants--
	return essais_restants
}

// fonction qui compare le contenu du mot que l'on cherche à ce que l'on tape
func word(mot, mot_caché string) {
	nombre_essai := 10
	var lettre string
	var lettres_deja_tapees string

	for nombre_essai > 0 {
		fmt.Println("Vous avez", nombre_essai, "essais !")
		fmt.Print("Tapez une lettre ou le mot entier SANS faute : ")
		fmt.Scanln(&lettre)

		//condition pour vérifier si ce que l'on tape est + long que le mot à deviner
		if len(lettre) > len(mot) {
			fmt.Printf("Entrée invalide : le mot saisi est trop long.")
			fmt.Println("")

			continue
		}

		// condition pour verifier les lettres en minuscule et majuscule
		if !((lettre[0] >= 'a' && lettre[0] <= 'z') || (lettre[0] >= 'A' && lettre[0] <= 'Z')) {
			fmt.Println("Entrée invalide. Veuillez taper une lettre alphabétique.")
			continue
		}

		//condition pour vérifier si l'utilisateur tente de deviner le mot en l'écrivant
		if len(lettre) > 1 {
			if lettre == mot {
				fmt.Println("Félicitations ! Vous avez deviné le mot :", mot)
				return

			} else {
				fmt.Println("Mauvaise tentative, le mot n'est pas correct !")
				fmt.Println("")
				hangman_dessin(nombre_essai)
				fmt.Println("")
				nombre_essai = nombre_essais(nombre_essai)
				continue
			}
		}

		//condition pour lettre afin de vérifier si c'est déjà tapé ou non
		if vérification_lettre_tapée(lettres_deja_tapees, lettre) {
			fmt.Println("Vous avez déjà tapé cette lettre. Essayez une autre.")
			continue
		}

		//collecte des lettres tapés par l'utilisateur afin de rappeler + facilement ce qu'il avait tapé
		lettres_deja_tapees += lettre
		fmt.Println("Les lettres utilisées sont: ", lettres_deja_tapees, ".")

		//vérifie si la lettre est dans le mot
		if verifier_lettres(lettre, mot) {
			fmt.Println("Vous avez trouvé une lettre !")
			mot_caché = nouv_mot_caché(lettre, mot, mot_caché)
		} else {
			fmt.Println("Vous avez raté votre chance! ", lettre, "n'est pas dans le mot.")
			fmt.Println("")
			hangman_dessin(nombre_essai)
			nombre_essai = nombre_essais(nombre_essai)
		}
		fmt.Println("")
		fmt.Println("Mot actuel :", mot_caché)

		if mot_caché == mot {
			fmt.Println("Félicitations ! Vous avez gagné !")
			return
		}
	}
	fmt.Println("Dommage, vous avez perdu ! Le mot était :", mot)

	if mot_caché == mot {
		fmt.Println("Félicitations ! Vous avez gagné !")
		rejouer()
		return
	}
	fmt.Println("Dommage, vous avez perdu ! Le mot était :", mot)
	rejouer()
}

// fonction pour le mot caché à chaque fois qu'une/des lettre/s est/sont trouvée/s
func nouv_mot_caché(lettres, mot, mot_caché string) string {
	maj_mot_caché := ""
	index := 0
	for i := 0; i < len(mot); i++ {

		//vérifie si l'indice du mot caché correspond à une lettre
		if string(mot[i]) == lettres {
			maj_mot_caché += string(mot[i]) + " "
		} else {
			maj_mot_caché += string(mot_caché[index]) + " "
		}
		index += 2 // espace de 2 entre les lettres
	}
	return maj_mot_caché
}

// vérifie si une lettre tapé est une lettre du mot
func verifier_lettres(lettres, mot string) bool {
	for i := 0; i < len(mot); i++ {
		if string(mot[i]) == lettres {
			return true
		}
	}
	return false
}

// vérifie si la lettre tapé est une lettre déjà tapée
func vérification_lettre_tapée(lettres_deja_tapees, lettre string) bool {
	for i := 0; i < len(lettres_deja_tapees); i++ {
		if string(lettres_deja_tapees[i]) == lettre {
			return true
		}
	}
	return false
}

func rejouer() {
	var choix string
	for {
		fmt.Print("Voulez-vous rejouer ? (oui/non) : ")
		fmt.Scanln(&choix)
		if choix == "oui" {
			fmt.Println("Super ! Relançons une nouvelle partie.")
			affichage_jeu()
			return // L'utilisateur veut rejouer
		} else if choix == "non" {
			fmt.Println("Merci d'avoir joué ! À bientôt.")
			return // L'utilisateur ne veut pas rejouer
		} else {
			fmt.Println("Entrée invalide. Veuillez taper 'oui' ou 'non'.")
		}
	}
}
