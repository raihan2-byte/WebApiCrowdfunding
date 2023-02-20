package user

import (
	"gorm.io/gorm"
)

// membuat interface dengan tujuan membuat method
type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindById(ID int) (User, error)
	Update(user User) (User, error)
}

// membuat struct private yang isinya database
type repository struct {
	db *gorm.DB
}

// membuat func yang isinya pada struct private diatas agar bisa dipanggil di main.go
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// membuat fungsi yang mana di buat pada interface diatas yang mana ada create yaitu membuat user dan disimpan pada database
func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil

}

// findbyemail yanbg bertujuan dalam login
func (r *repository) FindByEmail(email string) (User, error) {
	//inisiasi user yang bertujuan dalam mencari email yang ada di struct User
	var user User

	//menggunakan where yang mana email mana ni yang mau di loginin sesuai ga
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// FindById yang bertujuan mencari ID dimana bisa berfungsi dalam apakah email ini tersedia atau tidak dan lain lain
func (r *repository) FindById(ID int) (User, error) {
	var user User

	//sama halnya dengan findbyemail id nya brapa
	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// update yang bertujuan dalam update avatar dan lain lain
func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil

}
