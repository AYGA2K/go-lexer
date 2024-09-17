package lexer

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	INT   = "INT"
	IDENT = "IDENT"

	// Keywords
	IF       = "if"
	ELSE     = "else"
	FOR      = "for"
	RETURN   = "return"
	FUNCTION = "func"
	VAR      = "var"
	TRUE     = "true"
	FALSE    = "false"

	// Operators
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	ASSIGN   = "="
	EQ       = "=="
	NOT_EQ   = "!="
	BANG     = "!"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
)

type Token struct {
	Typ     string
	Literal string
}

var keywords = map[string]string{
	"func":   FUNCTION,
	"var":    VAR,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func searchIdent(ident string) string {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
