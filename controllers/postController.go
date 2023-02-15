package controllers

import (
	"log"

	"github.com/G-MURAKARU/go-simple-crud/initialisers"
	"github.com/G-MURAKARU/go-simple-crud/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(ctx *gin.Context) {
	// get data off request body - below struct should match request body
	var requestBody struct {
		Title   string
		Content string
	}

	// binding the gin context to the post body, to extract incoming data
	ctx.Bind(&requestBody)

	// create post from the request body
	post := models.Post{
		Title:   requestBody.Title,
		Content: requestBody.Content,
	}

	// pass a pointer to the data that should populate the db
	result := initialisers.DB.Create(&post)
	if result.Error != nil {
		ctx.Status(400)
	}

	// return created post (hover over JSON for info)
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func PostReadAll(ctx *gin.Context) {
	// get all posts
	var posts []models.Post
	result := initialisers.DB.Find(&posts)

	if result.Error != nil {
		log.Fatal("Could not retrieve all posts")
	}

	// return the posts
	ctx.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostReadOne(ctx *gin.Context) {
	// posts will be retrieved using their id

	// get id from url
	id := ctx.Param("id")

	// query the database for the post
	var post models.Post
	result := initialisers.DB.Find(&post, id)
	if result.Error != nil {
		log.Fatal("Could not find the queried post.")
	}

	// return the post
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(ctx *gin.Context) {
	// posts will be retrieved using their id

	// get id from url
	id := ctx.Param("id")

	// get data off request body
	var requestBody struct {
		Title   string
		Content string
	}

	// binding the gin context to the post body, to extract incoming data
	ctx.Bind(&requestBody)

	// query the database for the post
	var post models.Post
	result := initialisers.DB.Find(&post, id)
	if result.Error != nil {
		log.Fatal("Could not find the queried post.")
	}

	// update the post
	initialisers.DB.Model(&post).Updates(models.Post{
		Title:   requestBody.Title,
		Content: requestBody.Content,
	})

	// return the updated post
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(ctx *gin.Context) {
	// get id from url
	id := ctx.Param("id")

	// delete the post
	/*
		since the post model has DeletedAt, it will get a soft delete
		i.e. the record will not be deleted from the database
		the DeletedAt field will be populated at time of delete
		the associated entry will not be findable with normal queries
	*/
	initialisers.DB.Delete(&models.Post{}, id)

	// respond appropriately
	ctx.Status(204)
}
