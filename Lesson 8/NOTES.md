# Identifiers and Keywords

## Identifiers

- Name given to identify entities like variables, constants, types, functions etc.
- Formed by combining one or more Unicode letters and digits.
- The first character of an identifier must be a letter.
- Underscore is considered a letter.
- Eg: `myInt`, `_myString`, `MyName`, `αβ`.

## Predeclared Identifiers

- Used to identify types, const, function etc.
- Should not be used for user defined identifiers
- These are:

  > - Types:
  >   - any bool byte comparable complex64 complex128 error float32 float64
  >     int int8 int16 int32 int64 rune string uint uint8 uint16 uint32 uint64 uintptr
  > - Constants:
  >   - true false iota
  > - Zero Value:
  >   - nil
  > - Functions:
  >   - append cap cloase complex copy delete imag len make new panic print println real recover

## Keywords

- Have a specific meaning and function.
- Cannot be used as identifiers.
- List of keywords:

  > break default func interface select
  >
  > case defer go map struct
  >
  > chan else goto package switch
  >
  > const fallthrough if range type
