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
