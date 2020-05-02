package conf

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate
var validateOnce sync.Once

func GetValidate() *validator.Validate {
	validateOnce.Do(func() {
		validate = validator.New()
	})
	return validate
}
