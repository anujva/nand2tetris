package assembler

import (
	"fmt"
	"os"

	"github.com/anujva/nand2tetris/generator"
	"github.com/anujva/nand2tetris/parser"
	"github.com/anujva/nand2tetris/token"
)

//HackAssembler is the assembler for HACK computer.
type HackAssembler struct {
	Outputfile *os.File
	Parser     parser.Parser
	Code       *generator.CodeGenerator
	// It will be required for the assembler to store
	// some state values. There is a need of storing line
	// numbers that have not been resolved since we don't
	// know the value of the symbol.
	varAddress       int
	lineNumber       int
	unresolved       []token.Token
	varsymunresolved bool
}

//New returns a pointer to an object of HackAssembler
func New() *HackAssembler {
	return &HackAssembler{
		Parser:           parser.New(),
		Code:             generator.New(),
		varAddress:       16,
		lineNumber:       -1,
		unresolved:       make([]token.Token, 0),
		varsymunresolved: false,
	}
}

//NewWithFile returns a HackAssemblers output file specified
func NewWithFile(file *os.File) *HackAssembler {
	return &HackAssembler{
		Outputfile: file,
		Parser:     parser.New(),
		Code:       generator.New(),
	}
}

//SetOutputFile set the output file for the assembler
func (ha *HackAssembler) SetOutputFile(file *os.File) {
	ha.Outputfile = file
}

//FirstPass will fill up the code generator states
// for variables and labels
func (ha *HackAssembler) FirstPass(str string) {
	ha.lineNumber = ha.lineNumber + 1
	tkns := ha.Parser.Parse(str)
	if len(tkns) == 1 && tkns[0].Type == token.LABEL {
		ha.Code.VarMap[tkns[0].Val] = ha.lineNumber + 1
	} else if len(tkns) == 1 && tkns[0].Type == token.VARIABLE {
		if _, ok := ha.Code.VarMap[tkns[0].Val]; !ok {
			ha.Code.VarMap[tkns[0].Val] = ha.varAddress + 1
		}
	}
}

// AssembleFile What do I need an assembler to do? It should take in file
// And spit out another file that will be the assembled output
func (ha *HackAssembler) AssembleFile(str string) {
	tkns := ha.Parser.Parse(str)
	if len(tkns) == 0 {
		// skip this.
		return
	}
	finalString := ""
	if len(tkns) == 3 {
		//This is a c instruction
		finalString1, _ := ha.Code.TranslateToken(tkns[1])
		finalString2, _ := ha.Code.TranslateToken(tkns[0])
		finalString3, _ := ha.Code.TranslateToken(tkns[2])
		finalString = finalString1 + finalString2 + finalString3
	} else if len(tkns) == 1 {
		//This is a instruction
		finalString, _ = ha.Code.TranslateToken(tkns[0])
	}

	if len(finalString) > 0 {
		fmt.Println(finalString)
	}
}
