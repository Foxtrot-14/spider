package lexer
type Lexer struct{
	input        string
	position     int
	readPosition int
	ch           byte
}
func New(input string)*Lexer{
	l := &Lexer{input: input}
	return l
}
/*
By passing a pointer, you're essentially passing a reference to the original data in memory, rather than making a copy
*/
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}