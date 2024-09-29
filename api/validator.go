package api

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"regexp"
)

type ValidateItem struct {
	key   string
	value any
	tag   string
}

func V(key string, value any, tag string) *ValidateItem {
	return &ValidateItem{
		key:   key,
		value: value,
		tag:   tag,
	}
}

func ascii(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[\x20-\x7E]+$`)
	return re.MatchString(fl.Field().String())
}

func Validate(rules ...*ValidateItem) error {
	v := validator.New()
	_ = v.RegisterValidation("ascii", ascii)
	errMsg := make(map[string]string)
	for _, rule := range rules {
		err := v.Var(rule.value, rule.tag)
		if err != nil {
			var verr validator.ValidationErrors
			errors.As(err, &verr)
			errMsg[rule.key] = verr[0].Tag()
		}
	}
	if len(errMsg) == 0 {
		return nil
	}
	return NewFormValidationError(errMsg)
}
