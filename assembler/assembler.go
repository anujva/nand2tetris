package assembler

import "os"

//HackAssembler is the assembler for HACK computer.
type HackAssembler struct {
	immutableSymbolTable map[string]string
	outputfile           *os.File
	parser               Parser
	code                 CodeGenerator
}

func initializeSymbolTable() map[string]string {
	immutableSymbolTable := map[string]string{
		"R0":     "0",
		"R1":     "1",
		"R2":     "2",
		"R3":     "3",
		"R4":     "4",
		"R5":     "5",
		"R6":     "6",
		"R7":     "7",
		"R8":     "8",
		"R9":     "9",
		"R10":    "10",
		"R11":    "11",
		"R12":    "12",
		"R13":    "13",
		"R14":    "14",
		"R15":    "15",
		"SCREEN": "16384",
		"KBD":    "24576",
		"SP":     "0",
		"LCL":    "1",
		"ARG":    "2",
		"THIS":   "3",
		"THAT":   "4",
	}
	return immutableSymbolTable
}

//New returns a pointer to an object of HackAssembler
func New() *HackAssembler {
	immutableSymbolTable := initializeSymbolTable()
	return &HackAssembler{
		immutableSymbolTable: immutableSymbolTable,
	}
}

//NewWithFile returns a HackAssemblers output file specified
func NewWithFile(file *os.File) *HackAssembler {
	immutableSymbolTable := initializeSymbolTable()
	return &HackAssembler{
		immutableSymbolTable: immutableSymbolTable,
		outputfile:           file,
	}
}

//SetOutputFile set the output file for the assembler
func (ha *HackAssembler) SetOutputFile(file *os.File) {
	ha.outputfile = file
}
