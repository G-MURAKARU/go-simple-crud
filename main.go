package main

import (
	"github.com/G-MURAKARU/go-simple-crud/controllers"
	"github.com/G-MURAKARU/go-simple-crud/initialisers"
	"github.com/gin-gonic/gin"
)

func init() {
	// init will run before main

	// script to load our environment variables using godotenv
	initialisers.LoadEnvVariables()
	initialisers.ConnectToDB()
}

func main() {
	router := gin.Default()
	router.POST("/posts", controllers.PostCreate)
	router.GET("/posts", controllers.PostReadAll)
	router.GET("/posts/:id", controllers.PostReadOne)
	router.PUT("/posts/:id", controllers.PostUpdate)
	router.DELETE("/posts/:id", controllers.PostDelete)
	router.Run() // listen and serve on 0.0.0.0:8080
}
