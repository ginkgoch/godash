package godash

import (
	"gotest.tools/assert"
	"testing"
)

func TestCamelCase(t *testing.T) {
	str := "Hello world"
	str, _ = CamelCase(str)

	assert.Equal(t, str, "helloWorld")

	str, _ = CamelCase("Hello World from Ginkgo")
	assert.Equal(t, str, "helloWorldFromGinkgo")

	str, _ = CamelCase("001hello from gin")
	assert.Equal(t, str, "001HelloFromGin")

	str, _ = CamelCase("0aa1bb2cc")
	assert.Equal(t, str, "0Aa1Bb2Cc")

	str, _ = CamelCase("aa-bb-cc")
	assert.Equal(t, str, "aa-Bb-Cc")
}

func TestCamelCaseWithInit(t *testing.T) {
	str := "Hello world"
	str, _ = CamelCaseWithInit(str, true)

	assert.Equal(t, str, "HelloWorld")

	str, _ = CamelCaseWithInit("Hello World from Ginkgo", true)
	assert.Equal(t, str, "HelloWorldFromGinkgo")

	str, _ = CamelCaseWithInit("001hello from gin", true)
	assert.Equal(t, str, "001HelloFromGin")

	str, _ = CamelCaseWithInit("0aa1bb2cc", true)
	assert.Equal(t, str, "0Aa1Bb2Cc")

	str, _ = CamelCaseWithInit("aa-bb-cc", true)
	assert.Equal(t, str, "Aa-Bb-Cc")
}

func TestCapitalize(t *testing.T) {
	str := "jerry"
	str = Capitalize(str)

	assert.Equal(t, str, "Jerry")
}

func TestEndsWith(t *testing.T) {
	str := "Hello World"
	result := EndsWith(str, "d")
	assert.Equal(t, result, true)

	result = EndsWith(str, "l")
	assert.Equal(t, result, false)

	result = EndsWithFrom(str, "d", len(str))
	assert.Equal(t, result, true)

	result = EndsWithFrom(str, "l", len(str))
	assert.Equal(t, result, false)

	result = EndsWithFrom(str, "l", len(str)-1)
	assert.Equal(t, result, true)
}

func TestEscape(t *testing.T) {
	str := "<>"
	str = Escape(str)
	assert.Equal(t, str, "&lt;&gt;")
}

func TestEscapeRegExp(t *testing.T) {
	str := `Escaping symbols like: .+*?()|[]{}^$`
	str = EscapeRegExp(str)

	assert.Equal(t, str, `Escaping symbols like: \.\+\*\?\(\)\|\[\]\{\}\^\$`)
}

func TestLowerFirst(t *testing.T) {
	str := ""
	str = LowerFirst(str)
	assert.Equal(t, str, "")

	str = LowerFirst("A")
	assert.Equal(t, str, "a")

	str = LowerFirst("AA")
	assert.Equal(t, str, "aA")
}

func TestPad(t *testing.T) {
	str := "ABC"
	str = Pad(str, 2)
	assert.Equal(t, str, "ABC")

	str = Pad("ABC", 3)
	assert.Equal(t, str, "ABC")

	str = Pad("ABC", 4)
	assert.Equal(t, str, "ABC ")

	str = Pad("ABC", 5)
	assert.Equal(t, str, " ABC ")

	str = PadWith("ABC", 5, "-+")
	assert.Equal(t, str, "-ABC-")

	str = PadWith("ABC", 7, "-+")
	assert.Equal(t, str, "-+ABC-+")

	str = PadWith("ABC", 8, "-+")
	assert.Equal(t, str, "-+ABC-+-")
}

func TestPadLeft(t *testing.T) {
	str := "ABC"
	str = PadLeft(str, 2)
	assert.Equal(t, str, "ABC")

	str = PadLeft("ABC", 3)
	assert.Equal(t, str, "ABC")

	str = PadLeft("ABC", 4)
	assert.Equal(t, str, " ABC")

	str = PadLeft("ABC", 5)
	assert.Equal(t, str, "  ABC")

	str = PadLeftWith("ABC", 5, "-+")
	assert.Equal(t, str, "-+ABC")

	str = PadLeftWith("ABC", 7, "-+")
	assert.Equal(t, str, "-+-+ABC")
}

func TestPadRight(t *testing.T) {
	str := "ABC"
	str = PadRight(str, 2)
	assert.Equal(t, str, "ABC")

	str = PadRight("ABC", 3)
	assert.Equal(t, str, "ABC")

	str = PadRight("ABC", 4)
	assert.Equal(t, str, "ABC ")

	str = PadRight("ABC", 5)
	assert.Equal(t, str, "ABC  ")

	str = PadRightWith("ABC", 5, "-+")
	assert.Equal(t, str, "ABC-+")

	str = PadRightWith("ABC", 7, "-+")
	assert.Equal(t, str, "ABC-+-+")
}

func TestRepeat(t *testing.T) {
	str := "ABC"
	str = Repeat(str, 0)
	assert.Equal(t, str, "")

	str = Repeat("ABC", 1)
	assert.Equal(t, str, "ABC")

	str = Repeat("ABC", 2)
	assert.Equal(t, str, "ABCABC")
}

func TestReplace(t *testing.T) {
	str := "ABC-DEF-G"
	str = Replace(str, "-", "_")
	assert.Equal(t, str, "ABC_DEF_G")

	str, _ = ReplaceRegx(str, `[A-Z]`, "*")
	assert.Equal(t, str, "***_***_*")
}

func TestSplit(t *testing.T) {
	str := "A-B-C-D-E-F"
	r := Split(str, "-")
	assert.DeepEqual(t, r, []string{"A", "B", "C", "D", "E", "F"})

	r = SplitWithCountLimit(str, "-", 100)
	assert.DeepEqual(t, r, []string{"A", "B", "C", "D", "E", "F"})

	r = SplitWithCountLimit(str, "-", 3)
	assert.DeepEqual(t, r, []string{"A", "B", "C"})
}

func TestStartsWith(t *testing.T) {
	str := "Hello"
	r := StartsWith(str, "He")
	assert.Equal(t, r, true)

	r = StartsWith(str, "el")
	assert.Equal(t, r, false)

	r = StartsWithFrom(str, "el", 1)
	assert.Equal(t, r, true)
}

func TestUnescape(t *testing.T) {
	str := "&lt;&gt;"
	str = Unescape(str)
	assert.Equal(t, str, "<>")
}

func TestUpperFirst(t *testing.T) {
	str := "fred"
	str = UpperFirst(str)
	assert.Equal(t, str, "Fred")

	str = UpperFirst("FRED")
	assert.Equal(t, str, "FRED")
}
