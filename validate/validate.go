package validate

import (
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/vaidik-bajpai/testimonials/storer"
)

func RegisterValidators() bool {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("atLeastOne", atLeastOne)
		v.RegisterValidation("atLeastName", atLeastName)
		v.RegisterValidation("typesAllowed", typesAllowed)
		return ok
	}
	return false
}

var atLeastOne validator.Func = func(fl validator.FieldLevel) bool {
	questions, ok := fl.Field().Interface().([]string)
	if !ok {
		return false
	}

	length := len(questions)
	if length > 5 {
		return false
	}

	for _, question := range questions {
		if question == "" {
			return false
		}
	}

	return length > 0
}

var atLeastName validator.Func = func(fl validator.FieldLevel) bool {
	userInfos, ok := fl.Field().Interface().([]storer.Field)
	if !ok {
		return false
	}

	length := len(userInfos)
	if length <= 0 || length > 8 {
		return false
	}

	checkDup := make(map[string]bool)

	for _, fieldInfo := range userInfos {
		fieldName := strings.ToLower(fieldInfo.FieldName)
		if !checkDup[fieldName] {
			checkDup[fieldName] = true
		} else {
			return false
		}
	}

	uniqueLen := len(checkDup)

	return (uniqueLen == length) && checkDup["name"]
}

var typesAllowed validator.Func = func(fl validator.FieldLevel) bool {
	spaceType, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	return spaceType == "text" || spaceType == "video" || spaceType == "text&video"
}
