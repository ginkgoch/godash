# Lodash in Golang Version 
--
    import "github.com/ginkgoch/godash"

This is lodash in golang version. A modern Golang utility library delivering
modularity, performance & extras.

<!-- vscode-markdown-toc -->
## Index

* 1.1. [func  CountBy](#funcCountBy)
* 1.2. [func  Every](#funcEvery)
* 1.3. [func  Fill](#funcFill)
* 1.4. [func  FillInRange](#funcFillInRange)
* 1.5. [func  Find](#funcFind)
* 1.6. [func  FindFrom](#funcFindFrom)
* 1.7. [func  FindIndex](#funcFindIndex)
* 1.8. [func  FindIndexWith](#funcFindIndexWith)
* 1.9. [func  FindLast](#funcFindLast)
* 1.10. [func  FindLastFrom](#funcFindLastFrom)
* 1.11. [func  FindLastIndex](#funcFindLastIndex)
* 1.12. [func  FindLastIndexWith](#funcFindLastIndexWith)
* 1.13. [func  First](#funcFirst)
* 1.14. [func  FromPairs](#funcFromPairs)
* 1.15. [func  GroupBy](#funcGroupBy)
* 1.16. [func  Head](#funcHead)
* 1.17. [func  Identity](#funcIdentity)
* 1.18. [func  IdentityFloat64](#funcIdentityFloat64)
* 1.19. [func  IdentityInt](#funcIdentityInt)
* 1.20. [func  IdentityString](#funcIdentityString)
* 1.21. [func  Includes](#funcIncludes)
* 1.22. [func  IndexOf](#funcIndexOf)
* 1.23. [func  Join](#funcJoin)
* 1.24. [func  Last](#funcLast)
* 1.25. [func  LastIndexOf](#funcLastIndexOf)
* 1.26. [func  Nth](#funcNth)
* 1.27. [func  Reduce](#funcReduce)
* 1.28. [func  ReduceRight](#funcReduceRight)
* 1.29. [func  ReduceRightWithInitial](#funcReduceRightWithInitial)
* 1.30. [func  ReduceWithInitial](#funcReduceWithInitial)
* 1.31. [func  Sample](#funcSample)
* 1.32. [func  Size](#funcSize)
* 1.33. [func  Some](#funcSome)
* 1.34. [type Action](#typeAction)
* 1.35. [type Comparison](#typeComparison)
* 1.36. [type DashSlice](#typeDashSlice)
* 1.37. [func  Chunk](#funcChunk)
* 1.38. [func  Compact](#funcCompact)
* 1.39. [func  Concat](#funcConcat)
* 1.40. [func  ConcatSlices](#funcConcatSlices)
* 1.41. [func  Difference](#funcDifference)
* 1.42. [func  DifferenceBy](#funcDifferenceBy)
* 1.43. [func  DifferenceWith](#funcDifferenceWith)
* 1.44. [func  Drop](#funcDrop)
* 1.45. [func  DropRight](#funcDropRight)
* 1.46. [func  DropWhile](#funcDropWhile)
* 1.47. [func  Each](#funcEach)
* 1.48. [func  EachRight](#funcEachRight)
* 1.49. [func  Filter](#funcFilter)
* 1.50. [func  FlatMap](#funcFlatMap)
* 1.51. [func  FlatMapDeep](#funcFlatMapDeep)
* 1.52. [func  FlatMapDepth](#funcFlatMapDepth)
* 1.53. [func  Flatten](#funcFlatten)
* 1.54. [func  FlattenDeep](#funcFlattenDeep)
* 1.55. [func  FlattenDepth](#funcFlattenDepth)
* 1.56. [func  ForEach](#funcForEach)
* 1.57. [func  ForEachRight](#funcForEachRight)
* 1.58. [func  Initial](#funcInitial)
* 1.59. [func  Intersection](#funcIntersection)
* 1.60. [func  IntersectionBy](#funcIntersectionBy)
* 1.61. [func  IntersectionWith](#funcIntersectionWith)
* 1.62. [func  Map](#funcMap)
* 1.63. [func  NewDashSlice](#funcNewDashSlice)
* 1.64. [func  NewDashSliceFromIntArray](#funcNewDashSliceFromIntArray)
* 1.65. [func  Pull](#funcPull)
* 1.66. [func  PullAll](#funcPullAll)
* 1.67. [func  PullAllWith](#funcPullAllWith)
* 1.68. [func  PullAt](#funcPullAt)
* 1.69. [func  Reject](#funcReject)
* 1.70. [func  Remove](#funcRemove)
* 1.71. [func  Reverse](#funcReverse)
* 1.72. [func  SampleSize](#funcSampleSize)
* 1.73. [func  Shuffle](#funcShuffle)
* 1.74. [func  Slice](#funcSlice)
* 1.75. [func  SortByFloat64](#funcSortByFloat64)
* 1.76. [func  SortByInt](#funcSortByInt)
* 1.77. [func  SortByString](#funcSortByString)
* 1.78. [func  Tail](#funcTail)
* 1.79. [func  Take](#funcTake)
* 1.80. [func  TakeRight](#funcTakeRight)
* 1.81. [func  TakeRightWhile](#funcTakeRightWhile)
* 1.82. [func  TakeWhile](#funcTakeWhile)
* 1.83. [func  Union](#funcUnion)
* 1.84. [func  UnionBy](#funcUnionBy)
* 1.85. [func  UnionWith](#funcUnionWith)
* 1.86. [func  Uniq](#funcUniq)
* 1.87. [func  UniqBy](#funcUniqBy)
* 1.88. [func  UniqWith](#funcUniqWith)
* 1.89. [func  Without](#funcWithout)
* 1.90. [func  Xor](#funcXor)
* 1.91. [func  Zip](#funcZip)
* 1.92. [func  ZipWith](#funcZipWith)
* 1.93. [func (DashSlice) Map](#funcDashSliceMap)
* 1.94. [type Iteratee](#typeIteratee)
* 1.95. [type IterateeToFloat](#typeIterateeToFloat)
* 1.96. [type IterateeToInt](#typeIterateeToInt)
* 1.97. [type IterateeToString](#typeIterateeToString)
* 1.98. [type Predicate](#typePredicate)
* 1.99. [type Reducer](#typeReducer)

<!-- vscode-markdown-toc-config
	numbering=true
	autoSave=true
	/vscode-markdown-toc-config -->
<!-- /vscode-markdown-toc -->

##  1. <a name='Usage'></a>Usage

####  1.1. <a name='funcCountBy'></a>func  CountBy

```go
func CountBy(items DashSlice, iteratee Iteratee) map[interface{}]int
```
Creates an object composed of keys generated from the results of running each
element of collection thru iteratee. The corresponding value of each key is the
number of times the key was returned by iteratee. The iteratee is invoked with
one argument: (value).

####  1.2. <a name='funcEvery'></a>func  Every

```go
func Every(items DashSlice, predicate Predicate) bool
```
Checks if predicate returns truthy for all elements of collection. Iteration is
stopped once predicate returns falsy. The predicate is invoked with one
argument: (value).

####  1.3. <a name='funcFill'></a>func  Fill

```go
func Fill(items DashSlice, fillValue interface{})
```
Fills elements of array with value.

####  1.4. <a name='funcFillInRange'></a>func  FillInRange

```go
func FillInRange(items DashSlice, value interface{}, start int, end int)
```
Fills elements of array with value from start up to, but not including end.

####  1.5. <a name='funcFind'></a>func  Find

```go
func Find(items DashSlice, predicate Predicate) (interface{}, bool)
```
Iterates over elements of collection, returning the first element predicate
returns truthy for. The predicate is invoked with one argument: (value).

####  1.6. <a name='funcFindFrom'></a>func  FindFrom

```go
func FindFrom(items DashSlice, predicate Predicate, start int) (interface{}, bool)
```
Iterates over elements of collection from start index, returning the first
element predicate returns truthy for. The predicate is invoked with one
argument: (value).

####  1.7. <a name='funcFindIndex'></a>func  FindIndex

```go
func FindIndex(items DashSlice, predicate Predicate) (int, bool)
```
This method is like Find except that it returns the index of the first element
predicate returns truthy for instead of the element itself.

####  1.8. <a name='funcFindIndexWith'></a>func  FindIndexWith

```go
func FindIndexWith(items DashSlice, element interface{}, comparison Comparison) (int, bool)
```
Same to IndexOf. The difference is that, this method provides a comparison
function to compare programmatically.

####  1.9. <a name='funcFindLast'></a>func  FindLast

```go
func FindLast(items DashSlice, predicate Predicate) (interface{}, bool)
```
This method is like Find except that it iterates over elements of collection
from right to left.

####  1.10. <a name='funcFindLastFrom'></a>func  FindLastFrom

```go
func FindLastFrom(items DashSlice, predicate Predicate, start int) (interface{}, bool)
```
This method is like FindFrom except that it iterates over elements of collection
from right to left.

####  1.11. <a name='funcFindLastIndex'></a>func  FindLastIndex

```go
func FindLastIndex(items DashSlice, predicate Predicate) (int, bool)
```
This method is like Find except that it returns the index of the first element
predicate returns truthy for instead of the element itself.

####  1.12. <a name='funcFindLastIndexWith'></a>func  FindLastIndexWith

```go
func FindLastIndexWith(items DashSlice, element interface{}, comparison Comparison) (int, bool)
```
This method is like FindIndex except that it iterates over elements of
collection from right to left.

####  1.13. <a name='funcFirst'></a>func  First

```go
func First(items DashSlice) interface{}
```
Gets the first element of array.

####  1.14. <a name='funcFromPairs'></a>func  FromPairs

```go
func FromPairs(pairs []DashSlice) map[interface{}]interface{}
```
This method returns an object composed from key-value pairs.

####  1.15. <a name='funcGroupBy'></a>func  GroupBy

```go
func GroupBy(items DashSlice, iteratee Iteratee) map[interface{}]DashSlice
```
Creates an object composed of keys generated from the results of running each
element of collection thru iteratee. The order of grouped values is determined
by the order they occur in collection. The corresponding value of each key is an
array of elements responsible for generating the key. The iteratee is invoked
with one argument: (value).

####  1.16. <a name='funcHead'></a>func  Head

```go
func Head(items DashSlice) interface{}
```
Gets the first element of slice.

####  1.17. <a name='funcIdentity'></a>func  Identity

```go
func Identity(i interface{}) interface{}
```

####  1.18. <a name='funcIdentityFloat64'></a>func  IdentityFloat64

```go
func IdentityFloat64(i interface{}) float64
```

####  1.19. <a name='funcIdentityInt'></a>func  IdentityInt

```go
func IdentityInt(i interface{}) int
```

####  1.20. <a name='funcIdentityString'></a>func  IdentityString

```go
func IdentityString(i interface{}) string
```

####  1.21. <a name='funcIncludes'></a>func  Includes

```go
func Includes(items DashSlice, value interface{}) bool
```
Checks if value is in collection. If collection is a string, it's checked for a
substring of value, otherwise SameValueZero is used for equality comparisons. If
fromIndex is negative, it's used as the offset from the end of collection.

####  1.22. <a name='funcIndexOf'></a>func  IndexOf

```go
func IndexOf(items DashSlice, element interface{}) (int, bool)
```
This method is like _.find except that it returns the index of the first element
predicate returns truthy for instead of the element itself.

####  1.23. <a name='funcJoin'></a>func  Join

```go
func Join(items DashSlice, separator string) string
```
Converts all elements in array into a string separated by separator.

####  1.24. <a name='funcLast'></a>func  Last

```go
func Last(items DashSlice) interface{}
```
Gets the last element of array.

####  1.25. <a name='funcLastIndexOf'></a>func  LastIndexOf

```go
func LastIndexOf(items DashSlice, element interface{}) (int, bool)
```
This method is like IndexOf except that it iterates over elements of array from
right to left.

####  1.26. <a name='funcNth'></a>func  Nth

```go
func Nth(items DashSlice, n int) interface{}
```
Gets the element at index n of array. If n is negative, the nth element from the
end is returned.

####  1.27. <a name='funcReduce'></a>func  Reduce

```go
func Reduce(items DashSlice, reducer Reducer) interface{}
```
Reduces collection to a value which is the accumulated result of running each
element in collection thru iteratee, where each successive invocation is
supplied the return value of the previous. If accumulator is not given, the
first element of collection is used as the initial value. The reducer is invoked
with two arguments: (accumulator, value).

####  1.28. <a name='funcReduceRight'></a>func  ReduceRight

```go
func ReduceRight(items DashSlice, reducer Reducer) interface{}
```
This method is like Reduce except that it iterates over elements of collection
from right to left.

####  1.29. <a name='funcReduceRightWithInitial'></a>func  ReduceRightWithInitial

```go
func ReduceRightWithInitial(items DashSlice, reducer Reducer, initial interface{}) interface{}
```
This method is like ReduceWithInitial except that it iterates over elements of
collection from right to left.

####  1.30. <a name='funcReduceWithInitial'></a>func  ReduceWithInitial

```go
func ReduceWithInitial(items DashSlice, reducer Reducer, initial interface{}) interface{}
```
Reduces collection to a value which is the accumulated result of running each
element in collection thru iteratee, where each successive invocation is
supplied the return value of the previous. If accumulator is not given, the
first element of collection is used as the initial value. The reducer is invoked
with two arguments: (accumulator, value).

####  1.31. <a name='funcSample'></a>func  Sample

```go
func Sample(items DashSlice) interface{}
```
Gets a random element from collection.

####  1.32. <a name='funcSize'></a>func  Size

```go
func Size(items DashSlice) int
```
Gets the size of collection.

####  1.33. <a name='funcSome'></a>func  Some

```go
func Some(items DashSlice, predicate Predicate) bool
```
Checks if predicate returns truthy for any element of collection. Iteration is
stopped once predicate returns truthy. The predicate is invoked with one
argument: (value).

####  1.34. <a name='typeAction'></a>type Action

```go
type Action func(interface{}, int)
```


####  1.35. <a name='typeComparison'></a>type Comparison

```go
type Comparison func(interface{}, interface{}) bool
```


####  1.36. <a name='typeDashSlice'></a>type DashSlice

```go
type DashSlice []interface{}
```


####  1.37. <a name='funcChunk'></a>func  Chunk

```go
func Chunk(items DashSlice, size int) []DashSlice
```
Creates an array of elements split into groups the length of size. If array
can't be split evenly, the final chunk will be the remaining elements.

####  1.38. <a name='funcCompact'></a>func  Compact

```go
func Compact(items DashSlice) DashSlice
```
Creates an array with all falsy values removed. The values false, 0, "", nil are
falsy.

####  1.39. <a name='funcConcat'></a>func  Concat

```go
func Concat(items DashSlice, newItems ...interface{}) DashSlice
```
Creates a new array concatenating array with any additional arrays and/or
values.

####  1.40. <a name='funcConcatSlices'></a>func  ConcatSlices

```go
func ConcatSlices(slices ...DashSlice) DashSlice
```
Creates a new array concatenating array with any additional DashSlices.

####  1.41. <a name='funcDifference'></a>func  Difference

```go
func Difference(items DashSlice, itemsToCompare ...interface{}) DashSlice
```
Creates an array of array values not included in the other given arrays using
SameValueZero for equality comparisons. The order and references of result
values are determined by the first array.

####  1.42. <a name='funcDifferenceBy'></a>func  DifferenceBy

```go
func DifferenceBy(items DashSlice, itemsToCompare DashSlice, iteratee Iteratee) DashSlice
```
This method is like _.difference except that it accepts iteratee which is
invoked for each element of array and values to generate the criterion by which
they're compared. The order and references of result values are determined by
the first array. The iteratee is invoked with one argument: (value).

####  1.43. <a name='funcDifferenceWith'></a>func  DifferenceWith

```go
func DifferenceWith(items DashSlice, itemsToCompare DashSlice,
	comparison Comparison) DashSlice
```
This method is like _.difference except that it accepts comparator which is
invoked to compare elements of array to values. The order and references of
result values are determined by the first array. The comparator is invoked with
two arguments: (arrVal, othVal).

####  1.44. <a name='funcDrop'></a>func  Drop

```go
func Drop(items DashSlice, count int) DashSlice
```
Creates a slice of array with n elements dropped from the beginning.

####  1.45. <a name='funcDropRight'></a>func  DropRight

```go
func DropRight(items DashSlice, count int) DashSlice
```
Creates a slice of array with n elements dropped from the end.

####  1.46. <a name='funcDropWhile'></a>func  DropWhile

```go
func DropWhile(items DashSlice, predicate Predicate) DashSlice
```
Creates a slice of array excluding elements dropped from the beginning. Elements
are dropped until predicate returns falsy. The predicate is invoked with two
arguments: (value, index).

####  1.47. <a name='funcEach'></a>func  Each

```go
func Each(items DashSlice, action Action) DashSlice
```
Iterates over elements of collection and invokes iteratee for each element. The
iteratee is invoked with three arguments: (value, index|key, collection).
Iteratee functions may exit iteration early by explicitly returning false.

####  1.48. <a name='funcEachRight'></a>func  EachRight

```go
func EachRight(items DashSlice, action Action) DashSlice
```
This method is like ForEach except that it iterates over elements of collection
from right to left.

####  1.49. <a name='funcFilter'></a>func  Filter

```go
func Filter(items DashSlice, predicate Predicate) DashSlice
```
Iterates over elements of collection, returning an array of all elements
predicate returns truthy for. The predicate is invoked with one argument:
(value).

####  1.50. <a name='funcFlatMap'></a>func  FlatMap

```go
func FlatMap(items DashSlice, iteratee Iteratee) DashSlice
```
Creates a flattened array of values by running each element in collection thru
iteratee and flattening the mapped results. The iteratee is invoked with one
argument: (value).

####  1.51. <a name='funcFlatMapDeep'></a>func  FlatMapDeep

```go
func FlatMapDeep(items DashSlice, iteratee Iteratee) DashSlice
```
This method is like FlatMap except that it recursively flattens the mapped
results.

####  1.52. <a name='funcFlatMapDepth'></a>func  FlatMapDepth

```go
func FlatMapDepth(items DashSlice, iteratee Iteratee, depth int) DashSlice
```
This method is like FlatMap except that it recursively flattens the mapped
results up to depth times.

####  1.53. <a name='funcFlatten'></a>func  Flatten

```go
func Flatten(items DashSlice) DashSlice
```
Flattens array a single level deep.

####  1.54. <a name='funcFlattenDeep'></a>func  FlattenDeep

```go
func FlattenDeep(items DashSlice) DashSlice
```
Recursively flattens array.

####  1.55. <a name='funcFlattenDepth'></a>func  FlattenDepth

```go
func FlattenDepth(items DashSlice, depth int) DashSlice
```
Recursively flatten array up to depth times.

####  1.56. <a name='funcForEach'></a>func  ForEach

```go
func ForEach(items DashSlice, action Action) DashSlice
```
Iterates over elements of collection and invokes iteratee for each element. The
iteratee is invoked with three arguments: (value, index|key, collection).
Iteratee functions may exit iteration early by explicitly returning false.

####  1.57. <a name='funcForEachRight'></a>func  ForEachRight

```go
func ForEachRight(items DashSlice, action Action) DashSlice
```
This method is like ForEach except that it iterates over elements of collection
from right to left.

####  1.58. <a name='funcInitial'></a>func  Initial

```go
func Initial(items DashSlice) DashSlice
```
Gets all but the last element of array.

####  1.59. <a name='funcIntersection'></a>func  Intersection

```go
func Intersection(items1 DashSlice, items2 DashSlice) DashSlice
```
Creates an array of unique values that are included in all given arrays using
SameValueZero for equality comparisons. The order and references of result
values are determined by the first array.

####  1.60. <a name='funcIntersectionBy'></a>func  IntersectionBy

```go
func IntersectionBy(items1 DashSlice, items2 DashSlice, iteratee Iteratee) DashSlice
```
This method is like Intersection except that it accepts iteratee which is
invoked for each element of each arrays to generate the criterion by which
they're compared. The order and references of result values are determined by
the first array. The iteratee is invoked with one argument: (value).

####  1.61. <a name='funcIntersectionWith'></a>func  IntersectionWith

```go
func IntersectionWith(items1 DashSlice, items2 DashSlice, comparison Comparison) DashSlice
```
This method is like _.intersection except that it accepts comparator which is
invoked to compare elements of arrays. The order and references of result values
are determined by the first array. The comparator is invoked with two arguments:
(arrVal, othVal).

####  1.62. <a name='funcMap'></a>func  Map

```go
func Map(items DashSlice, iteratee Iteratee) DashSlice
```
Creates an array of values by running each element in collection thru iteratee.
The iteratee is invoked with one argument: (value).

####  1.63. <a name='funcNewDashSlice'></a>func  NewDashSlice

```go
func NewDashSlice(items ...interface{}) DashSlice
```

####  1.64. <a name='funcNewDashSliceFromIntArray'></a>func  NewDashSliceFromIntArray

```go
func NewDashSliceFromIntArray(items ...int) DashSlice
```

####  1.65. <a name='funcPull'></a>func  Pull

```go
func Pull(items *DashSlice, values ...interface{}) DashSlice
```
Removes all given values from array using SameValueZero for equality
comparisons.

####  1.66. <a name='funcPullAll'></a>func  PullAll

```go
func PullAll(items *DashSlice, values DashSlice) DashSlice
```
This method is like Pull except that it accepts an array of values to remove.

####  1.67. <a name='funcPullAllWith'></a>func  PullAllWith

```go
func PullAllWith(items *DashSlice, values DashSlice, comparison Comparison) DashSlice
```
This method is like PullAll except that it accepts comparator which is invoked
to compare elements of array to values. The comparator is invoked with two
arguments: (arrVal, othVal).

####  1.68. <a name='funcPullAt'></a>func  PullAt

```go
func PullAt(items *DashSlice, indices ...int) DashSlice
```
Removes elements from array corresponding to indexes and returns an array of
removed elements.

####  1.69. <a name='funcReject'></a>func  Reject

```go
func Reject(items DashSlice, predicate Predicate) DashSlice
```
The opposite of Filter; this method returns the elements of collection that
predicate does not return truthy for.

####  1.70. <a name='funcRemove'></a>func  Remove

```go
func Remove(items *DashSlice, predicate Predicate) DashSlice
```
Removes all elements from array that predicate returns truthy for and returns an
array of the removed elements. The predicate is invoked with two arguments:
(value, index).

####  1.71. <a name='funcReverse'></a>func  Reverse

```go
func Reverse(items DashSlice) DashSlice
```
Reverses array so that the first element becomes the last, the second element
becomes the second to last, and so on.

####  1.72. <a name='funcSampleSize'></a>func  SampleSize

```go
func SampleSize(items DashSlice, n int) DashSlice
```
Gets n random elements at unique keys from collection up to the size of
collection.

####  1.73. <a name='funcShuffle'></a>func  Shuffle

```go
func Shuffle(items DashSlice) DashSlice
```
Creates an array of shuffled values.

####  1.74. <a name='funcSlice'></a>func  Slice

```go
func Slice(items DashSlice, start int, end int) DashSlice
```
Creates a slice of array from start up to, but not including, end.

####  1.75. <a name='funcSortByFloat64'></a>func  SortByFloat64

```go
func SortByFloat64(items DashSlice, iteratee IterateeToFloat) DashSlice
```
Creates an array of elements, sorted in ascending order by the results of
running each element in a collection thru each iteratee. This method performs a
stable sort, that is, it preserves the original sort order of equal elements.
The iteratee is invoked with one argument: (value).

####  1.76. <a name='funcSortByInt'></a>func  SortByInt

```go
func SortByInt(items DashSlice, iteratee IterateeToInt) DashSlice
```
Creates an array of elements, sorted in ascending order by the results of
running each element in a collection thru each iteratee. This method performs a
stable sort, that is, it preserves the original sort order of equal elements.
The iteratee is invoked with one argument: (value).

####  1.77. <a name='funcSortByString'></a>func  SortByString

```go
func SortByString(items DashSlice, iteratee IterateeToString) DashSlice
```
Creates an array of elements, sorted in ascending order by the results of
running each element in a collection thru each iteratee. This method performs a
stable sort, that is, it preserves the original sort order of equal elements.
The iteratee is invoked with one argument: (value).

####  1.78. <a name='funcTail'></a>func  Tail

```go
func Tail(items DashSlice) DashSlice
```
Gets all but the first element of array.

####  1.79. <a name='funcTake'></a>func  Take

```go
func Take(items DashSlice, n int) DashSlice
```
Creates a slice of array with n elements taken from the beginning.

####  1.80. <a name='funcTakeRight'></a>func  TakeRight

```go
func TakeRight(items DashSlice, n int) DashSlice
```
Creates a slice of array with n elements taken from the end.

####  1.81. <a name='funcTakeRightWhile'></a>func  TakeRightWhile

```go
func TakeRightWhile(items DashSlice, predicate Predicate) DashSlice
```
Creates a slice of array with elements taken from the end. Elements are taken
until predicate returns falsy. The predicate is invoked with one argument:
(value).

####  1.82. <a name='funcTakeWhile'></a>func  TakeWhile

```go
func TakeWhile(items DashSlice, predicate Predicate) DashSlice
```
Creates a slice of array with elements taken from the beginning. Elements are
taken until predicate returns falsy. The predicate is invoked with one argument:
(value).

####  1.83. <a name='funcUnion'></a>func  Union

```go
func Union(slices ...DashSlice) DashSlice
```
Creates an array of unique values, in order, from all given arrays using
SameValueZero for equality comparisons.

####  1.84. <a name='funcUnionBy'></a>func  UnionBy

```go
func UnionBy(iteratee Iteratee, slices ...DashSlice) DashSlice
```
This method is like Uniq except that it accepts iteratee which is invoked for
each element in array to generate the criterion by which uniqueness is computed.
The order of result values is determined by the order they occur in the array.
The iteratee is invoked with one argument: (value).

####  1.85. <a name='funcUnionWith'></a>func  UnionWith

```go
func UnionWith(comparison Comparison, slices ...DashSlice) DashSlice
```
This method is like Uniq except that it accepts comparator which is invoked to
compare elements of array. The order of result values is determined by the order
they occur in the array. The comparator is invoked with two arguments: (arrVal,
othVal).

####  1.86. <a name='funcUniq'></a>func  Uniq

```go
func Uniq(items DashSlice) DashSlice
```
Creates a duplicate-free version of an array, using SameValueZero for equality
comparisons, in which only the first occurrence of each element is kept. The
order of result values is determined by the order they occur in the array.

####  1.87. <a name='funcUniqBy'></a>func  UniqBy

```go
func UniqBy(items DashSlice, iteratee Iteratee) DashSlice
```
This method is like Union except that it accepts iteratee which is invoked for
each element of each arrays to generate the criterion by which uniqueness is
computed. Result values are chosen from the first array in which the value
occurs. The iteratee is invoked with one argument: (value).

####  1.88. <a name='funcUniqWith'></a>func  UniqWith

```go
func UniqWith(items DashSlice, comparison Comparison) DashSlice
```
This method is like Uniq except that it accepts comparator which is invoked to
compare elements of array. The order of result values is determined by the order
they occur in the array. The comparator is invoked with two arguments: (arrVal,
othVal).

####  1.89. <a name='funcWithout'></a>func  Without

```go
func Without(items DashSlice, values ...interface{}) DashSlice
```
Creates an array excluding all given values using SameValueZero for equality
comparisons.

####  1.90. <a name='funcXor'></a>func  Xor

```go
func Xor(items ...DashSlice) DashSlice
```
Creates an array of unique values that is the symmetric difference of the given
arrays. The order of result values is determined by the order they occur in the
arrays.

####  1.91. <a name='funcZip'></a>func  Zip

```go
func Zip(slices ...DashSlice) []DashSlice
```
Creates an array of grouped elements, the first of which contains the first
elements of the given arrays, the second of which contains the second elements
of the given arrays, and so on.

####  1.92. <a name='funcZipWith'></a>func  ZipWith

```go
func ZipWith(iteratee func([]interface{}) interface{}, slices ...DashSlice) DashSlice
```
This method is like Zip except that it accepts iteratee to specify how grouped
values should be combined. The iteratee is invoked with the elements of each
group: (...group).

####  1.93. <a name='funcDashSliceMap'></a>func (DashSlice) Map

```go
func (ds DashSlice) Map(iteratee func(interface{}) interface{}) DashSlice
```

####  1.94. <a name='typeIteratee'></a>type Iteratee

```go
type Iteratee func(interface{}) interface{}
```


####  1.95. <a name='typeIterateeToFloat'></a>type IterateeToFloat

```go
type IterateeToFloat func(interface{}) float64
```


####  1.96. <a name='typeIterateeToInt'></a>type IterateeToInt

```go
type IterateeToInt func(interface{}) int
```


####  1.97. <a name='typeIterateeToString'></a>type IterateeToString

```go
type IterateeToString func(interface{}) string
```


####  1.98. <a name='typePredicate'></a>type Predicate

```go
type Predicate func(interface{}) bool
```


####  1.99. <a name='typeReducer'></a>type Reducer

```go
type Reducer func(interface{}, interface{}) interface{}
```
