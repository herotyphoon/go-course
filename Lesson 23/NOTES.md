# Control Flow

## Break Statement

- The `break` statement exits the current loop prematurely.
- When a specified condition is met, it allows you to leave the loop and continue with subsequent code.
- It can be utilized with a for loop to mimic the behaviour of a do-while loop.

### Syntax

```go
for {
  // code block
  condition {
    break
  }
}
```

## Continue Statement

- The `continue` statement is employed to bypass a specific iteration of a for loop.
- This is useful when you wish to skip executing the loop's code block under specific conditions.

### Syntax

```go
for {
  // code block
  condition {
    continue
  }
}
```

## Label

- The `continue` statement by itself only affects the loop in which it is placed.
- When dealing with multiple loops, `labels` allows us to countinue to an outer loop.
- In Go, usage of `labels` is infrequent and considered rare.
- Labels can be used with `break`, `continue and `goto` statements.

### Syntax

```go
outer:
  for for-clause {
    for for-clause {
      if condition {
        continue outer
      }
      // code block
    }
  }
}
```

## Goto

- The 'goto' statement is employed to redirect the flow of code.
- In Go, there are limitations on the use of 'goto':
  - It is not allowed to jump over variable declarations.
  - It is not allowed to jump into inner or same-level code blocks.

### Syntax

```go
  if condition {
    goto end
  }
  // skip code
end:
  // code
```