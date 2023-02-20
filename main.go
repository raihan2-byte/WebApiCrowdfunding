package main

import (
	"BWA/auth"
	"BWA/handler"
	"BWA/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	//Karena menggunakan gorm sehingga dibutuhkan pada line 15-21 yang bertujuan untuk menghubungkan ke database
	dsn := "root:@tcp(127.0.0.1:3306)/startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}

	//menginisiasi userRepository karena kita ingin mengambil repository yang akan dimasukkan kedalam database yang tampak dibawah parameternya "db"
	userRepository := user.NewRepository(db)
	//menginisiasi userService karena kita ingin juga butuh userRepository agar disimpan dalam database dan service juga bertujuan dalam menyimpan input
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	//menginisiasi userHandler yang mana memanggil handler agar kita bisa memasukan register dipostman sesuai method yang membutuhkan userService
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/users/v1")
	api.POST("/user", userHandler.RegisterUser)
	api.POST("/sesions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailabilty)
	api.POST("/avatars", userHandler.UploadAvatar)
	router.Run()

}
