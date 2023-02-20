package user

//karena kita membutuhkan output berhasil atau tidak dalam mengisi data register sehingga membutuhkan struct baru
type UsetFormatter struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

//ini biar nilai struct diatas dipakai sehingga dibuat func yang berisi nilai struc diatas
func FormatterUser(user User, Token string) UsetFormatter {
	formatter := UsetFormatter{
		Id:         user.Id,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      Token,
	}

	return formatter

}
