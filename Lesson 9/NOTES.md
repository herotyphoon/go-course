# Operators

- Operators are symbols that perform arithemetic and logical operations.
- Operation is performed on operands.
- Binary operators perform operation on two operands while unary operators perform operations on one operand.

## Unary Operators

| Operator | Operation       | Description                        |
| -------- | --------------- | ---------------------------------- |
| ++       | equal           | ++a incements the value of a by 1  |
| --       | not equal       | --a decrements the value of a by 1 |
| &        | less than       | &a gives the address of a          |
| \*       | less than equal | \*a get the value at address of a  |
| <-       | greater than    | <-a recieves value from channel a  |

## Arithemetic Operators

| Operator | Operation  | Types                                           |
| -------- | ---------- | ----------------------------------------------- |
| +        | sum        | a + b Integers, Floats, Complex Values, Strings |
| -        | difference | a - b Integers, Floats, Complex Values          |
| \*       | product    | a \* b Integers, Floats, Complex Values         |
| /        | division   | a / b Integers, Floats, Complex Values          |
| %        | remainder  | a % b Integers                                  |

## Bitwise Operators

| Operator | Operation             | Types           |
| -------- | --------------------- | --------------- |
| &        | Bitwise AND           | a & b Integers  |
| \|       | Bitwise OR            | a \| b Integers |
| ^        | Bitwise XOR           | a ^ b Integers  |
| &^       | Bit Clear ( AND NOT ) | a &^ b Integers |
| <<       | Left Shift            | a << b Integers |
| >>       | Right Shift           | a >> b Integers |

## Comparison Operators

| Operator | Operation          | Description                                                     |
| -------- | ------------------ | --------------------------------------------------------------- |
| ==       | equal              | a == b true if a is equal to b, false otherwise                 |
| !=       | not equal          | a != b true if a is not equal to b, false otherwise             |
| <        | less than          | a < b true if a is less than b, false otherwise                 |
| <=       | less than equal    | a <= b true if a is less than or equal to b, false otherwise    |
| >        | greater than       | a > b true if a is greater than b, false otherwise              |
| >=       | greater than equal | a >= b true if a is greater than or equal to b, false otherwise |

## Logical Operators

| Operator | Operation       | Description                                                      |
| -------- | --------------- | ---------------------------------------------------------------- |
| &&       | conditional and | a && b is true if a and b are both true, false otherwise         |
| \|\|     | conditional OR  | a \|\| b is true if either a or b are both true, false otherwise |
| !        | not             | !a is not a                                                      |
