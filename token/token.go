package token

// Type defines the kinds of strings that we can see in hack asm
type Type int

// Token is defined  to parse the strings that we see in assembly
const (
	DEST Type = 1 << iota
	COMP
	JUMP
	ADDRESS
)

// Token defines the type of the string that we are evaluating
type Token struct {
	Type
	val string
}
