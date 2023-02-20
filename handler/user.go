package handler

import (
	"BWA/auth"
	"BWA/helper"
	"BWA/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// membuat struct private yang isinya user.service karena service isi dari interface yang akan kita gunakan
type userHandler struct {
	userService user.Service
	authService auth.Service
}

// membuat fungsi yang parameter sesuai struct diatas dan mengembalikan nilai nama struct yang isinya variablenya
func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

// membuat fungsi RegisterUser yang ada pada service sebelumnya
func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		//inisiasi erors yang mengarah dalam package helper ada fungsi FormatValidationError yang isinya apakah ada eror atau tidak
		errors := helper.FormatValidationError(err)
		//inisiasi erorMessage yang mana isi erornya apa
		errorMessage := gin.H{"errors": errors}
		//inisiasi response yang mengarah ke package helper.APIresponse dimana fungsi itu status yang mana jika gagal menampilkan data dibawah
		response := helper.APIresponse("Account has failed registered", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		//tidak lupa ada return karena biar tidak lanjut ke step selanjutnya bisa dibilang sebagai "break;"
		return
	}

	//menggunakan userService yang bertujuan untuk memberikan respon pada postman
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		//
		response := helper.APIresponse("Account has failed registered", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := h.authService.GenerateToken(newUser.Id)
	if err != nil {
		response := helper.APIresponse("Account has failed registered", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//jika tidak eror maka inisiasi formatter yang mengarah ke package user.FormatterUser yang mana isinya
	//newUser yang uda diinisasi diatas lalu nilai kedua adalah token karena pada funsi formatterUser token belum dibuat
	formatter := user.FormatterUser(newUser, token)
	//inisiasi response yang mengarah ke helper.APIresponse yang bertujuan jika sukses maka mendapatkan respon seperi dibawah
	//dan kembalikan nilai formatter karena formatter merupakan isi response yang ada pada fungsi helper.APIresponse//
	response := helper.APIresponse("Account has been registered", http.StatusOK, "Success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		//inisiasi erors yang mengarah dalam package helper ada fungsi FormatValidationError yang isinya apakah ada eror atau tidak
		errors := helper.FormatValidationError(err)
		//inisiasi erorMessage yang mana isi erornya apa
		errorMessage := gin.H{"errors": errors}
		//inisiasi response yang mengarah ke package helper.APIresponse dimana fungsi itu status yang mana jika gagal menampilkan data dibawah
		response := helper.APIresponse("Login Failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		//tidak lupa ada return karena biar tidak lanjut ke step selanjutnya bisa dibilang sebagai "break;"
		return
	}

	LoggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIresponse("Login Failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(LoggedinUser.Id)
	if err != nil {
		response := helper.APIresponse("Login Failed", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatterUser(LoggedinUser, token)
	response := helper.APIresponse("Login Success", http.StatusOK, "Success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailabilty(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		//inisiasi erors yang mengarah dalam package helper ada fungsi FormatValidationError yang isinya apakah ada eror atau tidak
		errors := helper.FormatValidationError(err)
		//inisiasi erorMessage yang mana isi erornya apa
		errorMessage := gin.H{"errors": errors}
		//inisiasi response yang mengarah ke package helper.APIresponse dimana fungsi itu status yang mana jika gagal menampilkan data dibawah
		response := helper.APIresponse("Email checking failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		//tidak lupa ada return karena biar tidak lanjut ke step selanjutnya bisa dibilang sebagai "break;"
		return
	}

	IsEmailAvailable, err := h.userService.IsEmaillAvailabilty(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Eror"}
		//inisiasi response yang mengarah ke package helper.APIresponse dimana fungsi itu status yang mana jika gagal menampilkan data dibawah
		response := helper.APIresponse("Email checking failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		//tidak lupa ada return karena biar tidak lanjut ke step selanjutnya bisa dibilang sebagai "break;"
		return
	}

	//inisiasi data dengan tujuan mengecek apakah email yang didaftarkan tersedia atau tidak sehingga membutuhkan gin.H
	data := gin.H{
		"is_available": IsEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if IsEmailAvailable {
		metaMessage = "Email is available"
	}
	//kalau berhasil maka akan memberikan data dibawah
	response := helper.APIresponse(metaMessage, http.StatusOK, "Success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	//inisiasi file yang bertujuan untuk menangkap avatar yang akan diupload oleh user
	//formFile dimana kalau dipostman kita masukkan value maka avatar, contoh kalau register email:raihan@gmail.com kalau ini adalah avatar
	file, err := c.FormFile("avatar")
	if err != nil {
		//inisiasi data yang tujuan dalam return hasil ke postman
		data := gin.H{"is_uploaded": false}
		response := helper.APIresponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	userID := 1

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	//ini berfungsi dalam upload file dengan perbedaan diatas ialah punya si meta dan bawah punya si data
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIresponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//pengambilan pada file service
	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIresponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helper.APIresponse("Success to upload avatar image", http.StatusOK, "error", data)
	c.JSON(http.StatusOK, response)

}
