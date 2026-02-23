# Slices

- Slices are wrappers around arrays.
- They manage the underlying arrays for you.
- They help you overcome the limitations of arrays.
- Length is not attached to the type of the slice.
- Slices have length and capacity.
- Length of a slice is the length of elements stored in a slice (use len).
- Capacity is total size of the underlying array (use cap).

## Behind the Scenes

- When you create a slice of certain length underlying array is usually created of the double length.
- When you reach the capacity of a slice, it will create a array of 50 to 25 % larger than the existing array and copy values from old to new array.

## Declaring Slices in Go

- Declaring slice:
  - `var a []int`
  - `a := []int{1, 2, 3, 4, 5}`
  - `a := []int{10, 2:15, 20}`
  - `a := make([]int, 0)`

## Make Function

- The universal `make` function can be used to create slices.
- It takes type of slice as the first argument, length as second and capacity as an optional third argument.
- Eg:
  - `a := make([]int, 5)`
  - `a := make([]string, 5, 10)`
  - `a := make([]int, 0, 10)`

## Append Function

- The `append` universal function can be used to append elements to a slice.
- It throws an error is the return value of the append function is not assigned.
- Eg:
  - `a = append(a, 5)`
  - `a = append(a, 5, 2, 4, 6)`
  - `a = append(a, b...)`

## Similarities with Arrays

- Slices like arrays use index to fetch and manipulate value stored at given index.
- Both slice and arrays are zero-indexed.
- For both arrays and slices, accessing the elements out of the range (0 to length of array -1) leads to an out of range error (Go Code Panics).

## Empty & Nil Slice

- `nil` in Go is different from `null` in other languages.
- `nil` means a lack of value and type.
- `nil` is the zero value for slice.
- `nil` slice does not have an underlying array.
- An `empty` slice has underlying array with no values stored.
- A `nil` slice: `var a []int`
- An `empty` slice: `var a []int{}`
