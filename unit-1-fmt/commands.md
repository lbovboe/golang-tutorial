# Go Commands and Setup Guide

## Project Initialization

**Initialize Go Module**

```bash
go mod init hello-go
```

- **Purpose**: Creates a `go.mod` file that serves as the foundation for dependency management in your Go project
- **What it does**: Establishes your project as a Go module with the name "hello-go"
- **Benefits**: Enables automatic dependency tracking, version management, and ensures reproducible builds
- **Result**: Generates a `go.mod` file containing the module name and Go version

### Understanding the Module Name in go.mod

**The go.mod File Content:**

```go
module hello-go

go 1.24.5
```

**Module Name Explanation:**

- **What it is**: `hello-go` is the module name that uniquely identifies your Go project
- **Purpose**: Acts as the root import path for all packages within your module
- **Usage**: When other projects want to import your code, they use this module name
- **Local Development**: For local projects, you can use simple names like `hello-go`
- **Public Projects**: For projects meant to be shared, use domain-based names like `github.com/username/project-name`

**When the Module Name is Used:**

- **Internal Imports**: When importing local packages within your project
- **External References**: When other projects import your module as a dependency
- **Go Commands**: Tools like `go get`, `go mod download` use this name
- **Package Resolution**: Go uses this to resolve import paths

**How it Works:**

```go
// If you have a package structure like:
// hello-go/
//   ├── main.go
//   └── utils/
//       └── helper.go

// In main.go, you would import the utils package as:
import "hello-go/utils"
```

**Best Practices:**

- **Local Projects**: Use simple, descriptive names (e.g., `calculator`, `todo-app`)
- **Shared Projects**: Use full domain paths (e.g., `github.com/yourusername/project-name`)
- **Consistency**: Keep the name consistent with your project/repository name

## File Structure

**Create Main Program File**

```bash
touch main.go
# or simply create the file in your editor
```

- **Purpose**: `main.go` serves as the entry point and main program file for your Go application
- **Convention**: While you can name it anything, `main.go` is the standard convention for the main file
- **Function**: Contains the `main()` function where program execution begins
- **Structure**: Typically includes package declaration, imports, and the main function

## Code Structure

**Package Declaration**

```go
package main
```

- **Purpose**: The `main` package is special in Go - it tells the compiler this is an executable program
- **Significance**: Only the `main` package can contain a `main()` function that serves as the program's entry point
- **Requirement**: Every Go file must start with a package declaration
- **Difference**: Other packages (like libraries) use different names and cannot be executed directly

**Import Statement**

```go
import "fmt"
```

- **Purpose**: The `fmt` package provides formatted I/O functions similar to C's printf and scanf
- **Common Functions**:
  - `fmt.Println()` - prints text with a newline
  - `fmt.Printf()` - prints formatted text
  - `fmt.Sprintf()` - returns formatted string without printing
- **Standard Library**: Part of Go's built-in standard library, no external installation needed
- **Usage**: Essential for basic input/output operations in Go programs

## Running and Building

**Run the Program**

```bash
go run main.go
```

- **Purpose**: Compiles and executes the Go program in one step, without creating a permanent executable
- **Process**: Temporarily compiles the code to machine code and runs it immediately
- **Benefits**: Quick testing and development - no need to manually build and then run
- **Use Case**: Ideal for development and testing phases when you want instant feedback

**Build Executable**

```bash
go build
```

- **Purpose**: Compiles your Go source code into a standalone executable binary file
- **Output**: Creates an executable file named after your module (e.g., `hello-go` or `hello-go.exe` on Windows)
- **Benefits**:
  - Single file deployment - no dependencies needed on target machine
  - Fast execution - pre-compiled machine code
  - Cross-platform compilation support
- **Use Case**: Production deployment and distribution of your Go applications
