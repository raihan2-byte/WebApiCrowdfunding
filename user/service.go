package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// membuat interface dengan tujuan membuat method
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmaillAvailabilty(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
}

// membuat struct private yang isinya untuk database karena uda di inisiasi pada repository jadi kita ambil repository aja
type service struct {
	repository Repository
}

// membuat func yang isinya pada struct private diatas agar bisa dipanggil di main.go
func NewService(repository Repository) *service {
	return &service{repository}
}

// membuat fungsi yang mana di buat interface diatas bedanya dg repository ini bertujuan sesuai dengan register yaitu pada atribut input
func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	//inisiasi user agar input.go tuh tahu kalau dia digunakan sesuai atribut pada entity.go
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Ocuppation
	//bertujuan agar password yg diinput si user menjadi hash
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password_hash = string(passwordHash)
	user.Role = "user"

	//Menginisiasi newUser karena biar bisa menginput atribut pada input.go yang sesuai dengan repository
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil

}

// membuat fungsi login bertujuan dalam login
func (s *service) Login(input LoginInput) (User, error) {
	//inisiasi email dan password karena butuh nya dalam login hanya hal itu
	email := input.Email
	password := input.Password

	//pengambilan algoritmanya repository yaitu findbyemail
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	//ini jika user id nya 0 maka email itu ga ada karena id 0 kan ga ada
	if user.Id == 0 {
		return user, errors.New("User Not Found That Email")
	}

	//ini buat password menjadi hash karena sebelumnya jg password dibuat hash sehingaa jika kita ga buat fungsi hash seperti dibawah
	//ga akan bisa
	err = bcrypt.CompareHashAndPassword([]byte(user.Password_hash), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

// IsEmaillAvailabilty yang bertujuan apakah email yg didaftarkan tersedia atau tidak
func (s *service) IsEmaillAvailabilty(input CheckEmailInput) (bool, error) {
	//karena hanya email maka di inisiasi hanya email
	email := input.Email

	//pengambilan algoritmanya repository yaitu findbyemail
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	// ini nilainya true karena misal kita input email ini sama ga dengan email yang terdaftar dg id sekian
	//kalau g ada maka balikkanya 0 sehingga bisa di daftrakan atau availabilty
	if user.Id == 0 {
		return true, nil
	}

	return false, nil
}

// membuat fungsi saveAvatar
func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	//mengambil repository findbyid karena id mana ni yang mau upload avatar
	user, err := s.repository.FindById(ID)

	if err != nil {
		return user, err
	}

	//lalu ini adalah nama filenya apa disimpan dalam parameter
	user.Avatar_file_name = fileLocation

	//ini ambil dalam algonya repository kalau mau diupdate
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}
