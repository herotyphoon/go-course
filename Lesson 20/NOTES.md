# Conditional Statements

## The 'if' Statement

- The `if` statement is employed to conduct conditional checks based on boolean expressions, which are commonly referred to as conditionals.
- The code block within the `if` statement will be executed only when the boolean expression or condition evaluated to true.

## Example

```go
if age >= 18 {
  fmt.Println("You are an adult!")
}
```

## The 'if-else' Statement

- The `if-else` statement servers as an extension of the `if` statement.
- It is employed when there are precisely two conditions that need to be evaluated.
- This construct allows for distinct code blocks to be executed depending on whether the condition is `true` or `false`.

## Example

```go
if age >= 18 {
  fmt.Println("You are an adult!")
} else {
  fmt.Println("You are not an adult!")
}
```

## The 'if-else-if' Statement

- The `if-else-if` is an extension of `if-else` statement.
- This statement is employed when there are more than two conditions you want to evaluate.
- You can create a chain of `if-else-if` statments by adding blocks of `else-if` statements as needed.

## Example

```go
if age >= 18 {
  fmt.Println("You are an adult!")
} else if age >= 13 {
  fmt.Println("You are a teenager!")
} else {
  fmt.Println("You are a child!")
}
```

## The 'if' Scoped Variables

- In Go, you have the capability to initialize a new variable directly within the `if` statement.
- The variable declared in this manner is scoped exclusively to `if` or (`if-else`) block.
- The feature aids in writing more consise and organized code.

## Example

```go
if even := isEven(age); even {
  fmt.Println("Age is even")
}
```
