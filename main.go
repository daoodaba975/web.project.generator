package main

import (
	"fmt"
	"os"
)

// _____________________________________________________________
//
//  Déclarations des variables
// _____________________________________________________________
var lessCompiler bool
var answer string
var wait string
var easyLessConfig string
var lesspath string
var projectname string

// _____________________________________________________________
//
//  La fonction main est auto-exécutée par Go.
// _____________________________________________________________
func main() {

	// _____________________________________________________________
	//
	//  Le logiciel vous demande le nom du projet.
	// _____________________________________________________________
	fmt.Println("Quel est le nom du projet ?")
	fmt.Scanln(&projectname)
	if projectname == "" {
		projectname = "new-project"
	}
	fmt.Println("")

	// _____________________________________________________________
	//
	//  Sélectionnez "Y" si vous souhaitez l’utiliser ou sélectionnez "N" si vous ne l’utilisez pas.
	// _____________________________________________________________
	easyLessConfig = "// main:styles.less, out: ../css/styles.min.css, compress: true, sourceMap: false"

	fmt.Println("Souhaitez-vous utiliser LESS dans ce projet ? (Y/N)")
	fmt.Scanln(&answer)
	if answer == "Y" || answer == "y" {
		// _____________________________________________________________
		//
		//  Création des dossiers: public,css,images et js
		// _____________________________________________________________
		createFolder(projectname + "/public/css")
		createFolder(projectname + "/public/images")
		createFolder(projectname + "/public/js")
		createFolder(projectname + "/public/less")

		// _____________________________________________________________
		//
		//  Création des fichiers LESS
		// _____________________________________________________________
		lesspath = projectname + "/public/less/"
		createVariablesDocument()
		createResetDocument()
		createLessDocuments()
		createStylesDocument()
	} else {
		// _____________________________________________________________
		//
		//  Création des dossiers: public,css,images et js
		// _____________________________________________________________
		createFolder(projectname + "/public/css")
		createFolder(projectname + "/public/images")
		createFolder(projectname + "/public/js")
	}

	// _____________________________________________________________
	//
	//  Message de réussite
	// _____________________________________________________________
	fmt.Printf("\nProjet créé avec succès.\nAppuyez sur ENTER pour sortir")

	// _____________________________________________________________
	//
	//  Appuyez sur ENTER pour quitter
	// _____________________________________________________________
	fmt.Scanln(&wait)
}

// _____________________________________________________________
//
//  C'est une fonction générique pour créer les sous-dossiers
// _____________________________________________________________
func createFolder(pathname string) {
	os.MkdirAll(pathname, os.ModePerm)
}

// _____________________________________________________________
//
//  Cette fonction crée les @import dans le fichier styles.less
// _____________________________________________________________
func createStylesDocument() {
	file, err := os.Create(lesspath + "styles.less")
	if err != nil {
		fmt.Println("Problème avec la création de styles.less")
		return
	}
	if lessCompiler == true {
		fmt.Fprintln(file, easyLessConfig)
	}
	fmt.Fprintln(file, "@import url('reset.less');")
	fmt.Fprintln(file, "@import url('variables.less');")
	fmt.Fprintln(file, "@import url('functions.less');")
	fmt.Fprintln(file, "@import url('main.less');")
	fmt.Fprintln(file, "@import url('header.less');")
	fmt.Fprintln(file, "@import url('aside.less');")
	fmt.Fprintln(file, "@import url('navbar.less');")
	fmt.Fprintln(file, "@import url('articles.less');")
	fmt.Fprintln(file, "@import url('footer.less');")
	fmt.Fprintln(file, "@import url('forms.less');")
	fmt.Fprintln(file, "@import url('buttons.less');")
	fmt.Fprintln(file, "@import url('tables.less');")
	fmt.Fprintln(file, "@import url('helpers.less');")
	fmt.Fprintln(file, "@import url('responsive.less');")
	fmt.Fprintln(file, "@import url('animations.less');")
	defer file.Close()
}

// _____________________________________________________________
//
//  Cette fonction prédéfini le contenu dans variables.less
// _____________________________________________________________
func createVariablesDocument() {

	selectors := [9]string{"@imports", "background", "border", "colors", "font", "height", "margin", "padding", "width"}

	file, err := os.Create(lesspath + "variables.less")
	if err != nil {
		fmt.Println("Problème avec la création de variables.less")
		return
	}
	if lessCompiler == true {
		fmt.Fprintln(file, easyLessConfig)
	}
	for i := 0; i < 9; i++ {
		fmt.Fprintln(file, "/* =============================================================================")
		fmt.Fprintln(file, "* \t", selectors[i])
		fmt.Fprintln(file, "*  ========================================================================== */")
		fmt.Fprintln(file, "")
	}
	defer file.Close()
}

// _____________________________________________________________
//
//  Cette fonction prédéfini le contenu dans reset.less
// _____________________________________________________________
func createResetDocument() {
	file, err := os.Create(lesspath + "reset.less")
	if err != nil {
		fmt.Println("Problème avec la création de reset.Less")
		return
	}
	if lessCompiler == true {
		fmt.Fprintln(file, easyLessConfig)
	}
	fmt.Fprintln(file, "* {")
	fmt.Fprintln(file, "\t margin: 0;")
	fmt.Fprintln(file, "\t border: 0;")
	fmt.Fprintln(file, "\t padding: 0;")
	fmt.Fprintln(file, "\t outline: none;")
	fmt.Fprintln(file, "}")
	defer file.Close()
}

// _____________________________________________________________
//
//  Cette fonction crée et vérifie la bonne création des fichiers LESS
//	S'il y a erreur il alerte l'utilisateur
// _____________________________________________________________
func createLessDocuments() {

	files := [13]string{"functions.less", "header.less", "navbar.less", "main.less", "aside.less", "articles.less", "footer.less", "forms.less", "buttons.less", "tables.less", "helpers.less", "responsive.less", "animations.less"}

	for i := 0; i < len(files); i++ {
		file, err := os.Create(lesspath + files[i])
		if err != nil {
			fmt.Printf("Problème avec la création de %v", files[i])
			panic(err)
		}
		if lessCompiler == true {
			fmt.Fprintln(file, easyLessConfig)
		}
		defer file.Close()
	}

}
