package assembler

import "os"

//HackAssembler is the assembler for HACK computer.
type HackAssembler struct {
	outputfile *os.File
	parser     Parser
	code       CodeGenerator
}

//New returns a pointer to an object of HackAssembler
func New() *HackAssembler {
	immutableSymbolTable := initializeSymbolTable()
	return &HackAssembler{}
}

//NewWithFile returns a HackAssemblers output file specified
func NewWithFile(file *os.File) *HackAssembler {
	immutableSymbolTable := initializeSymbolTable()
	return &HackAssembler{
		outputfile: file,
	}
}

//SetOutputFile set the output file for the assembler
func (ha *HackAssembler) SetOutputFile(file *os.File) {
	ha.outputfile = file
}
