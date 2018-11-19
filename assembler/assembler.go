package assembler

import (
	"os"

	"github.com/anujva/nand2tetris/generator"
	"github.com/anujva/nand2tetris/parser"
)

//HackAssembler is the assembler for HACK computer.
type HackAssembler struct {
	outputfile *os.File
	parser     Parser
	code       CodeGenerator
}

//New returns a pointer to an object of HackAssembler
func New() *HackAssembler {
	return &HackAssembler{
		parser: parser.New(),
		code:   generator.New(),
	}
}

//NewWithFile returns a HackAssemblers output file specified
func NewWithFile(file *os.File) *HackAssembler {
	return &HackAssembler{
		outputfile: file,
		parser:     parser.New(),
		code:       generator.New(),
	}
}

//SetOutputFile set the output file for the assembler
func (ha *HackAssembler) SetOutputFile(file *os.File) {
	ha.outputfile = file
}

// What do I need an assembler to do? It should take in file
// And spit out another file that will be the assembled output

func (ha *HackAssembler) assembleFile() {

}
