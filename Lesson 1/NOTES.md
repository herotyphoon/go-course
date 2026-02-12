# Introduction

- It is a statically typed language with added features of memory safety and garbage collection.
- It is a highly efficient language, specially for multi-core processors.
- Go programs are compiled in binary thus it can be run on any device without any installation or dependencies.

## Installation

- Install go from this website [Go Install](https://go.dev/doc/install).
- Run `go version` in the terminal to verify its installation.
- For the editor, install VS Code and install the official Go extension in it.

## Getting Started

- Run `go help` to list all go commands.
- A go file has the extension `.go`.

## Creating Your First Go Program

1. Create a file with `.go` extension. This contains the go program to be executed.

2. Write the program in the file.

    ```go
    package main

    import "fmt"

    // Main program
    func main() {
        // Code to be executes
        fmt.Println("Hello, World!")
    }
    ```

    - Line 1 contains `package main` of the program, which have overall content of the program.It is the initial point to run the program, So it is compulsory to write.
    - Line 2 contains an import statement `"fmt"`. Import statement is used to include external packages that provide additional functionality beyond the built-in language features. In this case, "fmt" is a package that provides functions for formatting input and output.
    - Line 3 contains `main` function, it is beginning of execution of program.
    - Line 4 contain `fmt.Println()` is a standard library function to print something as a output on screen.In this, fmt package has transmitted Println method which is used to display the output.

3. Run the command `go run main.go`, where *main.go* is the filename.

4. If you want to build a binary out of the go code, this can be done using the `build` command. Run the command `go build -o "hello_world" main.go`, where *-o* is used to specify the name for the binary in this case *hello_world* and *main.go* is the filename. 

5. To run the binary file use `./hello_world` where *hello_world* is the name of the binary