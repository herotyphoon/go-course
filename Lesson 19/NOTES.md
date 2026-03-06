# Structs

- Go provide `struct` to allow you to create your own custom data types.
- A `struct` helps group similar data together.
- A `struct` is created using the `type` and `struct` keywords.

## Example

```go
type student struct {
  firstName string
  lastName string
  age int
  classes []string
}
```

## Zero Value for a Struct

- Zero value for a struct is a struct with all its fields having zero values for their corresponding types.
- A zero value struct can be created using a literal declaration: `var foo student`.
- It can also be created with a variable shorthand: `foo := student{}`.

## Initializing Field Values

- There are two ways to initialize `struct` field values during struct initialization.
- The first is implicit initialization, where values are provided inside the curly braces without field names.
- The second is explicit initialization, where values are provided inside the curly braces with field names.

## Example

```go
foo := student{
  "hero",
  "typhoon",
  25,
  []string{"maths", "english"}
}

foo := student{
  firstName: "hero",
  lastName: "typhoon"
}
```

## The Dot Notation

- The dot notation is used to read and write field values of a struct.
- For examples, for a stucent named foo, you can read the values with `foo.firstName`.
- Similarly to update the value you can do `foo.firstName = "legend"`.

