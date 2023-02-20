//tujuannya dibuat helper ialah karena dalam form registrasi jika sudah regis data mendapatkan hasil apakah sukses

package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"` //interface karena hasilnya tidak pasti sebab ga semua memasuki inputan name dll
}

//membuat struct bertujuan apa output kalauberhasil atau tidak
type Meta struct {
	Massage string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

//buat fungsi yang digunakan pada handler agar masuk ke dalam web
func APIresponse(massage string, code int, status string, data interface{}) Response {
	//buat metanya yaitu apasi isi dari meta itu yaitu ada status dll
	meta := Meta{
		Massage: massage,
		Code:    code,
		Status:  status,
	}
	//buat response nya ada meta yg di inisiasi diatas sama data
	JsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return JsonResponse
}

//membuat fungsi formatvalidator yang bertujuan dalam melihat apakah eror nya dimana, dengan kode apa sehingga
//membutuhkan parameter error yang mengembalikan nilai string berbentuk array
func FormatValidationError(err error) []string {
	//inisiasi errors aray string karena berguna dalam
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
