package godash

import (
	"html"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type InitCamelCase bool

// Converts string to camel case. First char is lower case.
func CamelCase(str string) (string, error) {
	return CamelCaseWithInit(str, false)
}

// Converts string to camel case. First char is lower case by default.
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

// Converts the first character of string to upper case and the remaining to lower case.
func Capitalize(str string) string {
	// return strings.Title(str)
	return cases.Title(language.Und, cases.NoLower).String(str)
}

// Checks if string ends with the given target string.
func EndsWith(str string, target string) bool {
	return strings.HasSuffix(str, target)
}

// Checks if string ends with the given target string with the position to search up to.
func EndsWithFrom(str string, target string, position int) bool {
	strLength := len(str)
	if position > strLength {
		position = strLength
	}

	str = str[0:position]
	return EndsWith(str, target)
}

// Converts the characters "&", "<", ">", '"', and "'" in string to their corresponding HTML entities.
func Escape(str string) string {
	return html.EscapeString(str)
}

// Escapes the RegExp special characters "^", "$", "", ".", "*", "+", "?", "(", ")", "[", "]", "{", "}", and "|" in string.
func EscapeRegExp(str string) string {
	return regexp.QuoteMeta(str)
}

func LowerFirst(str string) string {
	if len(str) > 1 {
		return strings.ToLower(str[0:1]) + str[1:]
	} else {
		return strings.ToLower(str)
	}
}

func padWithPosition(str string, length int, padChars string, position int) string {
	if len(str) >= length {
		return str
	}

	if padChars == "" {
		padChars = " "
	}

	length = length - len(str)
	leftPadLen := 0
	if position == 0 {
		leftPadLen = length / 2
	} else if position == 1 {
		leftPadLen = length
	}
	rightPadLen := length - leftPadLen

	charLen := len(padChars)
	leftPad := ""
	cur := 0
	for cur < leftPadLen {
		leftPad += string(padChars[cur%charLen])
		cur++
	}

	cur = 0
	rightPad := ""
	for cur < rightPadLen {
		rightPad += string(padChars[cur%charLen])
		cur++
	}

	return leftPad + str + rightPad
}

// Pads string on the left and right sides if it's shorter than length. Padding characters are truncated if they can't be evenly divided by length.
func PadWith(str string, length int, padChars string) string {
	return padWithPosition(str, length, padChars, 0)
}

// Pads string on the left sides if it's shorter than length.
func PadLeftWith(str string, length int, padChars string) string {
	return padWithPosition(str, length, padChars, 1)
}

// Pads string on the right sides if it's shorter than length.
func PadRightWith(str string, length int, padChars string) string {
	return padWithPosition(str, length, padChars, 2)
}

// Pads string on the left and right sides if it's shorter than length. Padding characters are truncated if they can't be evenly divided by length.
func Pad(str string, length int) string {
	return PadWith(str, length, " ")
}

// Pads string on the left sides if it's shorter than length.
func PadLeft(str string, length int) string {
	return PadLeftWith(str, length, " ")
}

// Pads string on the right sides if it's shorter than length.
func PadRight(str string, length int) string {
	return PadRightWith(str, length, " ")
}

// Repeats the given string n times.
func Repeat(str string, count int) string {
	result := strings.Builder{}

	for i := 0; i < count; i++ {
		result.WriteString(str)
	}

	return result.String()
}

// Replaces string with replacement.
func Replace(source string, target string, newStr string) string {
	return strings.ReplaceAll(source, target, newStr)
}

// Replaces matches for pattern in string with replacement.
func ReplaceRegx(source string, pattern string, newStr string) (string, error) {
	regx, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}

	result := regx.ReplaceAllString(source, newStr)
	return result, nil
}

// Splits string by separator.
func Split(str string, separator string) []string {
	return strings.Split(str, separator)
}

// Splits string by separator and return limit count items.
func SplitWithCountLimit(str string, separator string, n int) []string {
	result := Split(str, separator)
	if len(result) <= n {
		return result
	} else {
		return result[0:n]
	}
}

// Checks if string starts with the given target string.
func StartsWith(str string, target string) bool {
	return strings.HasPrefix(str, target)
}

// Checks if string starts with the given target string from a specific position.
func StartsWithFrom(str string, target string, position int) bool {
	if position < 0 {
		position = 0
	} else if position > len(str) {
		position = len(str)
	}

	return StartsWith(str[position:], target)
}

// Converts string, as a whole, to lower case.
func ToLower(str string) string {
	return strings.ToLower(str)
}

// Converts string, as a whole, to upper case
func ToUpper(str string) string {
	return strings.ToUpper(str)
}

// Removes leading and trailing whitespace from string.
func Trim(str string) string {
	return strings.Trim(str, " ")
}

// Removes leading and trailing whitespace or specified characters from string.
func TrimWith(str string, trimChars string) string {
	return strings.Trim(str, trimChars)
}

// Removes leading whitespace from string.
func TrimStart(str string) string {
	return strings.TrimLeft(str, " ")
}

// Removes leading whitespace or specified characters from string.
func TrimStartWith(str string, trimChars string) string {
	return strings.TrimLeft(str, trimChars)
}

// Removes tailing whitespace from string.
func TrimEnd(str string) string {
	return strings.TrimRight(str, " ")
}

// Removes tailing whitespace or specified characters from string.
func TrimEndWith(str string, trimChars string) string {
	return strings.TrimRight(str, trimChars)
}

// The inverse of Escape func; this method converts the HTML entities &amp;, &lt;, &gt;, &quot;, and &#39;
// in string to their corresponding characters.
func Unescape(str string) string {
	return html.UnescapeString(str)
}

// Converts the first character of string to upper case.
func UpperFirst(str string) string {
	if len(str) <= 1 {
		return ToUpper(str)
	} else {
		return ToUpper(string(str[0])) + str[1:]
	}
}

//TODO: kebabCase(str string) string
//TODO: lowerCase(str string) string
//TODO: snakeCase(str string) string
//TODO: words
