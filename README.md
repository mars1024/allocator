# Allocator

A universal allocator framework and some allocator implementations

## Universal interfaces

Allocator interface is as below, it has 4 universal methods,
- `Assign` is for assigning the using objects in range
- `Allocate` is for allocating an valid object from range
- `Release` is for releasing an object from range
- `Has` is for testing whether an object is in range

```go
type Interface interface {
	Assign(RangeID) error
	Allocate() (RangeID, interface{}, error)
	Release(RangeID) error
	Has(RangeID) bool
}
```

Range interface is as below, it has 2 universal methods,
- `Contains` is for testing whether an object is in range
- `First` is for returning an iterator which can traverse the whole range

```go
type Range interface {
	Contains(RangeID) bool
	First() RangeIterator
}
```

RangeIterator interface is as below, it has 3 universal methods,
the 3 methods can help us traverse the whole range.

```go
type RangeIterator interface {
	Get() (RangeID, interface{})
	Next()
	InRange() bool
}
```

## Universal implementation

There is an universal implementation to generate an Allocator from an Range.
You can use `func NewAllocator(ranger Range) Interface {}` method to get this.

## Some common allocators

There are also some common allocators implemented through the interfaces and methods above.

- [Port Allocator](https://github.com/mars1024/allocator/tree/master/allocators/port)