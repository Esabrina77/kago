package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//1. Definition du flag 'type (default : "simple")

	var projectType = flag.String("type", "simple", "type of project")

	//2. Analyse
	flag.Parse()

	//3. Recuperation du nom du projet (argument restant)
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("You must provide a project name")
		fmt.Println("Usage : kago -type= simple <project-name>")
		os.Exit(1)
	}

	//4. Affichage  pour le debug

	fmt.Printf("Initialisation du projet .... \n")
	fmt.Printf("Nom du projet : %s \n", args[0])
	fmt.Printf("Type de projet : %s \n", *projectType)
}
