# A Library for Others - README

## Project Overview

**A Library for Others** is an individual project aimed at developing a CSV parsing library in Go. The project focuses on mastering file handling, error management, interfaces, and resource management. This library will be versatile, user-friendly, and robust enough to handle various CSV formats and potential parsing challenges.

---

## Learning Objectives

1. **File Handling in Go**: Learn to handle file input/output efficiently using Go's standard library.
2. **CSV Parsing**: Develop skills to parse structured CSV data with variations in formatting.
3. **Interfaces**: Understand and implement Go interfaces to design flexible and maintainable code.
4. **Error Handling**: Handle and report errors effectively, ensuring reliable application behavior.

---

## Features to Implement

### CSVParser Interface

```go
type CSVParser interface {
    ReadLine(r io.Reader) (string, error)
    GetField(n int) (string, error)
    GetNumberOfFields() int
}
```

### Interface Method Details

#### 1. `ReadLine(r io.Reader) (string, error)`
- Reads a single line from the CSV file.
- Removes the line terminator (`\n`, `\r`, `\r\n`).
- Handles quoted fields and returns an `ErrQuote` error if quotes are mismatched.
- Efficiently reads large files line-by-line without loading the entire file into memory.

#### 2. `GetField(n int) (string, error)`
- Retrieves the **nth** field from the last line read by `ReadLine`.
- Removes enclosing quotes from fields.
- Returns an `ErrFieldCount` error for invalid field indices.

#### 3. `GetNumberOfFields() int`
- Returns the number of fields in the last line read by `ReadLine`.
- Ensures correct behavior when called after EOF or before any line is read.

---

## Error Handling

The library defines specific errors:
- **`ErrQuote`**: For missing or mismatched quotes in a quoted field.
- **`ErrFieldCount`**: For invalid field access (e.g., out-of-bounds or negative index).

---

## Usage Example

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("example.csv")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    var csvparser CSVParser = YourCSVParser{}

    for {
        line, err := csvparser.ReadLine(file)
        if err != nil {
            if err == io.EOF {
                break
            }
            fmt.Println("Error reading line:", err)
            return
        }

        fields := csvparser.GetNumberOfFields()
        fmt.Printf("Read line: %s with %d fields\n", line, fields)

        for i := 0; i < fields; i++ {
            field, err := csvparser.GetField(i)
            if err != nil {
                fmt.Printf("Error retrieving field %d: %v\n", i, err)
            } else {
                fmt.Printf("Field %d: %s\n", i, field)
            }
        }
    }
}
```

---

## Development Requirements

- **Coding Standards**: Code must conform to `gofumpt` formatting.
- **Error Handling**: No panics (e.g., nil pointer dereferences, index out of range).
- **File Handling**: Use only `io.Reader` for file input.
- **Testing**: Test the library rigorously with various input files (e.g., valid CSVs, malformed CSVs, binary files).

---

## Notes for Developers

- Focus on **information hiding** by exposing only necessary methods via the `CSVParser` interface.
- Efficiently handle large files by processing them line-by-line.
- Ensure that the library is resilient to edge cases, such as empty fields, inconsistent line endings, and malformed input.

---

## Deadline and Grading Criteria

- **Deadline**: 10.12, 18:56
- **Maximum XP**: 2500 XP
- **Grading Criteria**:
  - Code quality and adherence to standards.
  - Correct implementation of the `CSVParser` interface.
  - Robust error handling.
  - Comprehensive testing.

---

## Testing and Debugging Tips

1. **Test Input Files**:
   - Valid CSV files.
   - CSV files with mismatched quotes.
   - Binary files or invalid formats.

2. **Edge Cases**:
   - Empty lines and fields.
   - Very large files with millions of entries.
   - Lines with various line endings (`\n`, `\r`, `\r\n`).

3. **Memory Management**:
   - Ensure no memory leaks by managing resources correctly.
   - Avoid loading the entire file into memory.

4. **Error Scenarios**:
   - Missing or extra quotes.
   - Invalid field indices.
   - EOF behavior for `GetField` and `GetNumberOfFields`.

---
