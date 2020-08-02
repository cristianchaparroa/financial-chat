package handlers

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

type CustomValidator struct {
	Validator  *validator.Validate
	Translator ut.Translator
}

func NewValidator() *CustomValidator {
	v := validator.New()
	translator := en.New()
	uni := ut.New(translator, translator)
	trans, _ := uni.GetTranslator("en")
	v.RegisterTranslation("required", trans, registerField, translationMessage)
	return &CustomValidator{Validator: v, Translator: trans}
}

func registerField(ut ut.Translator) error {
	return ut.Add("required", "{0} is a required field", true)
}

func translationMessage(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("required", ToSnakeCase(fe.Field()))
	return t
}

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
