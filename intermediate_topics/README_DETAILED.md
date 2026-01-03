# Intermediate Go Topics - Complete Guide

## Overview
This directory contains comprehensive, well-commented Go examples for all intermediate topics in the course. Each file focuses on one specific topic and provides multiple detailed examples with explanations.

## Files Created

### Core Language Concepts

**57_closures_detailed.go** - Closures
- Basic closure concepts (capturing variables)
- Closure state persistence
- Loop variable capture patterns
- Factory pattern with closures
- HTTP handler patterns
- Middleware pattern implementation
- Practical examples (counters, rate limiters, validators)

**58_recursion_detailed.go** - Recursion
- Basic recursion with base cases
- Mathematical recursion (factorial, fibonacci)
- Inefficient vs optimized recursion
- Memoization for performance
- Binary search with recursion
- Tree traversal (pre-order, in-order, post-order)
- Backtracking and N-Queens problem
- Recursion depth visualization
- Tail recursion patterns

**59_pointers_detailed.go** - Pointers
- Basic pointer concepts (& and * operators)
- Memory layout and addresses
- Pass by value vs pass by reference
- Pointers to structs
- Pointer to pointer (double pointers)
- Nil pointers and safe dereferencing
- Functions returning pointers
- Pointers with arrays and slices
- Common pointer patterns
- Pointer efficiency with large data

**60_strings_and_runes_detailed.go** - Strings and Runes
- String and rune basics
- UTF-8 multi-byte characters
- String to rune conversion
- String operations (split, join, contains, etc.)
- Indexing vs range on strings
- Rune properties and Unicode categories
- String modification patterns
- String processing (filtering, mapping)
- String comparison and ordering
- Performance considerations

**61_formatting_verbs_detailed.go** - Formatting Verbs
- Integer formatting verbs (%d, %b, %o, %x)
- Float formatting (%f, %e, %g)
- String and character formatting (%s, %c, %U)
- Width and padding control
- Precision control for numbers and strings
- Special format verbs (boolean, pointer, percent)
- Flags (+, -, 0, #, space)
- Complex type formatting
- Practical formatting examples
- Common mistakes and solutions

**62_fmt_package_detailed.go** - Fmt Package
- Print functions (Print, Println, Printf)
- Sprintf (formatting into strings)
- Fprintf (formatting to writers)
- Scan functions (reading input)
- Writing to different destinations (stdout, stderr, files, buffers)
- Error messages with formatting
- Complex data structure formatting
- String building patterns
- Printf variants comparison
- Practical use cases (debug, CSV, reports)
- Error handling with formatting

### Object-Oriented Programming

**63_structs_detailed.go** - Structs
- Basic struct declaration and usage
- Struct pointers and modification
- Nested structs
- Zero values and field initialization
- Anonymous/embedded structs
- Struct methods and receivers
- Constructor functions
- Custom string representation
- Composition over inheritance
- Struct comparison
- Public and private fields
- Practical examples

**64_methods_detailed.go** - Methods
- Basic methods with value receiver
- Methods with pointer receiver (mutation)
- Value vs pointer receiver comparison
- String method (Stringer interface)
- Method chaining
- Methods on built-in types
- Methods vs functions comparison
- Complex example (game character)
- Receiver type consistency
- Method documentation

**65_interfaces_detailed.go** - Interfaces
- Basic interface definition and implementation
- Multiple implementations of same interface
- Empty interface (can hold any value)
- Type assertion and type switch
- Reader and Writer interfaces
- Embedded interfaces
- Interface composition
- Interface satisfaction checking
- Shape calculator example
- Best practices

**66_struct_embedding_detailed.go** - Struct Embedding
- Basic struct embedding (composition)
- Method promotion
- Overriding embedded methods
- Multiple embedding
- Embedding with initialization
- Embedding with interfaces
- Practical handler example
- Field shadowing
- Composition vs inheritance
- Anonymous embedded structs

**67_generics_detailed.go** - Generics (Go 1.18+)
- Basic generic functions
- Multiple type parameters
- Generic data structures (Stack, Queue)
- Type constraints
- Generic types with constraints
- Methods on generic types
- Generic slice operations (Find, Filter, Map)
- Ordered type constraints
- Generic cache implementation
- Method constraints

### Error Handling & Data Formats

**68_errors_detailed.go** - Errors
- Basic error handling with errors.New
- Error strings and formatting
- Custom error types
- Type assertion with errors
- Error wrapping with %w
- errors.Is vs errors.As
- Sentinel errors
- Rich error context
- Multiple return errors
- Panic and recover
- Error checking patterns

**94_json_detailed.go** - JSON
- Basic marshaling (Go to JSON)
- Basic unmarshaling (JSON to Go)
- Struct tags (json:"name", omitempty, -)
- Nested structures
- Slices and arrays in JSON
- Maps and JSON
- Custom marshaling with MarshalJSON
- Field visibility and privacy
- Error handling for invalid JSON
- Practical API response example

## Quick Start

Each file can be run independently:

```bash
go run 57_closures_detailed.go
go run 58_recursion_detailed.go
go run 59_pointers_detailed.go
# ... and so on
```

Or compile and run:

```bash
go build 60_strings_and_runes_detailed.go
./strings_and_runes_detailed
```

## Key Features of Each File

1. **Clear Organization**: Each example is numbered and has a descriptive heading
2. **Comprehensive Comments**: Inline comments explaining what's happening
3. **Multiple Examples**: Each topic has 8-10 practical examples
4. **Output Demonstrations**: Each example shows expected output
5. **Best Practices**: Includes patterns and anti-patterns
6. **Key Takeaways**: Summary section in main function output
7. **Practical Use Cases**: Real-world applications included

## Topics Covered (57-68, 94)

| # | Topic | File |
|---|-------|------|
| 57 | Closures | 57_closures_detailed.go |
| 58 | Recursion | 58_recursion_detailed.go |
| 59 | Pointers | 59_pointers_detailed.go |
| 60 | Strings and Runes | 60_strings_and_runes_detailed.go |
| 61 | Formatting Verbs | 61_formatting_verbs_detailed.go |
| 62 | Fmt Package | 62_fmt_package_detailed.go |
| 63 | Structs | 63_structs_detailed.go |
| 64 | Methods | 64_methods_detailed.go |
| 65 | Interfaces | 65_interfaces_detailed.go |
| 66 | Struct Embedding | 66_struct_embedding_detailed.go |
| 67 | Generics | 67_generics_detailed.go |
| 68 | Errors | 68_errors_detailed.go |
| 94 | JSON | 94_json_detailed.go |

## Learning Path Recommendation

1. **Start with Foundations**: 59 (Pointers) → 63 (Structs) → 60 (Strings)
2. **Learn Methods and Interfaces**: 64 (Methods) → 65 (Interfaces) → 66 (Embedding)
3. **Control Flow**: 57 (Closures) → 58 (Recursion)
4. **Output and Data**: 61 (Verbs) → 62 (Fmt) → 94 (JSON)
5. **Advanced**: 67 (Generics) → 68 (Errors)

## Common Patterns

### Error Handling
- Check errors immediately after operations
- Wrap errors with context using %w
- Use sentinel errors for specific conditions
- Create custom error types for control

### Interface Design
- Keep interfaces small (1-3 methods)
- Define interfaces where they're used
- Use composition to build larger interfaces
- Check interface satisfaction at compile time

### Closures and Generics
- Use closures for factories and state
- Use generics for type-safe collections
- Combine closures with defer for cleanup
- Type parameters enable code reuse

## Tips for Understanding

1. **Run Each Example**: Don't just read, execute the code
2. **Modify Examples**: Change values to see different behavior
3. **Combine Concepts**: Try using pointers with interfaces, etc.
4. **Read Comments**: Comments explain the "why", not just "what"
5. **Use Your IDE**: Use Go's tools to understand types and errors

## Additional Resources

- Go Documentation: https://golang.org/doc/
- Effective Go: https://golang.org/doc/effective_go
- Go by Example: https://gobyexample.com/

## Notes

- All files are self-contained and can be run independently
- Go 1.18+ features (Generics) require recent Go version
- Each file includes 10+ examples covering basics to advanced concepts
- Output includes clear separation and organization
- Best practices and anti-patterns are demonstrated

---

Created as comprehensive learning material for intermediate Go topics.
Each file is designed to be both an educational resource and a reference.
