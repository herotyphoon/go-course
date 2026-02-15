# Variables

## Introduction

- Variables are declared to store values that can change throughout.
- There can be package level variables (not recommended) or function level variables.¸
- Variable type declaration can be explicit or implicit.
- Variables declaration can be grouped with braces ().
- Multiple variables can be declared in single line.

## Type Assignment

- In explicit type assignment, the type is assigned in declaration.
- Eg: `var age int` or `var age int = 10`
- In implicit type assignment, the type is inferred from the right side value.
- The default type of value on right is used to assign type to the variable.
- Eg: `var age = 10`

## Short Variable Declaration

- The keyword `var` can be omitted along with the `type` which is inferred from the value on right in variable declaration.
- Can be used to declare multiple variables as long the variable on the left is new.
- Eg: `age := 10`

### Examples of variabl declaration

```go
var age int
var age int = 10
var age = 10
var age, name = 10, "Code"
age := 25
age, name := 10, "Code"
```
