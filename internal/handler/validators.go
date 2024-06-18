package handler

import (
	"reflect"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

func roleDependentFields(fl validator.FieldLevel) bool {
	role := fl.Parent().FieldByName("Role").String()
	headmanID := fl.Parent().FieldByName("HeadmanID").Interface()
	studentID := fl.Parent().FieldByName("StudentID").Interface()
	teacherID := fl.Parent().FieldByName("TeacherID").Interface()

	isFieldSet := func(field interface{}) bool {
		return !reflect.ValueOf(field).IsNil()
	}

	switch role {
	case "Староста":
		return isFieldSet(headmanID) && !isFieldSet(studentID) && !isFieldSet(teacherID)
	case "Студент":
		return !isFieldSet(headmanID) && isFieldSet(studentID) && !isFieldSet(teacherID)
	case "Преподаватель":
		return !isFieldSet(headmanID) && !isFieldSet(studentID) && isFieldSet(teacherID)
	default:
		return !isFieldSet(headmanID) && !isFieldSet(studentID) && !isFieldSet(teacherID)
	}
}

func ValidateTime(fl validator.FieldLevel) bool {
	timeStr := fl.Field().String()
	_, err := time.Parse("15:04", timeStr)
	return err == nil
}

func CustomSpecialtyCodeRegex(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	pattern := `^\d{2}\.\d{2}\.\d{2}$`
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(value)
}

func CustomGroupIDRegex(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	pattern := `^\d{4}-\d{2}\.\d{2}\.\d{2}-\d{1,6}$`
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(value)
}

func CustomPasswordRegex(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	pattern := "^[A-Za-z\\d$!%*#?&@]+$"
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(value)
}

func CustomFieldRusRegex(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	pattern := "^[а-яА-Я\\s-]+$"
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(value)
}
