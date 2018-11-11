package assembler

import (
	token "github.com/anujva/nand2tetris/token"
)

func initializeSymbolTable() map[string]string {
	immutableSymbolTable := map[string]int{
		"R0":     0,
		"R1":     1,
		"R2":     2,
		"R3":     3,
		"R4":     4,
		"R5":     5,
		"R6":     6,
		"R7":     7,
		"R8":     8,
		"R9":     9,
		"R10":    10,
		"R11":    11,
		"R12":    12,
		"R13":    13,
		"R14":    14,
		"R15":    15,
		"SCREEN": 16384,
		"KBD":    24576,
		"SP":     0,
		"LCL":    1,
		"ARG":    2,
		"THIS":   3,
		"THAT":   4,
	}
	return immutableSymbolTable
}

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

func (cg *CodeGenerator) translateInstruction(token token.Token) string {
	// The code generator will look at the token and translate it into
	// string.
}
