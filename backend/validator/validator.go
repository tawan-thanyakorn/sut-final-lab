package validator

import "github.com/asaskevich/govalidator"

func ValidateStruct(v interface{}) error {
	_, err := govalidator.ValidateStruct(v)
	return err
}