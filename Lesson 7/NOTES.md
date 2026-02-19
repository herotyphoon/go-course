# Constants

## Introduction

- Values that never change during the execution of the program or are immutable.
- Declared with the `const` keyword.
- In Go constants can only hold values that the compiler can infer are constants at the compile time.
- You cannot make a runtime value constant.

## Constant Values

- Boolean values i.e `true` or `false`.
- Numeric values such as `1`, `100`, `1.01`, etc.
- Global functions like `complex`, `real`, `imag`, `len` and `cap`.
- A string or a rune value, e.g. `"Typhoon"`, `'a'`
- Expression that can be evaluated at compile time, e.g. `1<2`.

## Constant Types

- There are two types of constants, _typed_ and _untyped_ constants.
- A typed constant can only be assigned directly to a variable of same kind.
- For untyped constant type can be inferred or it fallbacks to default type for the constant type.

## Variables & Constants

- Similarities
  - Package-level and function-level
  - Can be grouped using braces `()`
- Differences
  - Constant values cannot be change but variables can.
  - Constants can remain un-used.
