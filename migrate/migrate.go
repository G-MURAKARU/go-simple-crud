package main

import (
	"github.com/G-MURAKARU/go-simple-crud/initialisers"
	"github.com/G-MURAKARU/go-simple-crud/models"
)

func init() {
	initialisers.LoadEnvVariables()
	initialisers.ConnectToDB()
}

func main() {
	initialisers.DB.AutoMigrate(&models.Post{})
}
