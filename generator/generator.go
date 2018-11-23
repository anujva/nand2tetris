package generator

import (
	"fmt"
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
	TranslateToken(token token.Token) (string, error)
}

// New returns an implementation of the code generator
func New() CodeGenInterface {
	return &codeGenerator{
		destMap:    getDestMap(),
		jumpMap:    getJumpMap(),
		compMap:    getCompMap(),
		predefMap:  initializeSymbolTable(),
		varMap:     make(map[string]int),
		varAddress: 16,
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
	compMap["0"] = "1110101010"
	compMap["1"] = "1110111111"
	compMap["-1"] = "1110111010"
	compMap["D"] = "1110001100"
	compMap["A"] = "1110110000"
	compMap["!D"] = "1110001101"
	compMap["!A"] = "1110110001"
	compMap["-D"] = "1110001111"
	compMap["-A"] = "1110110011"
	compMap["D+1"] = "1110011111"
	compMap["A+1"] = "1110110111"
	compMap["D-1"] = "1110001110"
	compMap["A-1"] = "1110110010"
	compMap["D+A"] = "1110000010"
	compMap["D-A"] = "1110010011"
	compMap["A-D"] = "1110000111"
	compMap["D&A"] = "1110000000"
	compMap["D|A"] = "1110010101"
	compMap["M"] = "1111110000"
	compMap["!M"] = "1111110001"
	compMap["-M"] = "1111110011"
	compMap["M+1"] = "1111110111"
	compMap["M-1"] = "1111110010"
	compMap["D+M"] = "1111000010"
	compMap["D-M"] = "1111010011"
	compMap["M-D"] = "1111000111"
	compMap["D&M"] = "1111000000"
	compMap["D|M"] = "1111010101"
	return compMap
}

// CodeGenerator is an implementation of the
// CodeGenInterface, will be used to work the
// strings that are read from the source code.
type codeGenerator struct {
	destMap   map[string]string
	jumpMap   map[string]string
	compMap   map[string]string
	predefMap map[string]int
}

func (cg *codeGenerator) getAddressString(add string) (string, error) {
	val := 0
	var err error
	// Check if its a predefined symbol
	if val, ok := cg.predefMap[add]; ok {
		return getAsBinaryString(val), nil
	}
	// Address has to of type string
	val, err = strconv.Atoi(add)
	if err == nil {
		return getAsBinaryString(val), nil
	}
	// It is not a number or a predefined symbol
	// Which means it is a variable or it is a
	// label symbol.. we need to see if we know
	// the label already and have kept it in either
	// the varMap or the labelMap
	var ok bool
	if val, ok = cg.varMap[add]; ok {
		return getAsBinaryString(val), nil
	}

	if val, ok = cg.labelMap[add]; ok {
		return getAsBinaryString(val), nil
	}
	return "", &labelUnknown{add, "label is unknown"}
}

type labelUnknown struct {
	add  string
	prob string
}

func (l *labelUnknown) Error() string {
	return fmt.Sprintf("%s - %s", l.add, l.prob)
}

func getAsBinaryString(val int) string {
	binaryString := getAsBinaryStringSubroutine(val)
	//We will need to add a zero for the the binaryString
	//then append it with as many zeros as need to make it
	//15 bit long string
	lenApp := 15 - len(binaryString)
	for i := 0; i < lenApp; i++ {
		binaryString = "0" + binaryString
	}
	binaryString = "0" + binaryString
	return binaryString
}

func getAsBinaryStringSubroutine(val int) string {
	if val == 0 {
		return ""
	}

	if val%2 == 0 {
		return getAsBinaryStringSubroutine(val/2) + "0"
	}
	return getAsBinaryStringSubroutine(val/2) + "1"
}

func (cg *codeGenerator) TranslateToken(t token.Token) (string, error) {
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
