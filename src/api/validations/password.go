package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/hossein-repo/myproj1/common"
)

func PasswordValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		fld.Param()
		return false
	}

	return common.CheckPassword(value)
}
