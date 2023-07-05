package main

import (
	"github.com/amr-ahmed/go-crud/initializers"
	"github.com/amr-ahmed/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
