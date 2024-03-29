[![PkgGoDev](https://pkg.go.dev/badge/github.com/ginkgoch/godash)](https://pkg.go.dev/github.com/ginkgoch/godash/pkg)
[![Go Report Card](https://goreportcard.com/badge/github.com/ginkgoch/godash)](https://goreportcard.com/report/github.com/ginkgoch/godash)

[![forthebadge](https://forthebadge.com/images/badges/gluten-free.svg)](https://forthebadge.com)

# Lodash in Golang Version 
    

This is lodash in golang version. A modern Golang utility library delivering
modularity, performance & extras.

## Install & Import
```bash
go get -u "github.com/ginkgoch/godash/v1"
```

```go
import "github.com/ginkgoch/godash/v1"
```

## Example

* Slice Example

```go
items := godash.DashSlice{"a", "b", "c", "d"}
chunked := godash.Chunk(items, 2)

// out: [[a b] [c d]]
```

* Collection Example

```go
items := godash.DashSlice{1, 2, 3, 4, 5}
result := godash.Shuffle(items)

// out: [2 5 4 3 1]
```

* More examples TBD...

## Index
This is a index of the [complete API reference](REFERENCE.md).
<!-- vscode-markdown-toc -->
* [func  CamelCase](REFERENCE.md#funcCamelCase)
* [func  CamelCaseWithInit](REFERENCE.md#funcCamelCaseWithInit)
* [func  Capitalize](REFERENCE.md#funcCapitalize)
* [func  CountBy](REFERENCE.md#funcCountBy)
* [func  EndsWith](REFERENCE.md#funcEndsWith)
* [func  EndsWithFrom](REFERENCE.md#funcEndsWithFrom)
* [func  Escape](REFERENCE.md#funcEscape)
* [func  EscapeRegExp](REFERENCE.md#funcEscapeRegExp)
* [func  Every](REFERENCE.md#funcEvery)
* [func  Fill](REFERENCE.md#funcFill)
* [func  FillInRange](REFERENCE.md#funcFillInRange)
* [func  Find](REFERENCE.md#funcFind)
* [func  FindFrom](REFERENCE.md#funcFindFrom)
* [func  FindIndex](REFERENCE.md#funcFindIndex)
* [func  FindIndexWith](REFERENCE.md#funcFindIndexWith)
* [func  FindLast](REFERENCE.md#funcFindLast)
* [func  FindLastFrom](REFERENCE.md#funcFindLastFrom)
* [func  FindLastIndex](REFERENCE.md#funcFindLastIndex)
* [func  FindLastIndexWith](REFERENCE.md#funcFindLastIndexWith)
* [func  First](REFERENCE.md#funcFirst)
* [func  FromPairs](REFERENCE.md#funcFromPairs)
* [func  GroupBy](REFERENCE.md#funcGroupBy)
* [func  Head](REFERENCE.md#funcHead)
* [func  Identity](REFERENCE.md#funcIdentity)
* [func  IdentityFloat](REFERENCE.md#funcIdentityFloat)
* [func  IdentityInt](REFERENCE.md#funcIdentityInt)
* [func  IdentityString](REFERENCE.md#funcIdentityString)
* [func  Includes](REFERENCE.md#funcIncludes)
* [func  IndexOf](REFERENCE.md#funcIndexOf)
* [func  Join](REFERENCE.md#funcJoin)
* [func  Last](REFERENCE.md#funcLast)
* [func  LastIndexOf](REFERENCE.md#funcLastIndexOf)
* [func  LowerFirst](REFERENCE.md#funcLowerFirst)
* [func  Nth](REFERENCE.md#funcNth)
* [func  Pad](REFERENCE.md#funcPad)
* [func  PadLeft](REFERENCE.md#funcPadLeft)
* [func  PadLeftWith](REFERENCE.md#funcPadLeftWith)
* [func  PadRight](REFERENCE.md#funcPadRight)
* [func  PadRightWith](REFERENCE.md#funcPadRightWith)
* [func  PadWith](REFERENCE.md#funcPadWith)
* [func  Reduce](REFERENCE.md#funcReduce)
* [func  ReduceRight](REFERENCE.md#funcReduceRight)
* [func  ReduceRightWithInitial](REFERENCE.md#funcReduceRightWithInitial)
* [func  ReduceWithInitial](REFERENCE.md#funcReduceWithInitial)
* [func  Repeat](REFERENCE.md#funcRepeat)
* [func  Replace](REFERENCE.md#funcReplace)
* [func  ReplaceRegx](REFERENCE.md#funcReplaceRegx)
* [func  Sample](REFERENCE.md#funcSample)
* [func  Size](REFERENCE.md#funcSize)
* [func  Some](REFERENCE.md#funcSome)
* [func  Split](REFERENCE.md#funcSplit)
* [func  SplitWithCountLimit](REFERENCE.md#funcSplitWithCountLimit)
* [func  StartsWith](REFERENCE.md#funcStartsWith)
* [func  StartsWithFrom](REFERENCE.md#funcStartsWithFrom)
* [func  Ternary](REFERENCE.md#funcTernary)
* [func  ToLower](REFERENCE.md#funcToLower)
* [func  ToUpper](REFERENCE.md#funcToUpper)
* [func  Trim](REFERENCE.md#funcTrim)
* [func  TrimEnd](REFERENCE.md#funcTrimEnd)
* [func  TrimEndWith](REFERENCE.md#funcTrimEndWith)
* [func  TrimStart](REFERENCE.md#funcTrimStart)
* [func  TrimStartWith](REFERENCE.md#funcTrimStartWith)
* [func  TrimWith](REFERENCE.md#funcTrimWith)
* [func  Unescape](REFERENCE.md#funcUnescape)
* [func  UpperFirst](REFERENCE.md#funcUpperFirst)
* [type Action](REFERENCE.md#typeAction)
* [type Comparison](REFERENCE.md#typeComparison)
* [type DashSlice](REFERENCE.md#typeDashSlice)
* [func  Chunk](REFERENCE.md#funcChunk)
* [func  Compact](REFERENCE.md#funcCompact)
* [func  Concat](REFERENCE.md#funcConcat)
* [func  ConcatSlices](REFERENCE.md#funcConcatSlices)
* [func  Difference](REFERENCE.md#funcDifference)
* [func  DifferenceBy](REFERENCE.md#funcDifferenceBy)
* [func  DifferenceWith](REFERENCE.md#funcDifferenceWith)
* [func  Drop](REFERENCE.md#funcDrop)
* [func  DropRight](REFERENCE.md#funcDropRight)
* [func  DropWhile](REFERENCE.md#funcDropWhile)
* [func  Each](REFERENCE.md#funcEach)
* [func  EachRight](REFERENCE.md#funcEachRight)
* [func  Filter](REFERENCE.md#funcFilter)
* [func  FlatMap](REFERENCE.md#funcFlatMap)
* [func  FlatMapDeep](REFERENCE.md#funcFlatMapDeep)
* [func  FlatMapDepth](REFERENCE.md#funcFlatMapDepth)
* [func  Flatten](REFERENCE.md#funcFlatten)
* [func  FlattenDeep](REFERENCE.md#funcFlattenDeep)
* [func  FlattenDepth](REFERENCE.md#funcFlattenDepth)
* [func  ForEach](REFERENCE.md#funcForEach)
* [func  ForEachRight](REFERENCE.md#funcForEachRight)
* [func  Initial](REFERENCE.md#funcInitial)
* [func  Intersection](REFERENCE.md#funcIntersection)
* [func  IntersectionBy](REFERENCE.md#funcIntersectionBy)
* [func  IntersectionWith](REFERENCE.md#funcIntersectionWith)
* [func  Map](REFERENCE.md#funcMap)
* [func  NewDashSlice](REFERENCE.md#funcNewDashSlice)
* [func  NewDashSliceFromIntArray](REFERENCE.md#funcNewDashSliceFromIntArray)
* [func  Pull](REFERENCE.md#funcPull)
* [func  PullAll](REFERENCE.md#funcPullAll)
* [func  PullAllWith](REFERENCE.md#funcPullAllWith)
* [func  PullAt](REFERENCE.md#funcPullAt)
* [func  Reject](REFERENCE.md#funcReject)
* [func  Remove](REFERENCE.md#funcRemove)
* [func  Reverse](REFERENCE.md#funcReverse)
* [func  SampleSize](REFERENCE.md#funcSampleSize)
* [func  Shuffle](REFERENCE.md#funcShuffle)
* [func  Slice](REFERENCE.md#funcSlice)
* [func  SortByFloat](REFERENCE.md#funcSortByFloat)
* [func  SortByInt](REFERENCE.md#funcSortByInt)
* [func  SortByString](REFERENCE.md#funcSortByString)
* [func  Tail](REFERENCE.md#funcTail)
* [func  Take](REFERENCE.md#funcTake)
* [func  TakeRight](REFERENCE.md#funcTakeRight)
* [func  TakeRightWhile](REFERENCE.md#funcTakeRightWhile)
* [func  TakeWhile](REFERENCE.md#funcTakeWhile)
* [func  Union](REFERENCE.md#funcUnion)
* [func  UnionBy](REFERENCE.md#funcUnionBy)
* [func  UnionWith](REFERENCE.md#funcUnionWith)
* [func  Uniq](REFERENCE.md#funcUniq)
* [func  UniqBy](REFERENCE.md#funcUniqBy)
* [func  UniqWith](REFERENCE.md#funcUniqWith)
* [func  Without](REFERENCE.md#funcWithout)
* [func  Xor](REFERENCE.md#funcXor)
* [func  Zip](REFERENCE.md#funcZip)
* [func  ZipWith](REFERENCE.md#funcZipWith)
* [func (DashSlice) Map](REFERENCE.md#funcDashSliceMap)
* [type InitCamelCase](REFERENCE.md#typeInitCamelCase)
* [type Iteratee](REFERENCE.md#typeIteratee)
* [type IterateeToFloat](REFERENCE.md#typeIterateeToFloat)
* [type IterateeToInt](REFERENCE.md#typeIterateeToInt)
* [type IterateeToString](REFERENCE.md#typeIterateeToString)
* [type Predicate](REFERENCE.md#typePredicate)
* [type Reducer](REFERENCE.md#typeReducer)

<!-- vscode-markdown-toc-config
	numbering=true
	autoSave=true
	/vscode-markdown-toc-config -->
<!-- /vscode-markdown-toc -->
