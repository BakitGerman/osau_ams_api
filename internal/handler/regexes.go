package handler

import "regexp"

const (
	alphaRegexString        = "^[a-zA-Z]+$"
	alphaNumericRegexString = "^[a-zA-Z0-9]+$"
	groupRegexString        = "^\\d{4}-\\d{2}\\.\\d{2}\\.\\d{2}-\\d{1,6}$"
	passwordRegexString     = "^[A-Za-z\\d$!%*#?&@]+$"
	alphaRusRegexString     = "^[А-Яа-я ]+$"
	alphaNumRusRegexString  = "^[А-Яа-я\\d ]+$"
	dateRegexString         = `^\d{4}-\d{2}-\d{2}$`
)

var (
	alphaRegex        = regexp.MustCompile(alphaRegexString)
	alphaNumericRegex = regexp.MustCompile(alphaNumericRegexString)
	groupRegex        = regexp.MustCompile(groupRegexString)
	passwordRegex     = regexp.MustCompile(passwordRegexString)
	alphaRusRegex     = regexp.MustCompile(alphaRusRegexString)
	alphaNumRusRegex  = regexp.MustCompile(alphaNumRusRegexString)
	dateRegex         = regexp.MustCompile(dateRegexString)
)
