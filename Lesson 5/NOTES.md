# Strings

## Introduction to Strings

- Data type string is used for text in Go.
- You can define a string only with “” or \`\`. Example: "Hello World" or \`Hello World\`.
- Zero value for data type string is any empty string.
- Go supports Unicode which means it supports tons of characters, symbols and emojis and UTF-8 encoding is used for encode Unicode.

## Behind the Scenes

- Strings are stored as a byte array internally and are immutable.
- Variable holding the string can be updated but a string in Go cannot be modified directly.
- UTF-8 encoding means each Unicode character can take from 1 to 4 bytes memory.
- The rune data type is the best way to represent a character of string.

## Types of Strings

- A one line string can be created using “”. To add line breaks you would need to add "\n" escape character for new line.
- Multi-line strings can be represented with backticks `` also known as raw strings.
- Strings can be formatted with string formatters.