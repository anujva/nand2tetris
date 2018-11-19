package assembler

import (
	"os"

	"github.com/anujva/nand2tetris/generator"
	"github.com/anujva/nand2tetris/parser"
)

//HackAssembler is the assembler for HACK computer.
type HackAssembler struct {
	Outputfile *os.File
	Parser     parser.Parser
	Code       generator.CodeGenInterface
}

//New returns a pointer to an object of HackAssembler
func New() *HackAssembler {
	return &HackAssembler{
		Parser: parser.New(),
		Code:   generator.New(),
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

// What do I need an assembler to do? It should take in file
// And spit out another file that will be the assembled output
func (ha *HackAssembler) assembleFile() {

}
