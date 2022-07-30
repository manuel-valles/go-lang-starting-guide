# Go [Programming Language] Guide

A starting point for **Go** (the programming language) [and _Gopher_'s adventures].

## Basic Concepts

### Resources

- [Main resource](https://go.dev/learn/)
- [Download for PC](https://go.dev/dl/) and follow the wizards steps

### Workspaces & Packages

A **workspace** in **Go** is a directory where your Go files will go.

- There is actually a hierarchy of directories within your workspace where you will store the different types of Go files that you are working with. This is because common organization is good for sharing. Indeed, the following is **recommended**:

  - Three subdirectories:
    - `src` contains source code files
    - `pkg` contains packages/libraries
    - `bin` contains executables
  - Programmer typically has one workspace for many projects
  - Workspace directory is defined by `GOPATH` environment variable. This is automatically defined during the [wizard] installation, i.e. `C:\Users\{username}\go`
  - **Go tools** assume that code is in `GOPATH`

A **package** in **Go** is a group of related source code files.

- Each package cab be imported by other packages what also enables software reuse
- First line of file names the package. This name will be use for the imports
- `import` keyword is used to access other packages
- Go standard library includes many packages, i.e. "fmt"
- Searches directories are specified by `GOROOT` and `GOPATH`
- There must be one package called `main` where the execution starts:
  - Building/compiling it generates an executable program
  - It needs a `main()` function
- Example:
  ```go
    package main
    import "fmt"
    func main(){
        fmt.Printf("hello, world!\n")
    }
  ```

### Go Tools

**Go Tool** is a tool to manage Go source code. There is a bunch of different commands that you can use for the Go Tool to do. Some of them are:

- `go build` compiles the program:
  - The arguments can be a list of packages or a list of `.go` files
  - It creates an executable for the main packages with the same name as the first `.go` file
  - `.exe` suffix is qhat you will see for executables in Windows
- `go doc` prints documentation for a package. Basically you have to put the documentation in your package and this command will just pull it out of all your packages, and print it
- `go fmt` formats source code files. This will just indent in the way it should be done
- `go get` downloads packages and installs them
- `go list` lists all installed packages
- `go run` compiles `.go` files and runs the executable
- `go test` run tests using files ending in `_test.go`

### Variables

A **variable** is basically data stored in memory.

- It must have a `name` that start with a letter and they can have any number of letters, digits and underscores. It is canse sensitive and you can't use keywords, e.g. `if`, `case`,...
- It must have a `type`
- It must have `declarations` that specidy the name and the type of the variable. The most basic declaration would be: `{kyeword} {name} {type}`, e.g. `var x int`
- The compiler needs to know what type of variable it is, so it knows how much space to allocate, what operations to perform that type of,...
- It can be declared many on the same line: `var x, y int`
- A `variable type` defines the values that a variable can take and the operations that can be performed on it
- The _basic types_ types are:
  - `integer`: only integral values
  - `floating point`: fractional(decimal) values
  - `string`: byte(character) sequences

You can make `type declarations` where you define an alias, alternate name for a type.

- It is useful for clarity inside a particular application, e.g. `type Celsius float64`
- You can declare variables using the type alias, e.g. `var temp Celsius`

Every variable has to be **initialized** somehow before you use it:

- In the declaration itself:
  - `var x int = 100`
  - `var x = 100` the compiler will **infer** the type based on the type on the right hand side value
- After the declaration:

  ```go
  var x int // x=0
  x = 100

  var y string // x=""
  y = "hello"
  ```

- Perform a declaration and initialization together with the `:=` operator:
  - This is called `short variable declarations`, e.g. `x := 100`
  - The variable is declared as type of expression on the right hand side
  - It can be only done inside a function

### Basic Data Types

#### Pointers

A pointer is an address to some data in memory, so every variable and function is located somewhere.

There are two main operators that are associated with _pointers_:

- `&` returns the address of the variable or the function, whatever the name is referring to
- `*` does the opposite of the ampersand. It returns the data at the address (`dereferencing`)

```go
  var x int = 1
  var y int
  var ip *int // ip is pointer to an integer

  ip = &x // ip now points to x
  y = *ip // y is now 1
```

#### New

`new()` is another way to create a variable. It is a function that creates a variable (initialised to zero) and returns a pointer to the variable.

```go
  ptr := new(int)
  *ptr = 3 // The value 3 is placed at the address specified by the variable ptr
```

#### Variable Scope

The scope of a variable is the places in the code where a variable can be accessed. In Go, variable scoping is done using **blocks**.

A _Block_ is a sequence of declarations and statements within matching curly brackets, `{}`, including functions definitions.

There are also a hierarchy of _implicit blocks_:

- _universe block_ is all Go source
- _package block_ is all source in a package
- _file block_ is all source in a file
- all th code inside the _if_,_for_ and _switch_ statement
- individual clauses in _switch_ and _select_ get a block

Go is **lexically scoped** using blocks and the relationship of one block being defined inside another block. In other words, a variable is accessible from a block if the variable is declared in the same block or in a bigger block.

#### Deallocating Memory

A variable should be deallocated when a variable is no longer needed, so the memory space is made available. Otherwise, it will eventually run out of memory, what is called `memory leak` and pretty common in _C_.

- `Stack` is an area of memory that is dedicated to function calls. Local variables are stored here and they are deallocated automatically after function completes
- `Heap` is persistent. In compiled languages like _C_ you must deallocated the data on the heap manually. It's faster but it could also produce errors.

```c
  x = malloc(32);
  free(x);
```

#### Garbage Collection

Since it's hard to determine when a variable is no longer in use, there are some tools like **garbage collection** that is an automatic tool that deals with deallocation. This is part of interpreted languages and this is done by the interpreter such as the Java Virtual Machine in Java and the Python Interpreter in Python.

This use of this tool is easy for the programmer but also slow due to the need of an interpreter.

_Go_ is a compiled language which enables garbage collection. There are many ways of doing it but generally you have to keep track of the pointers to a particular object. Once all the pointers are gone, then you know that the object can be deallocated.

Garbage collection in Go allows to allocate stuff on the heap and the stack itself. Although there is a downside because of the act of garbage collection does take some time, it's a pretty efficient implementation. It slows things down a little bit but it is a great advantage because it makes programming a lot easier. You don't have to go as far as using a full-on interpreter like you would in an interpreted language.

#### Comments and Printing

- **Comments** are text for understandability, ignored by the compiler. They can be `single-line` (`// ...`) or `block` (`/* ... */`)
- **Printing** is done using the _format_(`fmt`) package:
  - `fmt.Printf()` or `fmt.Println()` prints a string
  - `%s` is the conversion character for a string, e.g. `fmt.Printf("Hi %s, x)`

#### Integers

Although the generic integer declaration is `int`, there are different lengths and signs: `int8`, `int16`, `int32`, `int64`, `uint8`, `uint16`, `uint32`, `uint64`. Th numbers are the number of bits that are used to represent the integer and memory, and the `u` stands for `unsigned` integers what means that the integer can get larger.

#### Floating Points

A floating point can be expressed using:

- decimals: `var x float64 = 123.45`
- scientific notation: `var y float64 = 1.2345e2`
- complex numbers represented as two floats: real and imaginary `var z complex128 = complex(2,3)` <-> `2 + 3i`

- `float32` will provide approximately 6 digits of precision
- `float64` will provide approximately 15 digits of precision

#### Type Conversion

Most binary operations need operands of the same type, including assignments:

```go
  var x int32 = 1
  var y int16 = 2
  x = y // This will fail because they are two different types of integers
```

You could use the covert type with `T()` operation: `x = int32(y)`

#### ASCII and Unicode

Each one of the characters that you want to store in a string has to be coded according to a standardized code. **American Standard Code for Information Interchange** (_ASCII_) associates each character with an 8-bit number (0-255 characters), e.g. `'A' = 0x41`

This has some limitations for many languages like Chinese. **Unicode** solves this using a 32-bit character code (2gig characters size). **UTF-8** is a subset of _Unicode_ with variable length code (from 8 to 32 bits)m where the first set of 8-bit matches _ASCII_.

The default in _Go_ is _UTF-8_, and the **code point** (called `Rune` in Go) is define by _unicode_ characters.

#### Strings

A **string** is a sequence of arbitrary bytes represented in UTF-8 often meant to be printed. Each byte is a rune represented as a UTF-8 code point. So they are read-only. You cannot modify a string but you can make a new string that is a modified version of an existing one.

- **String literal** is notated by double quotes, e.g. `x := "Hi world!"`

- **Unicode package** provides a set of functions that evaluates the properties of the different runes inside the strings.

  - Some examples that return `true` or `false`:
    - `IsDigit(r rune)`
    - `IsSpace(r rune)`
    - `IsLetter(r rune)`
    - `IsLower(r rune)`
    - `IsPunt(r rune)`
  - Some functions performs conversions:
    - `ToUpper(r rune)`
    - `ToLower(r rune)`

- **Strings package** provides a set of functions to look at the whole string and manipulate UTF-8 encoded strings

  - String search functions:
    - `Compare(a, b)` returns an integer comparing two string lexicographically: `0` if `a==b`, `-1` if `a<b` and `+1` if `a>b`
    - `Contains(s, substr)` returns `true` if substring is inside `s`
    - `HasPrefix(s, prefix)` returns `true` if the string `s` begins with `prefix`
    - `Index(s, substr)` returns the index of the first instance of `substr` in `s`

- **String Manipulation** provides a set of functions to return modified strings (original are immutable):

  - `Replace(s, old, new, n)` returns a copy of the string `s` with the first `n` instances of `old` replaced by `new`
  - `ToLower(s)`
  - `ToUpper(s)`
  - `TrimSpace(s)` returns a new string with all leading and trailing white space removed

- **Strconv Package** provides a set of functions for conversions to and from string representations of basic data types
  - `Atoi(s)` converts a string `s` to `int`
  - `Itoa(s)` converts `int(base 10)` to `string`
  - `FormatFloat(f, fmt, prec, bitSize)` converts floating point number to a string
  - `ParseFloat(s, bitSize)` converts a string to a floating point number

#### Constants

A **constant** is an expression whose value is known at compile time, and whose type is inferred from right-hand side (`boolean`, `string`, `number`)

```go
  const x = 1.3
  const (
    y = 4
    z = "Hi"
  )
```

**iota** generates a set of related but distinct constants (must be different but the actual value is not important). It often represents a property which has several distinct possible values, e.g. days of the week. It is like an enumerated type in other languages.

```go
  type Grades int
  const (
    A Grades = iota // Each constant is assigned to a unique integer
    B // The implementation starts at 1 and increments but the actual value is not important
    C
    D
    F
  )
```

#### Control Flow

**Control flow** describes the order in which statements are executed inside a program. The most basic control flow is executing one statement at a time, one after the other (top-down).

Control flow changes for a lot of reasons but the first one is because the programmer inserts control flow structures into their code, which changes the sequence in which the statements are executed:

- _**if** statement_ is the main control flow structure, where the expression `condition` is evaluated, and the `consequent` statements are executed if the condition is `true`

```go
  if x > 5 {
    fmt.Printf("Yep")
  }
```

- _**for** loops_ iterates while a condition is `true`, and may have an initialization and update operation:

```go
  for i:=0; i<10; i++ {
    fmt.Printf("hi")
  }

  j=0
  for j<10 {
    fmt.Printf("hello")
    j++
  }
```

- **_break_** exits the containing loop, and **_continue_** skips the rest of the current iteration:

```go
  i := 0
  for i<10 {
    i++
    if i == 3 { continue }
    if i == 5 { break }
    fmt.Printf("bye")
  }
```

- **_switch_** is a multi-way if statement that may contain a `tag` which is a variable to be compared to a constant defined in each `case`. The case that matches is executed. _NOTE_: You do not need a `break` in each case, it automatically breaks.

```go
  switch x {
    case 1:
      fmt.Printf("case 1")
    case 2:
      fmt.Printf("case 2")
    default:
      fmt.Printf("no case")
  }
```

A **tagless switch** is a switch which case contains a boolean expression to evaluate. First `true` case is executed:

```go
  switch {
    case x>1:
      fmt.Printf("case 1")
    case x<-1
      fmt.Printf("case 2")
    default:
      fmt.Printf("no case)
  }
```

#### Scan

The **Scan** reads the user input. It takes a pointer as an argument, waits for the user input, writes it to the pointer and returns the number of scanned items followed by a `null` or `error`:

```go
  var appleCount int

  ftm.Printf("Number of apples?")
  num, err := fmt.Scan(&appleCount) // the code will stop until the user input something and hits enter

  fmt.Printf(appleCount)
```
