package token
//custom type called TokenType
type TokenType string
//struct named token has two fields, Type and Literal
type Token struct {
	Type TokenType
	Literal string
}
const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	//Identifier
	IDENT = "IDENT" //name,x,y
	INT = "INT" //12346
	//Operators
	ASSIGN = "="
	PLUS = "+"
	//Delimiters
	COMMA = ","
	SEMICOLON = ";"
	//Blocks
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	//Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)
var keywords = map[string]TokenType{
	"fn":
	FUNCTION,
	"let": LET,
}
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}