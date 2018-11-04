package assembler

import (
	token "nand2tetris/token"
)

// CodeGenInterface defines the behavior that we
// want from the code generator. It will look at
// a token and will return a string which will be
// the machine language equivalent of it.
type CodeGenInterface interface {
	translateInstruction(token token.Token) string
}

// CodeGenerator is an implementation of the
// CodeGenInterface, will be used to work the
// strings that are read from the source code.
type CodeGenerator struct {
	tokenToMachine map[string]string
}
