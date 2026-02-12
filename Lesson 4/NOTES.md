# Floats & Complex Numbers

## Data Type `float`

- Floating point numbers have *integer* value, a *decimal point* and a *fraction*.
- Unlike our `int` data types, a `float` can store fractions.
- Go like many modern languages uses IEEE 754 standard to implement floats.
- Floats have limited level of precision to certain extent and should be used carefully.
- When comparing floats, they should not be compared directly but instead the check should be around if the difference is between a given range.

| Type    | Description                                           | Range                                                     |
|---------|-------------------------------------------------------|-----------------------------------------------------------|
| `float32` | the set of all IEEE-754 32-bit floating-point numbers | 1.401298464324817070923729583289916131280e-45 to 3.40282346638528859811704183484516925440e+38 |
| `float64` | the set of all IEEE-754 64-bit floating-point numbers | 4.940656458412465441765687928682213723651e-324 to 1.797693134862315708145274237317043567981e+308 |
 
## Data Type `complex`

- Complex numbers in Go have the real and imaginary part represented by `float`.
- Complex64 uses `float32` for real and imaginary part each, while `complex128` uses `float64` as the real and imaginary part each.
- Built-in function `complex()` can be used to create a complex number.
- Built-in function `real()` and `imag()` can be used to get the real and imaginary part.

| Type       | Description                                                          |
|------------|----------------------------------------------------------------------|
| `complex64`  | the set of all complex numbers with float32 real and imaginary parts |
| `complex128` | the set of all complex numbers with float64 real and imaginary parts |