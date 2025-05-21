package main

import (
	"fmt"

	"github.com/suhailkassar11/image_processing_system/initializers"
	"github.com/suhailkassar11/image_processing_system/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
}

func main() {
	initializers.DB.AutoMigrate(
		&models.User{},
	)
	fmt.Println("migrated")
}
