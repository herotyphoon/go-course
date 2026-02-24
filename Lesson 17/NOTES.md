# Strings, Runes & Bytes

## Strings

- Strings are stored as a sequence of bytes.
- You can access each byte of the string data using the index notation.
- The slice expression works with strings.
- Strings in Go are immutable, you cannot manipulate a string, once created.
- The string, rune and byte types can be converted to each other.

## Examples

- Given `s` is a string `"Code & Learn"`:

|     | Example         | Output                      |
| --- | --------------- | --------------------------- |
| 1.  | `string(s[0])`  | `"C"`                       |
| 2.  | `string(s[:5])` | `"Code"`                    |
| 3.  | `string(s[3:])` | `"e & Learn"`               |
| 4.  | `string(s[:])`  | `"Code & Learn"`            |
| 5.  | `s[1] = 'a'`    | `'UnsignableOperand' error` |

## UTF-8 & Unicode

- The byte data of string is assumed to be UTF-8 encoded code point.
- UTF-8 encoding and Unicode support means, any character of string in Go can take 1 to 4 bytes.
- If a code point takes more than 1 byte, fetching byte data directly and converting to a string will lead to unexpected results.

## Examples

- Example, we have string `s := "コーディング"` which includes unicode characters with code points of size more than 1 byte.
- Access a single byte `(string(s[1]))` will give you invalid output.
- Convert to rune slice, use strings or unicode/utf8 package or use for-range.
- `s := "コーディング"` convert to rune slice `r := []rune(s)` and then access each character like `string(r[2])`.
