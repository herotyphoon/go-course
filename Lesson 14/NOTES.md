# Arrays

- Arrays are data structures used to store data in contiguous blocks of memory.
- Each piece of data of an array is called and element.
- Each element of an array need to be of the same type.
- We use the index of the element in the array to access the element.
- Arrays are zero indexed, i.e. their index starts with zero.
- Accessing the elements out of the range (0 to length of array -1) leads to an out of range error (Go Code Panics).

## Benefits of an Array

- Arrays are very efficient to fetch elements from memory.
- Used to store a collection of data, such as a list of numbers etc.
- Uses include search and sort algorithms, buffers etc.

## Declaring Arrays in Go

- Declaring an array
  - `var a [5]int`
  - `var a = [5]int{1,2,3,4,5}`
  - `var a = [5]int{5, 2:10,50}`
  - `var a = [...]int{1,4,6}`
- Use the `len` function to get the length of an array.

## Limitations

- Once declared the size cannot be altered.
- The size and length of an array are stored in the type of an array.
- A variable cannot be used to declare the size of the array.
- Most cases slices are favoured over array.
