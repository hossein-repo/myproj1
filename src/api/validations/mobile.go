package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/hossein-repo/myproj1/common"
)




func IranianMobileNumberValidat(fld validator.FieldLevel) bool {

	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}
    return common.IranianMobileNumberValidat(value)

}