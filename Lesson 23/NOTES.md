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