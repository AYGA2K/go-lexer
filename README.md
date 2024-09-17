# Simple Lexer in Go

## Introduction

A lexer is the first stage of a compiler or interpreter. Its job is to read the source code, break it down into a series of tokens, and pass these tokens to the next phase (the parser).

Tokens are categorized into different types (keywords, operators, literals, etc.) and serve as the fundamental building blocks for the source code's structure.

In this project, we have implemented a lexer that can tokenize basic programming constructs such as keywords, operators, delimiters, and identifiers.

---

## Steps Followed to Build the Lexer

### 1. Define Token Types

    We started by defining the types of tokens that the lexer will recognize. This includes:
    - Keywords (`if`, `else`, `for`, `return`, etc.)
    - Operators (`+`, `-`, `*`, `/`, `=`, `!=`, `==`)
    - Delimiters (`;`, `,`, `(`, `)`, `{`, `}`)
    - Identifiers (user-defined names for variables and functions)
    - Literals (integers)

### 2. **Create the `Token` Struct**

    The `Token` struct has two fields: `Typ` (the token type) and `Literal` (the value of the token, such as the number or keyword).

### 3. **Define a Lexer Structure**

    The lexer keeps track of the input string and the position within it. It reads one character at a time and converts them into tokens based on the rules we defined.

### 4. **Initialize Keywords**

    We created a keyword map to differentiate between identifiers and keywords. If a token matches one of the keywords, it is classified as such. Otherwise, it is treated as an identifier.

### 5. **Implement the Lexer**

    The lexer reads the input character by character using the `nextChar` method. Based on the current character, it matches the token type (keyword, operator, delimiter, etc.) and moves to the next character to continue the process.

### 6. **Handle Special Cases**

    Special cases, such as multi-character operators (`==`, `!=`), are handled by looking ahead with the `peekChar` method. The lexer checks the next character before deciding which token type to assign.

### 7. **Whitespace Handling**

    We implemented a method `skipWhiteSpaces` to ignore spaces, tabs, and newlines between tokens, ensuring that tokens are not affected by unnecessary whitespace.

### 8. **EOF Handling**

    The lexer properly detects the end of the input (EOF) and returns an `EOF` token to indicate that there are no more tokens to process.

---

## Example Usage

The lexer can tokenize a simple Go-like program. For instance, the following code:

```go
func test(x, y int) int {
	if x == 10 {
		var z = y * 2
		for z != 0 {
			z = z - 1
		}
		return z + x
	} else if y != 5 {
		return y / x
	} else {
		return 0
	}
}
```

Will be tokenized into a series of tokens, each representing a part of the source code like keywords, operators, identifiers, and literals.

---

## Running the Lexer

To run the lexer:

1.  Clone this repository:

    ```bash
    git clone https://github.com/AYGA2K/go-lexer.git
    cd go-lexer
    ```

2.  Build and run the lexer:

    ```bash
    go run main.go
    ```

3.  Modify the `input` string in `main.go` inside the code directory with your test code to see how the lexer tokenizes it.

---

## Conclusion

This project demonstrates how to implement a simple lexer from scratch in Go. By following the steps above, we've created a tool that can tokenize basic programming constructs and lay the foundation for a full parser or interpreter. This lexer can be extended further to support more complex syntax and language features.
