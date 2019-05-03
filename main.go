package main

import "fmt"
import "os"
import "strings"
import "strconv"
import "os/exec"

//import "io"

/*
Fonction générique pour l'écriture via les pointeurs des param
*/
func ask_full_name(ptr_firstname *string, ptr_surname *string) {
	var firstname string
	var surname string

	if len(*ptr_firstname) == 0 {
		fmt.Println("Attention : Mets la majuscule au début de chaque mot !")
		fmt.Println(("Prénom :"))
		fmt.Scanln(&firstname)
		fmt.Println(("\nNom :"))
		fmt.Scanln(&surname)
		*ptr_firstname = strings.ToTitle(firstname)
		*ptr_surname = strings.ToTitle(surname)

	} else {
		var choix_id string
		show_parameters(ptr_firstname, ptr_surname)
		fmt.Println("\nGarder les mêmes paramètres ? (y/n)")
		fmt.Scanln(&choix_id)

		if choix_id == "y" {

		} else if choix_id == "n" {
			fmt.Println("Attention : Mets la majuscule au début de chaque mot !")
			fmt.Println(("Prénom :"))
			fmt.Scanln(&firstname)
			fmt.Println(("\nNom :"))
			fmt.Scanln(&surname)

			*ptr_firstname = firstname
			*ptr_surname = surname
		} else {
			fmt.Println("Réponse inconnue, retour au menu")
			return
		}
	}
}

/*
Création d'un .txt avec le nom et le prénom
*/
func create_name_list(ptr_firstname *string, ptr_surname *string) {
	ask_full_name(ptr_firstname, ptr_surname)

	//créer le fichier
	f, err := os.Create("name.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//écrire dedans
	f.WriteString(*ptr_firstname + *ptr_surname + "\n")
	f.WriteString(*ptr_surname + *ptr_firstname + "\n")
	f.WriteString(strings.ToUpper(*ptr_firstname) + strings.ToUpper(*ptr_surname) + "\n")
	f.WriteString(strings.ToUpper(*ptr_surname) + strings.ToUpper(*ptr_firstname) + "\n")
	f.WriteString(strings.ToLower(*ptr_firstname) + strings.ToLower(*ptr_surname) + "\n")
	f.WriteString(strings.ToLower(*ptr_surname) + strings.ToLower(*ptr_firstname))

	f.Close()
	fmt.Println("Fait !")
}

/*
Création d'un .txt avec le nom, prénom et année
*/
func create_name_and_year_list(ptr_firstname *string, ptr_surname *string) {
	var year string

	ask_full_name(ptr_firstname, ptr_surname)

	fmt.Println("Entre l'année : ")
	fmt.Scanln(&year)

	//créer le fichier
	f, err := os.Create("nameyear.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//écrire dedans
	f.WriteString(*ptr_firstname + *ptr_surname + year + "\n")
	f.WriteString(*ptr_surname + *ptr_firstname + year + "\n")
	f.WriteString(strings.ToUpper(*ptr_firstname) + strings.ToUpper(*ptr_surname) + year + "\n")
	f.WriteString(strings.ToUpper(*ptr_surname) + strings.ToUpper(*ptr_firstname) + year + "\n")
	f.WriteString(strings.ToLower(*ptr_firstname) + strings.ToLower(*ptr_surname) + year + "\n")
	f.WriteString(strings.ToLower(*ptr_surname) + strings.ToLower(*ptr_firstname) + year)

	f.Close()
	fmt.Println("Fait !")
}

/*
Création d'un .txt avec le nom, prénom et nombres
*/
func create_name_and_number_list(ptr_firstname *string, ptr_surname *string) {

	var min int
	var max int

	ask_full_name(ptr_firstname, ptr_surname)

	fmt.Println(("\nnombre min :"))
	fmt.Scanln(&min)
	fmt.Println("\nnombre max : ")
	fmt.Scanln(&max)

	//créer le fichier
	f, err := os.Create("namenumber.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//écrire dedans
	for i := min; i <= max; i++ {
		f.WriteString(*ptr_firstname + *ptr_surname + strconv.FormatInt(int64(i), 10) + "\n")
		f.WriteString(*ptr_surname + *ptr_firstname + strconv.FormatInt(int64(i), 10) + "\n")
		f.WriteString(strings.ToUpper(*ptr_firstname) + strings.ToUpper(*ptr_surname) + strconv.FormatInt(int64(i), 10) + "\n")
		f.WriteString(strings.ToUpper(*ptr_surname) + strings.ToUpper(*ptr_firstname) + strconv.FormatInt(int64(i), 10) + "\n")
		f.WriteString(strings.ToLower(*ptr_firstname) + strings.ToLower(*ptr_surname) + strconv.FormatInt(int64(i), 10) + "\n")
		f.WriteString(strings.ToLower(*ptr_surname) + strings.ToLower(*ptr_firstname) + strconv.FormatInt(int64(i), 10) + "\n")
	}

	f.Close()
	fmt.Println("Fait !")
}

/*
Création d'un .txt avec toutes les dates allant de 1/1/1960 à 31/12/2021
*/
func create_date_list() {

	f, err := os.Create("dates.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for y := 1960; y <= 2021; y++ {

		for m := 1; m <= 12; m++ {

			for d := 1; d <= 31; d++ {

				if m < 10 {

					if d < 10 {
						f.WriteString("0" + strconv.FormatInt(int64(d), 10) + "0" + strconv.FormatInt(int64(m), 10) + strconv.FormatInt(int64(y), 10) + "\n")
						f.WriteString(strconv.FormatInt(int64(y), 10) + "0" + strconv.FormatInt(int64(m), 10) + "0" + strconv.FormatInt(int64(d), 10) + "\n")
					} else {
						f.WriteString(strconv.FormatInt(int64(d), 10) + "0" + strconv.FormatInt(int64(m), 10) + strconv.FormatInt(int64(y), 10) + "\n")
						f.WriteString(strconv.FormatInt(int64(y), 10) + "0" + strconv.FormatInt(int64(m), 10) + strconv.FormatInt(int64(d), 10) + "\n")
					}
				} else if d < 10 {
					f.WriteString("0" + strconv.FormatInt(int64(d), 10) + strconv.FormatInt(int64(m), 10) + strconv.FormatInt(int64(y), 10) + "\n")
					f.WriteString(strconv.FormatInt(int64(y), 10) + strconv.FormatInt(int64(m), 10) + "0" + strconv.FormatInt(int64(d), 10) + "\n")
				} else {
					f.WriteString(strconv.FormatInt(int64(d), 10) + strconv.FormatInt(int64(m), 10) + strconv.FormatInt(int64(y), 10) + "\n")
					f.WriteString(strconv.FormatInt(int64(y), 10) + strconv.FormatInt(int64(m), 10) + strconv.FormatInt(int64(d), 10) + "\n")
				}
			}
		}
	}

	fmt.Println("Fait !")
}

/*
Supprimer un fichier
*/
func delete_list() {
	var name string

	fmt.Println("Entre le nom du fichier à supprimer :")
	fmt.Scanln(&name)

	if strings.Contains(name, "/") {
		fmt.Println("Le fichier doit être dans le même répertoire que ce script !")
	} else if strings.Contains(name, ".go") {
		fmt.Println("Bien tenté")
	} else {
		fmt.Println("Suppression de " + name)
		os.Remove(name)
		fmt.Println("Supprimé !")
	}
}

/*
Afficher les options
*/
func options_menu() {
	fmt.Println("Options de création :")
	fmt.Println(" [1] : nom uniquement")
	fmt.Println(" [2] : nom + année")
	fmt.Println(" [3] : nom + nombre")
	fmt.Println(" [4] : dates")
	fmt.Println(" [a] : toutes les options ci-dessus, dans des fichiers différents")
	fmt.Println(" [fa]: toutes les options ci-dessus (- a) dans un même fichier")
	fmt.Println("\nOptions de modification :")
	fmt.Println(" [ls]: montre tous les fichiers existants dans ce repertoire")
	fmt.Println(" [m] : merge des fichiers ensemble")
	fmt.Println(" [d] : supprimer un fichier")
	fmt.Println("\nOptions diverses :")
	fmt.Println(" [h] : menu d'aide")
	fmt.Println(" [p] : afficher les paramètres actuels")
	fmt.Println(" [q] : quitter")

}

func show_parameters(ptr_firstname *string, ptr_surname *string) {
	fmt.Println("\nParamètres actuels :")
	fmt.Println(" Prénom : ", *ptr_firstname)
	fmt.Println(" Nom : ", *ptr_surname)
}

/*
Afficher les fichiers
*/
func show_existing_lists() {
	cmd := exec.Command("ls")
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(stdout))
}

/*
Merger 2 fichiers
*/
func merge_files(first_file string, second_file string) {
	in, err := os.Open(second_file)
	if err != nil {
		fmt.Println("Impossible d'ouvrir le second fichier:", err)
	}
	defer in.Close()

	out, err := os.OpenFile(first_file, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Impossible d'ouvrir le premier fichier:", err)
	}
	defer out.Close()
}

/*
Fonction MAIN
*/

func main() {
	//fmt.Println("*** h3ll0 h@ck3r! ***\n")

	app := "figlet"

	arg0 := "h3ll0"
	arg1 := "\nH@ck3r !"

	cmd := exec.Command(app, arg0, arg1)
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(stdout))

	var firstname string
	var surname string

	var ptr_firstname *string = &firstname
	var ptr_surname *string = &surname

	options_menu()

	for {
		fmt.Print("\nChoisis une option :\n> ")

		var choix string
		fmt.Scanln(&choix)

		switch choix {
		case "1":
			create_name_list(ptr_firstname, ptr_surname)

		case "2":
			create_name_and_year_list(ptr_firstname, ptr_surname)

		case "3":
			create_name_and_number_list(ptr_firstname, ptr_surname)

		case "4":
			create_date_list()

		case "a":
			fmt.Println("en dev...")
			create_name_list(ptr_firstname, ptr_surname)
			create_name_and_year_list(ptr_firstname, ptr_surname)
			create_name_and_number_list(ptr_firstname, ptr_surname)
			create_date_list()

		case "fa":
			fmt.Println("en dev...")

		case "ls":
			show_existing_lists()

		case "m":
			/*
				var first_file string
				var second_file string
			*/

			fmt.Println("en dev...")
			/*
							fmt.Println("Entre le nom du premier fichier :")
			  				fmt.Scanln(&first_file)
			  				fmt.Println("Entre le nom du second fichier :")
			  				fmt.Scanln(&second_file)
							merge_files(first_file, second_file)
			*/

		case "d":
			delete_list()

		case "h":
			options_menu()

		case "p":
			show_parameters(ptr_firstname, ptr_surname)

		case "q":
			fmt.Println("Au revoir !")
			os.Exit(0)

		default:
			fmt.Println("Commande inconnue, h pour afficher les options")
		}
	}
}

/* TODO
merge -> entre les noms des listes qu'on veut fusionner pour une seule big liste

quand 2 fois la même option, pas écraser -> compteur avec nb d'itération qui change le nom du file

automatiser la majuscule lors de la saisie user du nom/prénom (strings.ToUpper un truc du style)
...
*/
