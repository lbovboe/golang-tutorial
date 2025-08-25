# Go Type Conversions

This guide covers the essential type conversion functions in Go, helping you convert between different data types safely and effectively.

## Table of Contents

1. [Numeric Conversions](#numeric-conversions)
2. [String Conversions](#string-conversions)
3. [Boolean Conversions](#boolean-conversions)
4. [Best Practices](#best-practices)

## Numeric Conversions

Go allows explicit type conversions between numeric types using type casting syntax.

### Basic Numeric Type Conversions

```go
// Converting between integer types
var i int = 42
var i8 int8 = int8(i)
var i32 int32 = int32(i)
var i64 int64 = int64(i)

// Converting between floating-point types
var f32 float32 = 3.14
var f64 float64 = float64(f32)

// Converting between integers and floats
var intVal int = 42
var floatVal float64 = float64(intVal)
var backToInt int = int(floatVal)

// Converting to unsigned integers
var signed int = -42
var unsigned uint = uint(signed) // Be careful with negative numbers!
```

### Important Notes:

- **Data Loss**: Converting from larger to smaller types may cause data loss
- **Overflow**: Converting negative numbers to unsigned types can cause unexpected results
- **Precision Loss**: Converting from float to int truncates the decimal part

## String Conversions

The `strconv` package provides functions to convert between strings and other basic data types.

### String to Number Conversions

#### `strconv.Atoi(s string) (int, error)`

Converts a string to an integer.

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "123"
    num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("String '%s' converted to integer: %d\n", str, num)
}
```

#### `strconv.ParseFloat(s string, bitSize int) (float64, error)`

Converts a string to a floating-point number.

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "3.14159"

    // Parse as float64
    f64, err := strconv.ParseFloat(str, 64)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    // Parse as float32 (but returns float64)
    f32, err := strconv.ParseFloat(str, 32)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    fmt.Printf("String '%s' as float64: %f\n", str, f64)
    fmt.Printf("String '%s' as float32: %f\n", str, float32(f32))
}
```

### Number to String Conversions

#### `strconv.Itoa(i int) string`

Converts an integer to a string.

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    num := 42
    str := strconv.Itoa(num)
    fmt.Printf("Integer %d converted to string: '%s'\n", num, str)
}
```

#### `strconv.FormatFloat(f float64, fmt byte, prec, bitSize int) string`

Converts a floating-point number to a string with specified formatting.

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    f := 3.14159265359

    // Different formatting options
    str1 := strconv.FormatFloat(f, 'f', 2, 64)  // Fixed decimal: 3.14
    str2 := strconv.FormatFloat(f, 'e', 2, 64)  // Scientific: 3.14e+00
    str3 := strconv.FormatFloat(f, 'g', -1, 64) // Automatic: 3.14159265359

    fmt.Printf("Original float: %f\n", f)
    fmt.Printf("Fixed format (2 decimals): %s\n", str1)
    fmt.Printf("Scientific format: %s\n", str2)
    fmt.Printf("Automatic format: %s\n", str3)
}
```

**Format options for `strconv.FormatFloat`:**

- `'f'`: Fixed decimal point (e.g., 123.456)
- `'e'`: Scientific notation (e.g., 1.23e+02)
- `'E'`: Scientific notation with capital E (e.g., 1.23E+02)
- `'g'`: Automatic format (chooses between 'f' and 'e')
- `'G'`: Automatic format with capital E

## Boolean Conversions

### `strconv.ParseBool(str string) (bool, error)`

Converts a string to a boolean value.

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Valid boolean strings
    validBools := []string{"true", "false", "TRUE", "FALSE", "True", "False", "1", "0", "t", "f", "T", "F"}

    for _, str := range validBools {
        b, err := strconv.ParseBool(str)
        if err != nil {
            fmt.Printf("Error parsing '%s': %v\n", str, err)
        } else {
            fmt.Printf("String '%s' -> Boolean: %t\n", str, b)
        }
    }

    // Invalid boolean string
    invalid := "maybe"
    _, err := strconv.ParseBool(invalid)
    if err != nil {
        fmt.Printf("Error parsing '%s': %v\n", invalid, err)
    }
}
```

**Accepted values for `strconv.ParseBool`:**

- `true`: "true", "TRUE", "True", "1", "t", "T"
- `false`: "false", "FALSE", "False", "0", "f", "F"

### `strconv.FormatBool(b bool) string`

Converts a boolean to a string.

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    trueBool := true
    falseBool := false

    trueStr := strconv.FormatBool(trueBool)
    falseStr := strconv.FormatBool(falseBool)

    fmt.Printf("Boolean %t -> String: '%s'\n", trueBool, trueStr)
    fmt.Printf("Boolean %t -> String: '%s'\n", falseBool, falseStr)
}
```

## Best Practices

### 1. Always Handle Errors

```go
// Good: Handle conversion errors
str := "not-a-number"
num, err := strconv.Atoi(str)
if err != nil {
    fmt.Printf("Conversion failed: %v\n", err)
    return
}

// Bad: Ignoring errors
num, _ := strconv.Atoi(str) // Don't do this in production code
```

### 2. Use Type Assertions Carefully

```go
// When converting between numeric types, be aware of potential data loss
var bigInt int64 = 9223372036854775807
var smallInt int8 = int8(bigInt) // Data loss!
fmt.Printf("Original: %d, Converted: %d\n", bigInt, smallInt)
```

### 3. Validate Input Before Conversion

```go
func safeStringToInt(s string) (int, error) {
    // Trim whitespace
    s = strings.TrimSpace(s)

    // Check if string is empty
    if s == "" {
        return 0, fmt.Errorf("empty string")
    }

    // Perform conversion
    return strconv.Atoi(s)
}
```

### 4. Use Specific Parsing Functions for Better Control

```go
// Instead of strconv.Atoi for specific int types
i64, err := strconv.ParseInt("123", 10, 64)  // base 10, 64-bit
i32 := int32(i64)

// For unsigned integers
u64, err := strconv.ParseUint("123", 10, 64)
```

## Common Conversion Examples

### Example 1: User Input Processing

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter your age: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    age, err := strconv.Atoi(input)
    if err != nil {
        fmt.Printf("Invalid age: %v\n", err)
        return
    }

    if age >= 18 {
        fmt.Println("You are an adult.")
    } else {
        fmt.Println("You are a minor.")
    }
}
```

### Example 2: Configuration File Processing

```go
package main

import (
    "fmt"
    "strconv"
)

type Config struct {
    Port        int
    Debug       bool
    Timeout     float64
    ServiceName string
}

func parseConfig(configMap map[string]string) (*Config, error) {
    config := &Config{}

    // Parse port
    if portStr, exists := configMap["port"]; exists {
        port, err := strconv.Atoi(portStr)
        if err != nil {
            return nil, fmt.Errorf("invalid port: %v", err)
        }
        config.Port = port
    }

    // Parse debug flag
    if debugStr, exists := configMap["debug"]; exists {
        debug, err := strconv.ParseBool(debugStr)
        if err != nil {
            return nil, fmt.Errorf("invalid debug flag: %v", err)
        }
        config.Debug = debug
    }

    // Parse timeout
    if timeoutStr, exists := configMap["timeout"]; exists {
        timeout, err := strconv.ParseFloat(timeoutStr, 64)
        if err != nil {
            return nil, fmt.Errorf("invalid timeout: %v", err)
        }
        config.Timeout = timeout
    }

    // Service name (no conversion needed)
    config.ServiceName = configMap["service_name"]

    return config, nil
}

func main() {
    configData := map[string]string{
        "port":         "8080",
        "debug":        "true",
        "timeout":      "30.5",
        "service_name": "my-service",
    }

    config, err := parseConfig(configData)
    if err != nil {
        fmt.Printf("Configuration error: %v\n", err)
        return
    }

    fmt.Printf("Config: %+v\n", config)
}
```

This guide covers the essential type conversion functions in Go. Remember to always handle errors when using the `strconv` package functions, as they can fail if the input string is not in the expected format.
