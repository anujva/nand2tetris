package generator

import (
	"strconv"

	token "github.com/anujva/nand2tetris/token"
)

func initializeSymbolTable() map[string]int {
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
	translateToken(token token.Token) (string, error)
}

// New returns an implementation of the code generator
func New() CodeGenInterface {
	return &codeGenerator{
		destMap: getDestMap(),
		jumpMap: getJumpMap(),
		compMap: getCompMap(),
	}
}

func getDestMap() map[string]string {
	destMap := make(map[string]string)
	destMap["null"] = "000"
	destMap["M"] = "001"
	destMap["D"] = "010"
	destMap["MD"] = "011"
	destMap["A"] = "100"
	destMap["AM"] = "101"
	destMap["AD"] = "110"
	destMap["AMD"] = "111"
	return destMap
}

func getJumpMap() map[string]string {
	jumpMap := make(map[string]string)
	jumpMap["null"] = "000"
	jumpMap["JGT"] = "001"
	jumpMap["JEQ"] = "010"
	jumpMap["JGE"] = "011"
	jumpMap["JLT"] = "100"
	jumpMap["JNE"] = "101"
	jumpMap["JLE"] = "110"
	jumpMap["JMP"] = "111"
	return jumpMap
}

func getCompMap() map[string]string {
	compMap := make(map[string]string)
	compMap["0"] = "0101010"
	compMap["1"] = "0111111"
	compMap["-1"] = "0111010"
	compMap["D"] = "0001100"
	compMap["A"] = "0110000"
	compMap["!D"] = "0001101"
	compMap["!A"] = "0110001"
	compMap["-D"] = "0001111"
	compMap["-A"] = "0110011"
	compMap["D+1"] = "0011111"
	compMap["A+1"] = "0110111"
	compMap["D-1"] = "0001110"
	compMap["A-1"] = "0110010"
	compMap["D+A"] = "0000010"
	compMap["D-A"] = "0010011"
	compMap["A-D"] = "0000111"
	compMap["D&A"] = "0000000"
	compMap["D|A"] = "0010101"
	compMap["M"] = "1110000"
	compMap["!M"] = "1110001"
	compMap["-M"] = "1110011"
	compMap["M+1"] = "1110111"
	compMap["M-1"] = "1110010"
	compMap["D+M"] = "1000010"
	compMap["D-M"] = "1010011"
	compMap["M-D"] = "1000111"
	compMap["D&M"] = "1000000"
	compMap["D|M"] = "1010101"
	return compMap
}

// CodeGenerator is an implementation of the
// CodeGenInterface, will be used to work the
// strings that are read from the source code.
type codeGenerator struct {
	destMap map[string]string
	jumpMap map[string]string
	compMap map[string]string
}

func (cg *codeGenerator) getAddressString(add string) (string, error) {
	// Address has to of type string
	valAsInt, err := strconv.Atoi(add)
	if err != nil {
		return "", err
	}
	binaryString := getAsBinaryString(valAsInt)

	//We will need to add a zero for the the binaryString
	//then append it with as many zeros as need to make it
	//15 bit long string
	lenApp := 15 - len(binaryString)
	for i := 0; i < lenApp; i++ {
		binaryString = "0" + binaryString
	}
	binaryString = "0" + binaryString
	return binaryString, nil
}

func getAsBinaryString(val int) string {
	if val == 0 {
		return ""
	}

	if val%2 == 0 {
		return getAsBinaryString(val/2) + "0"
	}
	return getAsBinaryString(val/2) + "1"
}

func (cg *codeGenerator) translateToken(t token.Token) (string, error) {
	// The code generator will look at the token and translate it into
	// string.
	switch t.Type {
	case token.DEST:
		// lookup the dest map to find the string to return
		return cg.destMap[t.Val], nil
	case token.COMP:
		return cg.compMap[t.Val], nil
	case token.JUMP:
		return cg.jumpMap[t.Val], nil
	case token.ADDRESS:
		s, err := cg.getAddressString(t.Val)
		if err != nil {
			return "", err
		}
		return s, nil
	}
	return "", nil
}
