package user

import "time"

//sesuai dengan atribut semua yang ada pada profile
type User struct {
	Id               int
	Name             string
	Occupation       string
	Email            string
	Password_hash    string
	Avatar_file_name string
	Role             string
	Created_at       time.Time
	Update_at        time.Time
}
