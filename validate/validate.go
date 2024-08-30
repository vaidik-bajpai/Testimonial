package validate

import (
	"fmt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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
	userInfos, ok := fl.Field().Interface().([]string)
	fmt.Println(userInfos)
	hasName := false
	if !ok {
		return false
	}

	if len(userInfos) > 8 {
		return false
	}

	for _, info := range userInfos {
		if len(info) < 4 {
			return false
		}
		if info == "name" {
			hasName = true
		}
	}

	return true && hasName
}

var typesAllowed validator.Func = func(fl validator.FieldLevel) bool {
	spaceType, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	return spaceType == "text" || spaceType == "video" || spaceType == "text&video"
}
