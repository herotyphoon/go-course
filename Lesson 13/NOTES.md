# Binary Shift Operators

- They are use to remove or append bits in integer values
- Two types:
  - > > right shift
  - << left shift

## Right Shift Operator

- Removes bits from integers and shifts them to right
- `a >> i` removes i bits from a and shifts i bits from the number to right
- Right shifting i bits of the number a reuslts in a number that is equal to a divided by 2^i (^ meaning raise to the power)

## Left Shift Operator

- Removes bits from integers and shifts them to left
- `a >> i` removes i bits from a and shifts i bits from the number to left
- Right shifting i bits of the number a reuslts in a number that is equal to a multiplied by 2^i (^ meaning raise to the power)

## Interesting Facts

- `1 << i` will result in the a single bit being set in the i™ position while rest of the bits are zero.
- The operation `a | 1<<i` will set the i bit of a to 1.
- The operation `a &^ 1<<i` will set the ith bitof a to 0 (bit clear).
- The operation a \* `1<<i` will inverts the ith bit of a.

NOTE: All bitwise operators should be used with care for signed integers
