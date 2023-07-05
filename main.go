package main

import (
	"github.com/amr-ahmed/go-crud/controllers"
	"github.com/amr-ahmed/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()
	postController := controllers.PostController{}
	router.POST("/posts", postController.PostsCreate)
	router.GET("/posts", postController.PostsIndex)
	router.GET("/posts/:id", postController.PostsShow)
	router.PUT("/posts/:id", postController.PostsUpdate)
	router.DELETE("/posts/:id", postController.PostsDelete)
	router.Run()
}
