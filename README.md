[![PkgGoDev](https://pkg.go.dev/badge/github.com/ginkgoch/godash)](https://pkg.go.dev/github.com/ginkgoch/godash/pkg)
[![Go Report Card](https://goreportcard.com/badge/github.com/ginkgoch/godash)](https://goreportcard.com/report/github.com/ginkgoch/godash)

[![forthebadge](https://forthebadge.com/images/badges/gluten-free.svg)](https://forthebadge.com)

# Lodash in Golang Version V2
This is lodash in golang version v2. A modern Golang utility library delivering modularity, performance & extras. Compare to [v1](./README_V1.md), all APIs are refactored with generic types support. [API Reference](https://pkg.go.dev/github.com/ginkgoch/godash/v2).


## Install & Import
```bash
go get -u "github.com/ginkgoch/godash/v2"
```

```go
import "github.com/ginkgoch/godash/v2"
```


## Usage

#### func  CamelCase

```go
func CamelCase(str string) (string, error)
```
Converts string to camel case. First char is lower case.

#### func  CamelCaseWithInit

```go
func CamelCaseWithInit(str string, upperCase InitCamelCase) (string, error)
```
Converts string to camel case. First char is lower case by default.

#### func  Capitalize

```go
func Capitalize(str string) string
```
Converts the first character of string to upper case and the remaining to lower
case.

#### func  Chunk

```go
func Chunk[E any](items []E, size int) [][]E
```
Creates an array of elements split into groups the length of size. If array
can't be split evenly, the final chunk will be the remaining elements.

#### func  Compact

```go
func Compact[E any](items []E) []E
```
Creates an array with all falsy values removed. The values false, 0, "", nil are
falsy.

#### func  Concat

```go
func Concat[E any](items []E, newItems []E) []E
```
Creates a new array concatenating array with any additional arrays and/or
values.

#### func  ConcatSlices

```go
func ConcatSlices[E any](slices ...[]E) []E
```
Creates a new array concatenating array with any additional DashSlices.

#### func  Difference

```go
func Difference[E any](items []E, itemsToCompare []E) []E
```
Creates an array of array values not included in the other given arrays using
SameValueZero for equality comparisons. The order and references of result
values are determined by the first array.

#### func  DifferenceBy

```go
func DifferenceBy[E any](items []E, itemsToCompare []E, iteratee Iteratee[E, E]) []E
```
This method is like _.difference except that it accepts iteratee which is
invoked for each element of array and values to generate the criterion by which
they're compared. The order and references of result values are determined by
the first array. The iteratee is invoked with one argument: (value).

#### func  DifferenceWith

```go
func DifferenceWith[E any](items []E, itemsToCompare []E,
	comparison Comparison[E]) []E
```
This method is like _.difference except that it accepts comparator which is
invoked to compare elements of array to values. The order and references of
result values are determined by the first array. The comparator is invoked with
two arguments: (arrVal, othVal).

#### func  Drop

```go
func Drop[E any](items []E, count int) []E
```
Creates a slice of array with n elements dropped from the beginning.

#### func  DropRight

```go
func DropRight[E any](items []E, count int) []E
```
Creates a slice of array with n elements dropped from the end.

#### func  DropWhile

```go
func DropWhile[E any](items []E, predicate Predicate[E]) []E
```
Creates a slice of array excluding elements dropped from the beginning. Elements
are dropped until predicate returns falsy. The predicate is invoked with two
arguments: (value, index).

#### func  EndsWith

```go
func EndsWith(str string, target string) bool
```
Checks if string ends with the given target string.

#### func  EndsWithFrom

```go
func EndsWithFrom(str string, target string, position int) bool
```
Checks if string ends with the given target string with the position to search
up to.

#### func  Escape

```go
func Escape(str string) string
```
Converts the characters "&", "<", ">", '"', and "'" in string to their
corresponding HTML entities.

#### func  EscapeRegExp

```go
func EscapeRegExp(str string) string
```
Escapes the RegExp special characters "^", "$", "", ".", "*", "+", "?", "(",
")", "[", "]", "{", "}", and "|" in string.

#### func  Fill

```go
func Fill[E any](items []E, fillValue E)
```
Fills elements of array with value.

#### func  FillInRange

```go
func FillInRange[E any](items []E, value E, start int, end int)
```
Fills elements of array with value from start up to, but not including end.

#### func  Filter

```go
func Filter[E any](slice []E, predicate Predicate[E]) []E
```

#### func  FindIndex

```go
func FindIndex[E any](items []E, predicate Predicate[E]) (int, bool)
```
This method is like Find except that it returns the index of the first element
predicate returns truthy for instead of the element itself.

#### func  FindIndexWith

```go
func FindIndexWith[E any](items []E, element E, comparison Comparison[E]) (index int, ok bool)
```
Same to IndexOf. The difference is that, this method provides a comparison
function to compare programmatically.

#### func  FindLastIndex

```go
func FindLastIndex[E any](items []E, predicate Predicate[E]) (int, bool)
```
This method is like Find except that it returns the index of the first element
predicate returns truthy for instead of the element itself.

#### func  FindLastIndexWith

```go
func FindLastIndexWith[E any](items []E, element E, comparison Comparison[E]) (int, bool)
```
This method is like FindIndex except that it iterates over elements of
collection from right to left.

#### func  First

```go
func First[E any](items []E) *E
```
Gets the first element of array.

#### func  FromPairs

```go
func FromPairs[K comparable, V any](pairs []KeyValuePair[K, V]) map[K]V
```
This method returns an object composed from key-value pairs.

#### func  FromPairsAny

```go
func FromPairsAny(pairs [][]any) map[any]any
```
This method returns an object composed from key-value pairs.

#### func  Head

```go
func Head[E any](items []E) *E
```
Gets the first element of slice.

#### func  IndexOf

```go
func IndexOf[E any](items []E, element E) (int, bool)
```
This method is like _.find except that it returns the index of the first element
predicate returns truthy for instead of the element itself.

#### func  Initial

```go
func Initial[E any](slice []E) []E
```
Gets all but the last element of array.

#### func  Intersection

```go
func Intersection[E any](items1 []E, items2 []E) (intersectedItems []E)
```
Creates an array of unique values that are included in all given arrays using
SameValueZero for equality comparisons. The order and references of result
values are determined by the first array.

#### func  IntersectionBy

```go
func IntersectionBy[E, T any](items1 []E, items2 []E, iteratee Iteratee[E, T]) (intersectedItems []E)
```
This method is like Intersection except that it accepts iteratee which is
invoked for each element of each arrays to generate the criterion by which
they're compared. The order and references of result values are determined by
the first array. The iteratee is invoked with one argument: (value).

#### func  IntersectionWith

```go
func IntersectionWith[E any](items1 []E, items2 []E, comparison Comparison[E]) (intersectedItems []E)
```
This method is like _.intersection except that it accepts comparator which is
invoked to compare elements of arrays. The order and references of result values
are determined by the first array. The comparator is invoked with two arguments:
(arrVal, othVal).

#### func  Join

```go
func Join[E any](items []E, separator string) string
```
Converts all elements in array into a string separated by separator.

#### func  Last

```go
func Last[E any](items []E) (exists bool, lastItem E)
```
Gets the last element of array.

#### func  LastIndexOf

```go
func LastIndexOf[E any](items []E, element E) (int, bool)
```
This method is like IndexOf except that it iterates over elements of array from
right to left.

#### func  LowerFirst

```go
func LowerFirst(str string) string
```

#### func  Map

```go
func Map[E, V any](slice []E, iteratee func(E) V) []V
```

#### func  Nth

```go
func Nth[E any](items []E, n int) (exists bool, item E)
```
Gets the element at index n of array. If n is negative, the nth element from the
end is returned.

#### func  Pad

```go
func Pad(str string, length int) string
```
Pads string on the left and right sides if it's shorter than length. Padding
characters are truncated if they can't be evenly divided by length.

#### func  PadLeft

```go
func PadLeft(str string, length int) string
```
Pads string on the left sides if it's shorter than length.

#### func  PadLeftWith

```go
func PadLeftWith(str string, length int, padChars string) string
```
Pads string on the left sides if it's shorter than length.

#### func  PadRight

```go
func PadRight(str string, length int) string
```
Pads string on the right sides if it's shorter than length.

#### func  PadRightWith

```go
func PadRightWith(str string, length int, padChars string) string
```
Pads string on the right sides if it's shorter than length.

#### func  PadWith

```go
func PadWith(str string, length int, padChars string) string
```
Pads string on the left and right sides if it's shorter than length. Padding
characters are truncated if they can't be evenly divided by length.

#### func  Pull

```go
func Pull[E comparable](items *[]E, values ...E) []E
```
Removes all given values from array using SameValueZero for equality
comparisons.

#### func  PullAll

```go
func PullAll[E comparable](items *[]E, values []E) []E
```
This method is like Pull except that it accepts an array of values to remove.

#### func  PullAllWith

```go
func PullAllWith[E any](items *[]E, values []E, comparison Comparison[E]) []E
```
This method is like PullAll except that it accepts comparator which is invoked
to compare elements of array to values. The comparator is invoked with two
arguments: (arrVal, othVal).

#### func  PullAt

```go
func PullAt[E any](items *[]E, indices ...int) (pulled []E)
```
Removes elements from array corresponding to indexes and returns an array of
removed elements.

#### func  Remove

```go
func Remove[E any](items *[]E, predicate Predicate[E]) (removed []E)
```
Removes all elements from array that predicate returns truthy for and returns an
array of the removed elements. The predicate is invoked with two arguments:
(value, index).

#### func  Repeat

```go
func Repeat(str string, count int) string
```
Repeats the given string n times.

#### func  Replace

```go
func Replace(source string, target string, newStr string) string
```
Replaces string with replacement.

#### func  ReplaceRegx

```go
func ReplaceRegx(source string, pattern string, newStr string) (string, error)
```
Replaces matches for pattern in string with replacement.

#### func  Reverse

```go
func Reverse[E any](items []E) []E
```
Reverses array so that the first element becomes the last, the second element
becomes the second to last, and so on.

#### func  Slice

```go
func Slice[E any](items []E, start int, end int) []E
```
Creates a slice of array from start up to, but not including, end.

#### func  Split

```go
func Split(str string, separator string) []string
```
Splits string by separator.

#### func  SplitWithCountLimit

```go
func SplitWithCountLimit(str string, separator string, n int) []string
```
Splits string by separator and return limit count items.

#### func  StartsWith

```go
func StartsWith(str string, target string) bool
```
Checks if string starts with the given target string.

#### func  StartsWithFrom

```go
func StartsWithFrom(str string, target string, position int) bool
```
Checks if string starts with the given target string from a specific position.

#### func  Tail

```go
func Tail[E any](items []E) (result []E)
```
Gets all but the first element of array.

#### func  Take

```go
func Take[E any](items []E, n int) (results []E)
```
Creates a slice of array with n elements taken from the beginning.

#### func  TakeRight

```go
func TakeRight[E any](items []E, n int) []E
```
Creates a slice of array with n elements taken from the end.

#### func  TakeRightWhile

```go
func TakeRightWhile[E any](items []E, predicate Predicate[E]) []E
```
Creates a slice of array with elements taken from the end. Elements are taken
until predicate returns falsy. The predicate is invoked with one argument:
(value).

#### func  TakeWhile

```go
func TakeWhile[E any](items []E, predicate Predicate[E]) []E
```
Creates a slice of array with elements taken from the beginning. Elements are
taken until predicate returns falsy. The predicate is invoked with one argument:
(value).

#### func  Ternary

```go
func Ternary(satisfy bool, truthyValue interface{}, falsyValue interface{}) interface{}
```

#### func  ToLower

```go
func ToLower(str string) string
```
Converts string, as a whole, to lower case.

#### func  ToUpper

```go
func ToUpper(str string) string
```
Converts string, as a whole, to upper case

#### func  Trim

```go
func Trim(str string) string
```
Removes leading and trailing whitespace from string.

#### func  TrimEnd

```go
func TrimEnd(str string) string
```
Removes tailing whitespace from string.

#### func  TrimEndWith

```go
func TrimEndWith(str string, trimChars string) string
```
Removes tailing whitespace or specified characters from string.

#### func  TrimStart

```go
func TrimStart(str string) string
```
Removes leading whitespace from string.

#### func  TrimStartWith

```go
func TrimStartWith(str string, trimChars string) string
```
Removes leading whitespace or specified characters from string.

#### func  TrimWith

```go
func TrimWith(str string, trimChars string) string
```
Removes leading and trailing whitespace or specified characters from string.

#### func  Unescape

```go
func Unescape(str string) string
```
The inverse of Escape func; this method converts the HTML entities &amp;, &lt;,
&gt;, &quot;, and &#39; in string to their corresponding characters.

#### func  Union

```go
func Union[E comparable](slices ...[]E) []E
```
Creates an array of unique values, in order, from all given arrays using
SameValueZero for equality comparisons.

#### func  UnionBy

```go
func UnionBy[I any, O comparable](iteratee Iteratee[I, O], slices ...[]I) []I
```
This method is like Uniq except that it accepts iteratee which is invoked for
each element in array to generate the criterion by which uniqueness is computed.
The order of result values is determined by the order they occur in the array.
The iteratee is invoked with one argument: (value).

#### func  UnionWith

```go
func UnionWith[E any](comparison Comparison[E], slices ...[]E) []E
```
This method is like Uniq except that it accepts comparator which is invoked to
compare elements of array. The order of result values is determined by the order
they occur in the array. The comparator is invoked with two arguments: (arrVal,
othVal).

#### func  Uniq

```go
func Uniq[E comparable](items []E) []E
```
Creates a duplicate-free version of an array, using SameValueZero for equality
comparisons, in which only the first occurrence of each element is kept. The
order of result values is determined by the order they occur in the array.

#### func  UniqBy

```go
func UniqBy[I any, O comparable](items []I, iteratee Iteratee[I, O]) []I
```
This method is like Union except that it accepts iteratee which is invoked for
each element of each arrays to generate the criterion by which uniqueness is
computed. Result values are chosen from the first array in which the value
occurs. The iteratee is invoked with one argument: (value).

#### func  UniqWith

```go
func UniqWith[E any](items []E, comparison Comparison[E]) []E
```
This method is like Uniq except that it accepts comparator which is invoked to
compare elements of array. The order of result values is determined by the order
they occur in the array. The comparator is invoked with two arguments: (arrVal,
othVal).

#### func  UpperFirst

```go
func UpperFirst(str string) string
```
Converts the first character of string to upper case.

#### func  Without

```go
func Without[E comparable](items []E, values ...E) []E
```
Creates an array excluding all given values using SameValueZero for equality
comparisons.

#### func  Xor

```go
func Xor[E comparable](items ...[]E) []E
```
Creates an array of unique values that is the symmetric difference of the given
arrays. The order of result values is determined by the order they occur in the
arrays.

#### func  Zip

```go
func Zip[E any](slices ...[]E) [][]E
```
Creates an array of grouped elements, the first of which contains the first
elements of the given arrays, the second of which contains the second elements
of the given arrays, and so on.

#### func  ZipWith

```go
func ZipWith[E any, V any](iteratee Iteratee[[]E, V], slices ...[]E) []V
```
This method is like Zip except that it accepts iteratee to specify how grouped
values should be combined. The iteratee is invoked with the elements of each
group: (...group).

#### type Comparison

```go
type Comparison[E any] func(E, E) bool
```


#### type InitCamelCase

```go
type InitCamelCase bool
```


#### type Iteratee

```go
type Iteratee[E any, V any] func(E) V
```


#### type KeyValuePair

```go
type KeyValuePair[K comparable, V any] struct {
}
```


#### type Number

```go
type Number interface {
	int | int16 | int32 | int64 | int8 | float32 | float64 | uint | uint16 | uint32 | uint64 | uint8
}
```


#### type Predicate

```go
type Predicate[E any] func(E) bool
```
