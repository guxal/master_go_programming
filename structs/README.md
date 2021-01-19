# Structs in Go

- A `struct` is a sequence of **named elements**, called **fields**. Each of them has a name and a
type.
- If you are familiar with OOP from other languages **you can think of a struct as of a class**.
The struct fields are like the instance attributes we define in OOP.
- Unlike traditional Object-Oriented Programming, Go does not have a class-object
architecture. Rather we have structs which hold complex data structures.
- A structs is nothing more that a schema containing a blueprint of data a structure will
hold. This blueprint is fixed at compile time. It’s not allowed to change the name or the
type of the fields at runtime. You can’t add or remove fields from a struct at runtime.