package godash

import (
	"regexp"
	"strings"
)

type InitCamelCase bool

const (
	UpperCase InitCamelCase = true
	LowerCase InitCamelCase = false
)

func CamelCase(str string) (string, error) {
	return CamelCaseWithInit(str, false)
}

func CamelCaseWithInit(str string, upperCase InitCamelCase) (string, error) {
	initCamelCaseFunc := Ternary(bool(upperCase), strings.ToUpper, strings.ToLower).(func(string) string)

	str, err := replaceWithRegexFunc(`\s(.)`, str, strings.ToUpper)
	if err != nil {
		return "", err
	}

	if str, err = replaceWithRegexFunc(`([\d.,\-_])[a-zA-Z]`, str, strings.ToUpper); err != nil {
		return "", err
	}

	if str, err = replaceWithRegex(`\s`, str, ""); err != nil {
		return "", err
	}

	if str, err = replaceWithRegexFunc(`^(.)`, str, initCamelCaseFunc); err != nil {
		return "", err
	}

	return str, nil
}

func replaceWithRegex(expression string, src string, repl string) (string, error) {
	r, err := regexp.Compile(expression)
	if err != nil {
		return "", err
	}

	result := src
	result = r.ReplaceAllString(result, repl)
	return result, nil
}

func replaceWithRegexFunc(expression string, src string, repl func(string) string) (string, error) {
	r, err := regexp.Compile(expression)
	if err != nil {
		return "", err
	}

	result := src
	result = r.ReplaceAllStringFunc(result, repl)
	return result, nil
}
