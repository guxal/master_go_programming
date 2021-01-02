# Slices in Go

### Array 
- Has a **fixed length** defined at compile time;

- The length of an array is part of its type, defined at compile time and cannot be changed;

- By default an uninitialized array has all elements equal to zero;

### Slice
- Has a **dynamic length** (it can shrink ir grow);
- The length of a slice is not part of its type and it belongs to runtime;
- An uninitialized slice is equal to nil (its zero value is nil).


### Array & Slice 

- Both a slice and an array can contain only the same type of elements;
- We can create a keyed slice like a keyed array;