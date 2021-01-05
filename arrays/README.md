# Arrays in Go

- An array is a composite, indexable type that stores a collection of elements of same type.

- An array has a fixed length.
- Every element in an array (or slice) must be of same type.
- Go stores the elements of the array in contiguos memory locations and this way it's very efficient.

- The length and the elements type determines the type of an array. The Length belongs to array type and it's determine at compile time.

```go
accounts := [3]int {50, 60, 70}

```

The array called accounts that consists of 3 integers has it's type [3]int.