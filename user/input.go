package user

//sesuai dengan atribut yang ada pada isi register
type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Ocuppation string `json:"ocuppation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
