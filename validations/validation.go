package validations

import (
	"regexp"

	"gopkg.in/go-playground/validator.v9"
)

var (
	validade = validator.New()
)

func init() {
	validade.RegisterValidation("passwd", passwd)
}

func passwd(fl validator.FieldLevel) bool {
	re, _ := regexp.Compile("^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[@$!%*?&])[A-Za-z\\d@$!%*?&]{6,}$")
	return re.FindString(fl.Field().String()) != ""
}

func profile(fl validator.FieldLevel) bool {
	return fl.Field().String() != "ADMIN" || fl.Field().String() != "ADMIN"
}

func Get() *validator.Validate {
	return validade
}
