package posts

import (
	"fmt"
	"net/http"

	"github.com/amr-ahmed/go-crud/initializers"
	"github.com/amr-ahmed/go-crud/models"
)

type PostService struct {
}

func NewPostService() *PostService {
	service := PostService{}
	return &service
}

func (service *PostService) PostCreate(post *models.Post) (int, map[string]interface{}) {
	response := map[string]interface{}{
		"message": "post has been stored successfully.",
	}
	if result := initializers.DB.Create(&post); result.Error != nil {
		response["message"] = "Can not create new post."
		return http.StatusBadRequest, response
	}
	response["post"] = &post

	return http.StatusOK, response
}

func (service *PostService) PostsIndex() (int, map[string]interface{}) {
	posts := []models.Post{}
	response := map[string]interface{}{
		"message": "posts have been fetched successfully.",
	}
	if result := initializers.DB.Find(&posts); result.Error != nil {
		response["message"] = "can not fetch posts."
		return http.StatusBadRequest, response
	}
	response["posts"] = &posts

	return http.StatusOK, response

}

func (service *PostService) PostsShow(id string) (int, map[string]interface{}) {
	post := models.Post{}
	response := map[string]interface{}{
		"message": "post has been fetched successfully.",
	}
	if result := initializers.DB.First(&post, id); result.Error != nil {
		response["message"] = "post not found."
		return http.StatusNotFound, response
	}
	response["post"] = &post
	return http.StatusOK, response
}

func (service *PostService) PostsUpdate(id string, postWithNewValues *models.Post) (int, map[string]interface{}) {
	//cast updated value into post struct
	response := map[string]interface{}{
		"message": "post has been updated successfully",
	}
	//find the post based on id
	postWithDBValues := models.Post{}
	if find := initializers.DB.First(&postWithDBValues, id); find.Error != nil {
		response["message"] = "post not found."
		return http.StatusNotFound, response
	}
	//update the post in db based on id
	if update := initializers.DB.Model(&postWithDBValues).Updates(&postWithNewValues); update.Error != nil {
		response["message"] = "post has not been updated."
		return http.StatusBadRequest, response
	}
	//send the post with response if there is no error
	response["post"] = &postWithDBValues
	return http.StatusOK, response
}

func (service *PostService) PostsDelete(id string) (int, string) {
	response := fmt.Sprintf("post with id %s has been deleted successfully.", id)
	if find := initializers.DB.First(&models.Post{}, id); find.Error != nil {
		response = fmt.Sprintf("post with id %s not found.", id)
		return http.StatusNotFound, response
	}
	if delete := initializers.DB.Delete(&models.Post{}, id); delete.Error != nil {
		response = "post has not been deleted."
		return http.StatusBadRequest, response
	}
	return http.StatusOK, response
}
