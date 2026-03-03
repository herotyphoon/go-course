# Maps

- Maps are data structures used to store key-value pairs.
- Behind the scenes, Go uses hashmaps to implement maps.
- The `len` function can be used to get the key value pairs in a map.

## Example

```go
nameAge := map[string]int{
    "Foo": 25,
    "Bar": 20,
}

firstNameLastName := map[string]string{
    "Foo": "Bar",
    "Bar": "Foo",
}
```

## Nil Map

- The zero value for a null map is `nil`.
- A `nil` map can be created using a variable declaration.
- Eg: `var nameAge map[string]int`
- Length of a `nil` map is `0`.
- Reading from the `nil` map will return zero value for the value type.
- Writing to a `nil` map will cause panic.

## Empty Map

- Empty maps can be created by many ways .
- Length of an empty map is 0.
- You can read and write to an empty map.

## Examples

```go
nameAge := map[string]int{}
nameAge := make(map[string]int)
var nameAge map[string]int = map[string]int{}
```

## Read From a Map

- Value from a `map` can be accessed using the key.
- Example: `nameAge["Foobar"]`
- If key does not exist, returns the zero value.
- The comma ok idiom in Go:

```go
fooAge, ok := nameAge["foo"]
if !ok { :
fmt.Println("not found")
}
```

## Write to a Map

- Value can be written to a `map` using a key.
- Example: `nameAge["Foo"] = 20`
- A `map` can have unique keys only.
- Using duplicate keys will overwrite any existing key value pair.

## Some More Facts

- Map keys can only be of comparable types.
- Values can be of any type.
- Map can only be compared to `nil`.
- Since `map` keys are unique, maps can be used as `sets`.
