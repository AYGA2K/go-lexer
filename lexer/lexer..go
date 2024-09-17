package lexer

type Lexer struct {
	input   string
	curPos  int
	nextPos int
	char    byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.nextChar()
	return l
}

func (l *Lexer) NextToken() Token {
	var token Token
	l.skipWhiteSpaces()

	switch l.char {
	case '=':
		if l.peekChar() == '=' {
			char := l.char
			l.nextChar()
			literal := string(char) + string(l.char)
			token = Token{EQ, literal}
		} else {
			token = Token{ASSIGN, string(l.char)}
		}
	case '+':
		token = Token{PLUS, string(l.char)}
	case '-':
		token = Token{MINUS, string(l.char)}
	case '*':
		token = Token{ASTERISK, string(l.char)}
	case '/':
		token = Token{SLASH, string(l.char)}
	case '!':
		if l.peekChar() == '=' {
			char := l.char
			l.nextChar()
			literal := string(char) + string(l.char)
			token = Token{NOT_EQ, literal}
		} else {
			token = Token{BANG, string(l.char)}
		}
	case ',':
		token = Token{COMMA, string(l.char)}
	case ';':
		token = Token{SEMICOLON, string(l.char)}
	case '(':
		token = Token{LPAREN, string(l.char)}
	case ')':
		token = Token{RPAREN, string(l.char)}
	case '{':
		token = Token{LBRACE, string(l.char)}
	case '}':
		token = Token{RBRACE, string(l.char)}
	case 0:
		token.Literal = ""
		token.Typ = EOF
	default:
		if isLetter(l.char) {
			literal := l.readWord()
			tokenType := searchIdent(literal)
			return Token{tokenType, literal}
		} else if isDigit(l.char) {
			literal := l.readNumber()
			return Token{INT, literal}
		} else {
			token = Token{ILLEGAL, string(l.char)}
		}
	}

	l.nextChar()
	return token
}
func (l *Lexer) nextChar() {
	if l.nextPos >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.nextPos]
		l.curPos = l.nextPos
		l.nextPos++
	}
}
func (l *Lexer) skipWhiteSpaces() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.nextChar()
	}
}
func (l *Lexer) peekChar() byte {
	if l.nextPos >= len(l.input) {
		l.char = 0
	}
	return l.input[l.nextPos]
}

func (l *Lexer) readNumber() string {
	pos := l.curPos
	for isDigit(l.char) {
		l.nextChar()
	}
	return l.input[pos:l.curPos]
}
func (l *Lexer) readWord() string {
	pos := l.curPos
	for isLetter(l.char) {
		l.nextChar()
	}
	return l.input[pos:l.curPos]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
