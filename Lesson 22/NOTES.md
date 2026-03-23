# Multifunctional For Loop

- Loops in programming serve to repetitively execute a block of code.
- Throughout history, programming languages have featured constructs such as the for loop, while loop, and do-while loop.
- In Go, the singular type of loop available is the `for` loop.

## Types of Loops

- Three kinds of loops:
  1. For with a condition
  2. For with a clause
  3. For with a range

### For loop with a Condition

- The 'for with condition' loop type requires a single condition.
- It will continue executing as long as the condition remains true.
- If the condition is omitted, the loop will default to true, resulting in an infinite loop.
- To exit an infinite loop, we can use the `break` statement.

#### Syntax

```go
for condition {
    // code block
}
```

### For loop with a Clause

- This is the most standard form of a for loop, present in nearly all programming languages that support this construct.
- It comprises three components separated by semicolons:
  - Initialization.
  - Condition.
  - Post Statement.
- Any of these parts can be left empty.

#### Syntax

```go
for init; condition; post {
    // code block
}
```

### For loop with a Range

- A 'for loop with a range clause' iterates through elements of an array, slice, string, map, channel, or any derived type from these composite structures.
- It yields two iteration variables for arrays, slices, strings, and maps, but only a single iteration variable for channels.

#### Syntax

```go
for val1, val2 := range x {
    // code block
}
```

- When used with an array or slice, range provides the index of the element as the first iteration variable, and the value as the second.
- When applied to a string, range iterates over the Unicode code points within the string.
- The first value represents the index of the starting byte, and the second value corresponds to the rune value for the Unicode character.
- In a map, the first value represents the key, while the second value corresponds to the associated value for that key.
- However, it's important to note that the order of iteration is not guaranteed.
- If the map is modified during iterations, the presence of newly added or deleted values may or may not be observed.
