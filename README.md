# pkg
--
    import "."

This is lodash golang version. A modern Golang utility library delivering
modularity, performance & extras.

## Usage

#### func  CountBy

```go
func CountBy(items DashSlice, iteratee Iteratee) map[interface{}]int
```
Creates an object composed of keys generated from the results of running each
element of collection thru iteratee. The corresponding value of each key is the
number of times the key was returned by iteratee. The iteratee is invoked with
one argument: (value).

#### func  Every

```go
func Every(items DashSlice, predicate Predicate) bool
```
Checks if predicate returns truthy for all elements of collection. Iteration is
stopped once predicate returns falsy. The predicate is invoked with one
argument: (value).

#### func  Fill

```go
func Fill(items DashSlice, fillValue interface{})
```
Fills elements of array with value.

#### func  FillInRange

```go
func FillInRange(items DashSlice, value interface{}, start int, end int)
```
Fills elements of array with value from start up to, but not including end.

#### func  Find

```go
func Find(items DashSlice, predicate Predicate) (interface{}, bool)
```
Iterates over elements of collection, returning the first element predicate
returns truthy for. The predicate is invoked with one argument: (value).

#### func  FindFrom

```go
func FindFrom(items DashSlice, predicate Predicate, start int) (interface{}, bool)
```
Iterates over elements of collection from start index, returning the first
element predicate returns truthy for. The predicate is invoked with one
argument: (value).

#### func  FindIndex

```go
func FindIndex(items DashSlice, predicate Predicate) (int, bool)
```
This method is like Find except that it returns the index of the first element
predicate returns truthy for instead of the element itself.

#### func  FindIndexWith

```go
func FindIndexWith(items DashSlice, element interface{}, comparison Comparison) (int, bool)
```
Same to IndexOf. The difference is that, this method provides a comparison
function to compare programmatically.

#### func  FindLast

```go
func FindLast(items DashSlice, predicate Predicate) (interface{}, bool)
```
This method is like Find except that it iterates over elements of collection
from right to left.

#### func  FindLastFrom

```go
func FindLastFrom(items DashSlice, predicate Predicate, start int) (interface{}, bool)
```
This method is like FindFrom except that it iterates over elements of collection
from right to left.

#### func  FindLastIndex

```go
func FindLastIndex(items DashSlice, predicate Predicate) (int, bool)
```
This method is like Find except that it returns the index of the first element
predicate returns truthy for instead of the element itself.

#### func  FindLastIndexWith

```go
func FindLastIndexWith(items DashSlice, element interface{}, comparison Comparison) (int, bool)
```
This method is like FindIndex except that it iterates over elements of
collection from right to left.

#### func  First

```go
func First(items DashSlice) interface{}
```
Gets the first element of array.

#### func  FromPairs

```go
func FromPairs(pairs []DashSlice) map[interface{}]interface{}
```
This method returns an object composed from key-value pairs.

#### func  GroupBy

```go
func GroupBy(items DashSlice, iteratee Iteratee) map[interface{}]DashSlice
```
Creates an object composed of keys generated from the results of running each
element of collection thru iteratee. The order of grouped values is determined
by the order they occur in collection. The corresponding value of each key is an
array of elements responsible for generating the key. The iteratee is invoked
with one argument: (value).

#### func  Head

```go
func Head(items DashSlice) interface{}
```
Gets the first element of slice.

#### func  Identity

```go
func Identity(i interface{}) interface{}
```

#### func  IdentityFloat64

```go
func IdentityFloat64(i interface{}) float64
```

#### func  IdentityInt

```go
func IdentityInt(i interface{}) int
```

#### func  IdentityString

```go
func IdentityString(i interface{}) string
```

#### func  Includes

```go
func Includes(items DashSlice, value interface{}) bool
```
Checks if value is in collection. If collection is a string, it's checked for a
substring of value, otherwise SameValueZero is used for equality comparisons. If
fromIndex is negative, it's used as the offset from the end of collection.

#### func  IndexOf

```go
func IndexOf(items DashSlice, element interface{}) (int, bool)
```
This method is like _.find except that it returns the index of the first element
predicate returns truthy for instead of the element itself.

#### func  Join

```go
func Join(items DashSlice, separator string) string
```
Converts all elements in array into a string separated by separator.

#### func  Last

```go
func Last(items DashSlice) interface{}
```
Gets the last element of array.

#### func  LastIndexOf

```go
func LastIndexOf(items DashSlice, element interface{}) (int, bool)
```
This method is like IndexOf except that it iterates over elements of array from
right to left.

#### func  Nth

```go
func Nth(items DashSlice, n int) interface{}
```
Gets the element at index n of array. If n is negative, the nth element from the
end is returned.

#### func  Reduce

```go
func Reduce(items DashSlice, reducer Reducer) interface{}
```
Reduces collection to a value which is the accumulated result of running each
element in collection thru iteratee, where each successive invocation is
supplied the return value of the previous. If accumulator is not given, the
first element of collection is used as the initial value. The reducer is invoked
with two arguments: (accumulator, value).

#### func  ReduceRight

```go
func ReduceRight(items DashSlice, reducer Reducer) interface{}
```
This method is like Reduce except that it iterates over elements of collection
from right to left.

#### func  ReduceRightWithInitial

```go
func ReduceRightWithInitial(items DashSlice, reducer Reducer, initial interface{}) interface{}
```
This method is like ReduceWithInitial except that it iterates over elements of
collection from right to left.

#### func  ReduceWithInitial

```go
func ReduceWithInitial(items DashSlice, reducer Reducer, initial interface{}) interface{}
```
Reduces collection to a value which is the accumulated result of running each
element in collection thru iteratee, where each successive invocation is
supplied the return value of the previous. If accumulator is not given, the
first element of collection is used as the initial value. The reducer is invoked
with two arguments: (accumulator, value).

#### func  Sample

```go
func Sample(items DashSlice) interface{}
```
Gets a random element from collection.

#### func  Size

```go
func Size(items DashSlice) int
```
Gets the size of collection.

#### func  Some

```go
func Some(items DashSlice, predicate Predicate) bool
```
Checks if predicate returns truthy for any element of collection. Iteration is
stopped once predicate returns truthy. The predicate is invoked with one
argument: (value).

#### type Action

```go
type Action func(interface{}, int)
```


#### type Comparison

```go
type Comparison func(interface{}, interface{}) bool
```


#### type DashSlice

```go
type DashSlice []interface{}
```


#### func  Chunk

```go
func Chunk(items DashSlice, size int) []DashSlice
```
Creates an array of elements split into groups the length of size. If array
can't be split evenly, the final chunk will be the remaining elements.

#### func  Compact

```go
func Compact(items DashSlice) DashSlice
```
Creates an array with all falsy values removed. The values false, 0, "", nil are
falsy.

#### func  Concat

```go
func Concat(items DashSlice, newItems ...interface{}) DashSlice
```
Creates a new array concatenating array with any additional arrays and/or
values.

#### func  ConcatSlices

```go
func ConcatSlices(slices ...DashSlice) DashSlice
```
Creates a new array concatenating array with any additional DashSlices.

#### func  Difference

```go
func Difference(items DashSlice, itemsToCompare ...interface{}) DashSlice
```
Creates an array of array values not included in the other given arrays using
SameValueZero for equality comparisons. The order and references of result
values are determined by the first array.

#### func  DifferenceBy

```go
func DifferenceBy(items DashSlice, itemsToCompare DashSlice, iteratee Iteratee) DashSlice
```
This method is like _.difference except that it accepts iteratee which is
invoked for each element of array and values to generate the criterion by which
they're compared. The order and references of result values are determined by
the first array. The iteratee is invoked with one argument: (value).

#### func  DifferenceWith

```go
func DifferenceWith(items DashSlice, itemsToCompare DashSlice,
	comparison Comparison) DashSlice
```
This method is like _.difference except that it accepts comparator which is
invoked to compare elements of array to values. The order and references of
result values are determined by the first array. The comparator is invoked with
two arguments: (arrVal, othVal).

#### func  Drop

```go
func Drop(items DashSlice, count int) DashSlice
```
Creates a slice of array with n elements dropped from the beginning.

#### func  DropRight

```go
func DropRight(items DashSlice, count int) DashSlice
```
Creates a slice of array with n elements dropped from the end.

#### func  DropWhile

```go
func DropWhile(items DashSlice, predicate Predicate) DashSlice
```
Creates a slice of array excluding elements dropped from the beginning. Elements
are dropped until predicate returns falsy. The predicate is invoked with two
arguments: (value, index).

#### func  Each

```go
func Each(items DashSlice, action Action) DashSlice
```
Iterates over elements of collection and invokes iteratee for each element. The
iteratee is invoked with three arguments: (value, index|key, collection).
Iteratee functions may exit iteration early by explicitly returning false.

#### func  EachRight

```go
func EachRight(items DashSlice, action Action) DashSlice
```
This method is like ForEach except that it iterates over elements of collection
from right to left.

#### func  Filter

```go
func Filter(items DashSlice, predicate Predicate) DashSlice
```
Iterates over elements of collection, returning an array of all elements
predicate returns truthy for. The predicate is invoked with one argument:
(value).

#### func  FlatMap

```go
func FlatMap(items DashSlice, iteratee Iteratee) DashSlice
```
Creates a flattened array of values by running each element in collection thru
iteratee and flattening the mapped results. The iteratee is invoked with one
argument: (value).

#### func  FlatMapDeep

```go
func FlatMapDeep(items DashSlice, iteratee Iteratee) DashSlice
```
This method is like FlatMap except that it recursively flattens the mapped
results.

#### func  FlatMapDepth

```go
func FlatMapDepth(items DashSlice, iteratee Iteratee, depth int) DashSlice
```
This method is like FlatMap except that it recursively flattens the mapped
results up to depth times.

#### func  Flatten

```go
func Flatten(items DashSlice) DashSlice
```
Flattens array a single level deep.

#### func  FlattenDeep

```go
func FlattenDeep(items DashSlice) DashSlice
```
Recursively flattens array.

#### func  FlattenDepth

```go
func FlattenDepth(items DashSlice, depth int) DashSlice
```
Recursively flatten array up to depth times.

#### func  ForEach

```go
func ForEach(items DashSlice, action Action) DashSlice
```
Iterates over elements of collection and invokes iteratee for each element. The
iteratee is invoked with three arguments: (value, index|key, collection).
Iteratee functions may exit iteration early by explicitly returning false.

#### func  ForEachRight

```go
func ForEachRight(items DashSlice, action Action) DashSlice
```
This method is like ForEach except that it iterates over elements of collection
from right to left.

#### func  Initial

```go
func Initial(items DashSlice) DashSlice
```
Gets all but the last element of array.

#### func  Intersection

```go
func Intersection(items1 DashSlice, items2 DashSlice) DashSlice
```
Creates an array of unique values that are included in all given arrays using
SameValueZero for equality comparisons. The order and references of result
values are determined by the first array.

#### func  IntersectionBy

```go
func IntersectionBy(items1 DashSlice, items2 DashSlice, iteratee Iteratee) DashSlice
```
This method is like Intersection except that it accepts iteratee which is
invoked for each element of each arrays to generate the criterion by which
they're compared. The order and references of result values are determined by
the first array. The iteratee is invoked with one argument: (value).

#### func  IntersectionWith

```go
func IntersectionWith(items1 DashSlice, items2 DashSlice, comparison Comparison) DashSlice
```
This method is like _.intersection except that it accepts comparator which is
invoked to compare elements of arrays. The order and references of result values
are determined by the first array. The comparator is invoked with two arguments:
(arrVal, othVal).

#### func  Map

```go
func Map(items DashSlice, iteratee Iteratee) DashSlice
```
Creates an array of values by running each element in collection thru iteratee.
The iteratee is invoked with one argument: (value).

#### func  NewDashSlice

```go
func NewDashSlice(items ...interface{}) DashSlice
```

#### func  NewDashSliceFromIntArray

```go
func NewDashSliceFromIntArray(items ...int) DashSlice
```

#### func  Pull

```go
func Pull(items *DashSlice, values ...interface{}) DashSlice
```
Removes all given values from array using SameValueZero for equality
comparisons.

#### func  PullAll

```go
func PullAll(items *DashSlice, values DashSlice) DashSlice
```
This method is like Pull except that it accepts an array of values to remove.

#### func  PullAllWith

```go
func PullAllWith(items *DashSlice, values DashSlice, comparison Comparison) DashSlice
```
This method is like PullAll except that it accepts comparator which is invoked
to compare elements of array to values. The comparator is invoked with two
arguments: (arrVal, othVal).

#### func  PullAt

```go
func PullAt(items *DashSlice, indices ...int) DashSlice
```
Removes elements from array corresponding to indexes and returns an array of
removed elements.

#### func  Reject

```go
func Reject(items DashSlice, predicate Predicate) DashSlice
```
The opposite of Filter; this method returns the elements of collection that
predicate does not return truthy for.

#### func  Remove

```go
func Remove(items *DashSlice, predicate Predicate) DashSlice
```
Removes all elements from array that predicate returns truthy for and returns an
array of the removed elements. The predicate is invoked with two arguments:
(value, index).

#### func  Reverse

```go
func Reverse(items DashSlice) DashSlice
```
Reverses array so that the first element becomes the last, the second element
becomes the second to last, and so on.

#### func  SampleSize

```go
func SampleSize(items DashSlice, n int) DashSlice
```
Gets n random elements at unique keys from collection up to the size of
collection.

#### func  Shuffle

```go
func Shuffle(items DashSlice) DashSlice
```
Creates an array of shuffled values.

#### func  Slice

```go
func Slice(items DashSlice, start int, end int) DashSlice
```
Creates a slice of array from start up to, but not including, end.

#### func  SortByFloat64

```go
func SortByFloat64(items DashSlice, iteratee IterateeToFloat) DashSlice
```
Creates an array of elements, sorted in ascending order by the results of
running each element in a collection thru each iteratee. This method performs a
stable sort, that is, it preserves the original sort order of equal elements.
The iteratee is invoked with one argument: (value).

#### func  SortByInt

```go
func SortByInt(items DashSlice, iteratee IterateeToInt) DashSlice
```
Creates an array of elements, sorted in ascending order by the results of
running each element in a collection thru each iteratee. This method performs a
stable sort, that is, it preserves the original sort order of equal elements.
The iteratee is invoked with one argument: (value).

#### func  SortByString

```go
func SortByString(items DashSlice, iteratee IterateeToString) DashSlice
```
Creates an array of elements, sorted in ascending order by the results of
running each element in a collection thru each iteratee. This method performs a
stable sort, that is, it preserves the original sort order of equal elements.
The iteratee is invoked with one argument: (value).

#### func  Tail

```go
func Tail(items DashSlice) DashSlice
```
Gets all but the first element of array.

#### func  Take

```go
func Take(items DashSlice, n int) DashSlice
```
Creates a slice of array with n elements taken from the beginning.

#### func  TakeRight

```go
func TakeRight(items DashSlice, n int) DashSlice
```
Creates a slice of array with n elements taken from the end.

#### func  TakeRightWhile

```go
func TakeRightWhile(items DashSlice, predicate Predicate) DashSlice
```
Creates a slice of array with elements taken from the end. Elements are taken
until predicate returns falsy. The predicate is invoked with one argument:
(value).

#### func  TakeWhile

```go
func TakeWhile(items DashSlice, predicate Predicate) DashSlice
```
Creates a slice of array with elements taken from the beginning. Elements are
taken until predicate returns falsy. The predicate is invoked with one argument:
(value).

#### func  Union

```go
func Union(slices ...DashSlice) DashSlice
```
Creates an array of unique values, in order, from all given arrays using
SameValueZero for equality comparisons.

#### func  UnionBy

```go
func UnionBy(iteratee Iteratee, slices ...DashSlice) DashSlice
```
This method is like Uniq except that it accepts iteratee which is invoked for
each element in array to generate the criterion by which uniqueness is computed.
The order of result values is determined by the order they occur in the array.
The iteratee is invoked with one argument: (value).

#### func  UnionWith

```go
func UnionWith(comparison Comparison, slices ...DashSlice) DashSlice
```
This method is like Uniq except that it accepts comparator which is invoked to
compare elements of array. The order of result values is determined by the order
they occur in the array. The comparator is invoked with two arguments: (arrVal,
othVal).

#### func  Uniq

```go
func Uniq(items DashSlice) DashSlice
```
Creates a duplicate-free version of an array, using SameValueZero for equality
comparisons, in which only the first occurrence of each element is kept. The
order of result values is determined by the order they occur in the array.

#### func  UniqBy

```go
func UniqBy(items DashSlice, iteratee Iteratee) DashSlice
```
This method is like Union except that it accepts iteratee which is invoked for
each element of each arrays to generate the criterion by which uniqueness is
computed. Result values are chosen from the first array in which the value
occurs. The iteratee is invoked with one argument: (value).

#### func  UniqWith

```go
func UniqWith(items DashSlice, comparison Comparison) DashSlice
```
This method is like Uniq except that it accepts comparator which is invoked to
compare elements of array. The order of result values is determined by the order
they occur in the array. The comparator is invoked with two arguments: (arrVal,
othVal).

#### func  Without

```go
func Without(items DashSlice, values ...interface{}) DashSlice
```
Creates an array excluding all given values using SameValueZero for equality
comparisons.

#### func  Xor

```go
func Xor(items ...DashSlice) DashSlice
```
Creates an array of unique values that is the symmetric difference of the given
arrays. The order of result values is determined by the order they occur in the
arrays.

#### func  Zip

```go
func Zip(slices ...DashSlice) []DashSlice
```
Creates an array of grouped elements, the first of which contains the first
elements of the given arrays, the second of which contains the second elements
of the given arrays, and so on.

#### func  ZipWith

```go
func ZipWith(iteratee func([]interface{}) interface{}, slices ...DashSlice) DashSlice
```
This method is like Zip except that it accepts iteratee to specify how grouped
values should be combined. The iteratee is invoked with the elements of each
group: (...group).

#### func (DashSlice) Map

```go
func (ds DashSlice) Map(iteratee func(interface{}) interface{}) DashSlice
```

#### type Iteratee

```go
type Iteratee func(interface{}) interface{}
```


#### type IterateeToFloat

```go
type IterateeToFloat func(interface{}) float64
```


#### type IterateeToInt

```go
type IterateeToInt func(interface{}) int
```


#### type IterateeToString

```go
type IterateeToString func(interface{}) string
```


#### type Predicate

```go
type Predicate func(interface{}) bool
```


#### type Reducer

```go
type Reducer func(interface{}, interface{}) interface{}
```
