package controllers

import (
	"github.com/amr-ahmed/go-crud/models"
	"github.com/amr-ahmed/go-crud/services/posts"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	posts.PostService
}

func NewPostController() *PostController {
	controller := PostController{}
	return &controller
}

func (controller *PostController) PostsCreate(ctx *gin.Context) {
	//get data from request and cast body into post struct
	post := models.Post{}
	ctx.Bind(&post)
	//store the new post in db
	status, res := controller.PostService.PostCreate(&post)
	//send the post in response if there is no error
	ctx.JSON(status, res)
}

func (controller *PostController) PostsIndex(ctx *gin.Context) {
	//get all posts
	status, res := controller.PostService.PostsIndex()
	ctx.JSON(status, gin.H{
		"response": res,
	})
}

func (controller *PostController) PostsShow(ctx *gin.Context) {
	//get id off url
	id := ctx.Param("id")
	//get post based on id
	status, res := controller.PostService.PostsShow(id)
	ctx.JSON(status, gin.H{
		"response": res,
	})
}

func (controller *PostController) PostsUpdate(ctx *gin.Context) {
	//get the id off url
	id := ctx.Param("id")
	//cast updated values into post struct
	postWithNewValues := models.Post{}
	ctx.Bind(&postWithNewValues)
	status, res := controller.PostService.PostsUpdate(id, &postWithNewValues)
	ctx.JSON(status, gin.H{
		"response": res,
	})
}

func (controller *PostController) PostsDelete(ctx *gin.Context) {
	//get the id off url
	id := ctx.Param("id")
	//delete the post from database based on id
	status, res := controller.PostService.PostsDelete(id)
	ctx.String(status, res)
}
