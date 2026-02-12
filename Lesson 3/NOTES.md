# Integers

## Data Type `int`

- The `int` data type is used to store positive and negative integers (without fractions).
- Eg: 1, 100, -987, -78432 etc.
- If a fraction is stored in the `int` data types the fraction part is dropped.
- There are different kind of `int` data types such as `int8`, `int16`, `int32`, etc.
- The different `int` data types have different size in memory and as such have a different range of numbers that it can store.
- The first bit in `int` data type is used to store the sign of the value and the rest of the bits are used to store the value itself.

| Туре  | Description                           | Range                                       |
|-------|---------------------------------------|---------------------------------------------|
| int8  | the set of all signed 8-bit integers  | -128 to 127                                 |
| int16 | the set of all signed 16-bit integers | -32768 to 32767                             |
| int32 | the set of all signed 32-bit integers | -2147483648 to 2147483647                   |
| int64 | the set of all signed 64-bit integers | -9223372036854775808 to 9223372036854775807 |

## Data Type `uint`

- The `uint` data type is used to store positive integers only (without fractions).
- Eg: 1, 2349087, 12, 23423545345.
- There are different kind of int data types such as `uint8`, `unt16`, `uint32`, etc.
- The different uint data types are have different size in memory and as such have a different range of numbers that can be store.
- All the bits are used to store the value as such uint has greater range.

| Туре     | Description                               | Range                   |
|----------|-------------------------------------------|-------------------------|
| `uint8`  | the set of all unsigned 8-bit integers  | 0 to 255                  |
| `uint16` | the set of all unsigned 16-bit integers | 0 to 65535                |
| `uint32` | the set of all unsigned 32-bit integers | 0 to 4294967295           |
| `uint64` | the set of all unsigned 64-bit integers | 0 to 18446744073709551615 |

## Other Related Data Types

| Type   | Description                                               |
|--------|-----------------------------------------------------------|
|`byte`  | alias for `uint8`                                         |
|`rune`  | alias for `int32`                                         |
|`uint`  | either 32 or 64 bits dependent on the system architecture |
|`int`   | either 32 or 64 bits dependent on the system architecture |

**Note**: Explicitly specify the size of `uint` or `int` as needed or it could lead to memory leaks.