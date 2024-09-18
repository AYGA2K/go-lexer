# Simple Lexer in Go

<!--toc:start-->

- [Simple Lexer in Go](#simple-lexer-in-go)
  - [Introduction](#introduction)
  - [Steps Followed to Build the Lexer](#steps-followed-to-build-the-lexer)
    - [1. **Define Token Types**](#1-define-token-types)
    - [2. **Create the `Token` Struct**](#2-create-the-token-struct)
    - [3. **Define a Lexer Structure**](#3-define-a-lexer-structure)
    - [4. **Initialize Keywords**](#4-initialize-keywords)
    - [5. **Implement the Lexer**](#5-implement-the-lexer)
    - [6. **Handle Special Cases**](#6-handle-special-cases)
    - [7. **Whitespace Handling**](#7-whitespace-handling)
    - [8. **EOF Handling**](#8-eof-handling)
  - [Example Usage](#example-usage)
  - [Running the Lexer](#running-the-lexer)
  - [Conclusion](#conclusion)
  <!--toc:end-->

## Introduction

A lexer is the first stage of a compiler or interpreter. Its job is to read the source code, break it down into a series of tokens, and pass these tokens to the next phase (the parser).

Tokens are categorized into different types (keywords, operators, literals, etc.) and serve as the fundamental building blocks for the source code's structure.

In this project, we have implemented a lexer that can tokenize basic programming constructs such as keywords, operators, delimiters, and identifiers.

---

## Steps Followed to Build the Lexer

### 1. **Define Token Types**

We started by defining the types of tokens that the lexer will recognize. This includes: - Keywords (`if`, `else`, `for`, `return`, etc.) - Operators (`+`, `-`, `*`, `/`, `=`, `!=`, `==`) - Delimiters (`;`, `,`, `(`, `)`, `{`, `}`) - Identifiers (user-defined names for variables and functions) - Literals (integers)

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

Will be tokenized into the following series of tokens:

```bash

{Typ:func Literal:func}
{Typ:IDENT Literal:test}
{Typ:( Literal:(}
{Typ:IDENT Literal:x}
{Typ:, Literal:,}
{Typ:IDENT Literal:y}
{Typ:IDENT Literal:int}
{Typ:) Literal:)}
{Typ:IDENT Literal:int}
{Typ:{ Literal:{}
{Typ:if Literal:if}
{Typ:IDENT Literal:x}
{Typ:== Literal:==}
{Typ:INT Literal:10}
{Typ:{ Literal:{}
{Typ:var Literal:var}
{Typ:IDENT Literal:z}
{Typ:= Literal:=}
{Typ:IDENT Literal:y}
{Typ:_Literal:_}
{Typ:INT Literal:2}
{Typ:IDENT Literal:for}
{Typ:IDENT Literal:z}
{Typ:!= Literal:!=}
{Typ:INT Literal:0}
{Typ:{ Literal:{}
{Typ:IDENT Literal:z}
{Typ:= Literal:=}
{Typ:IDENT Literal:z}
{Typ:- Literal:-}
{Typ:INT Literal:1}
{Typ:} Literal:}}
{Typ:return Literal:return}
{Typ:IDENT Literal:z}
{Typ:+ Literal:+}
{Typ:IDENT Literal:x}
{Typ:} Literal:}}
{Typ:else Literal:else}
{Typ:if Literal:if}
{Typ:IDENT Literal:y}
{Typ:!= Literal:!=}
{Typ:INT Literal:5}
{Typ:{ Literal:{}
{Typ:return Literal:return}
{Typ:IDENT Literal:y}
{Typ:/ Literal:/}
{Typ:IDENT Literal:x}
{Typ:} Literal:}}
{Typ:else Literal:else}
{Typ:{ Literal:{}
{Typ:return Literal:return}
{Typ:INT Literal:0}
{Typ:} Literal:}}
{Typ:} Literal:}}


```

---

## Running the Lexer

To run the lexer:

1. Clone this repository:

   ```bash
   git clone https://github.com/AYGA2K/go-lexer.git
   cd go-lexer
   ```

2. Build and run the lexer:

   ```bash
   go run main.go
   ```

3. Modify the `input` string in `main.go` inside the code directory with your test code to see how the lexer tokenizes it.

---

## Conclusion

This project demonstrates how to implement a simple lexer from scratch in Go. By following the steps above, we've created a tool that can tokenize basic programming constructs and lay the foundation for a full parser or interpreter. This lexer can be extended further to support more complex syntax and language features.
