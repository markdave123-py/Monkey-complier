package lexer

import "compiler/token"

// type Lexer struct {
// 	cursor int
// 	position int
// 	character byte
// 	inputStream string
// }

type Lexer struct{
	input string
	next int  // point to the next character to be read(next after position)
	pos int // points to the character just read
	ch byte // current character under examination
}

func New(input string) *Lexer{
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {

	if l.next >= len(l.input){
		l.ch = 0
	}else{
		l.ch = l.input[l.next]
	}

	l.pos = l.next
	l.next += 1
}

func (l *Lexer) NextToken() token.Token{

	var tok token.Token

	switch l.ch {
		case '=':
			tok = newToken(token.ASSIGN, l.ch)
		case ';':
			tok = newToken(token.SEMICOLON, l.ch)
		case '(':
			tok = newToken(token.LPAREN, l.ch)
		case ')':
			tok = newToken(token.RPAREN, l.ch)
		case ',':
			tok = newToken(token.COMMA, l.ch)
		case '+':
			tok = newToken(token.PLUS, l.ch)
		case '{':
			tok = newToken(token.LBRACE, l.ch)
		case '}':
			tok = newToken(token.RBRACE, l.ch)
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF
	}

	l.readChar()

	return tok

}

func newToken(tokenType token.TokenType, ch byte) token.Token{
	return token.Token{Type: tokenType, Literal: string(ch)}
}



// func New(sourceCode string) *Lexer {
// 	return &Lexer{ inputStream: sourceCode }
// }

// func IsNumber(ch byte) bool { 	return '0' <= ch && ch <= '9'
// }