# Slice's Backing (Underlying) Array

- When creating a slice, behind the scenes Go creates a hidden array called Backing Array.
- The backing array stores the elements, not the slice.
- Go implements a slice as a data structure called slice header.


### **Slice Header** contains 3 fields:

1. *the address* of the backing array (pointer)
2. *the length* of the slice. len() returns it.
3. *the capacity* of the slice. cap() returns it.

- Slice Header is the runtime representation of a slice.
- A **nil slice** doesn't a have backing array