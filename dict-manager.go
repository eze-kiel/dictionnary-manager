package main

import "fmt"
import "os"
import "strings"
import "strconv"
import "os/exec"

/*
Création d'un .txt avec le nom et le prénom
*/
func create_name_list() {
	var firstname string
	var surname string

	fmt.Println("Attention : Mets la majuscule au début de chaque mot !")
	fmt.Println(("Prénom :"))
	fmt.Scanln(&firstname)
	fmt.Println(("\nNom :"))
	fmt.Scanln(&surname)

	//créer le fichier
	f, err := os.Create("name.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//écrire dedans
	f.WriteString(firstname + surname + "\n")
	f.WriteString(surname + firstname + "\n")
	f.WriteString(strings.ToUpper(firstname) + strings.ToUpper(surname) + "\n")
	f.WriteString(strings.ToUpper(surname) + strings.ToUpper(firstname) + "\n")
	f.WriteString(strings.ToLower(firstname) + strings.ToLower(surname) + "\n")
	f.WriteString(strings.ToLower(surname) + strings.ToLower(firstname))

	f.Close()
	fmt.Println("Fait !")
}


/*
Création d'un .txt avec le nom, prénom et année
*/
func create_name_and_year_list() {
	var firstname string
	var surname string
	var year string

	fmt.Println("Attention : Mets la majuscule au début de chaque mot !")
	fmt.Println(("Prénom :"))
	fmt.Scanln(&firstname)
	fmt.Println(("\nNom :"))
	fmt.Scanln(&surname)
	fmt.Println(("\nAnnée :"))
	fmt.Scanln(&year)

	//créer le fichier
	f, err := os.Create("nameyear.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//écrire dedans
	f.WriteString(firstname + surname + year + "\n")
	f.WriteString(surname + firstname + year + "\n")
	f.WriteString(strings.ToUpper(firstname) + strings.ToUpper(surname) + year + "\n")
	f.WriteString(strings.ToUpper(surname) + strings.ToUpper(firstname) + year + "\n")
	f.WriteString(strings.ToLower(firstname) + strings.ToLower(surname) + year + "\n")
	f.WriteString(strings.ToLower(surname) + strings.ToLower(firstname) + year)

	f.Close()
	fmt.Println("Fait !")
}

/*
Création d'un .txt avec le nom, prénom et nombres
*/
func create_name_and_number_list() {
  	var firstname string
	var surname string
	var min int
  	var max int

	fmt.Println("Attention : Mets la majuscule au début de chaque mot !")
	fmt.Println(("Prénom :"))
	fmt.Scanln(&firstname)
	fmt.Println(("\nNom :"))
	fmt.Scanln(&surname)

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
    	f.WriteString(firstname + surname + strconv.FormatInt(int64(i), 10) + "\n")
    	f.WriteString(surname + firstname + strconv.FormatInt(int64(i), 10) + "\n")
	  	f.WriteString(strings.ToUpper(firstname) + strings.ToUpper(surname) + strconv.FormatInt(int64(i), 10) + "\n")
	  	f.WriteString(strings.ToUpper(surname) + strings.ToUpper(firstname) + strconv.FormatInt(int64(i), 10) + "\n")
	  	f.WriteString(strings.ToLower(firstname) + strings.ToLower(surname) + strconv.FormatInt(int64(i), 10) + "\n")
	  	f.WriteString(strings.ToLower(surname) + strings.ToLower(firstname) + strconv.FormatInt(int64(i), 10) + "\n")
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
						f.WriteString("0"+strconv.FormatInt(int64(d), 10)+"0"+strconv.FormatInt(int64(m), 10)+strconv.FormatInt(int64(y), 10)+"\n")
						f.WriteString(strconv.FormatInt(int64(y), 10)+"0"+strconv.FormatInt(int64(m), 10)+"0"+strconv.FormatInt(int64(d), 10)+"\n")
					} else {
						f.WriteString(strconv.FormatInt(int64(d), 10)+"0"+strconv.FormatInt(int64(m), 10)+strconv.FormatInt(int64(y), 10)+"\n")
						f.WriteString(strconv.FormatInt(int64(y), 10)+"0"+strconv.FormatInt(int64(m), 10)+strconv.FormatInt(int64(d), 10)+"\n")
					}
				} else if d < 10 {
					f.WriteString("0"+strconv.FormatInt(int64(d), 10)+strconv.FormatInt(int64(m), 10)+strconv.FormatInt(int64(y), 10)+"\n")
					f.WriteString(strconv.FormatInt(int64(y), 10)+strconv.FormatInt(int64(m), 10)+"0"+strconv.FormatInt(int64(d), 10)+"\n")
				} else {
					f.WriteString(strconv.FormatInt(int64(d), 10)+strconv.FormatInt(int64(m), 10)+strconv.FormatInt(int64(y), 10)+"\n")
					f.WriteString(strconv.FormatInt(int64(y), 10)+strconv.FormatInt(int64(m), 10)+strconv.FormatInt(int64(d), 10)+"\n")
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

  	if strings.Contains(name, "/") == true {
    	fmt.Println("Le fichier doit être dans le même répertoire que ce script !")
  	} else if strings.Contains(name, ".go") == true {
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
	fmt.Println(" [q] : quitter")
}

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

Fonction MAIN

*/

func main() {
	//fmt.Println("*** 1337 h@x3r i5 b@ck! ***\n")

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

	options_menu()

	for {
		fmt.Print("\nChoisis une option :\n> ")

		var choix string
		fmt.Scanln(&choix)

		switch choix {
			case "1":
				create_name_list()

			case "2":
				create_name_and_year_list()

	    	case "3":
	      		create_name_and_number_list()

	      	case "4":
	      		create_date_list()

			case "a":
				fmt.Println("en dev...")
			
			case "fa":
				fmt.Println("en dev...")

			case "ls":
				show_existing_lists()

			case "m":
				fmt.Println("en dev...")

	    	case "d":
	      		delete_list()

			case "h":
				options_menu()

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
...
*/