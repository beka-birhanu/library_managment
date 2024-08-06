package main

import (
	"os"

	"github.com/beka-birhanu/library_managment/controllers"
	"github.com/beka-birhanu/library_managment/services"
)

func main() {
	libraryService := services.NewLibrary()
	console := controllers.NewConsole(libraryService, os.Stdin, os.Stdout)

	console.Run()
}
