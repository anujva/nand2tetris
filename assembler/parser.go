package assembler

import (
	"regexp"
	"strings"

	token "github.com/anujva/nand2tetris/token"
)

//Parser for parsing the hack program
type Parser interface {
	parse(line string) []token.Token
}

//HackParser is an implementation of Parser
//and knows how to parse for hack language specs
type HackParser struct {
	line string
}

func (h *HackParser) parse(line string) []token.Token {
	wr := removeWhiteSpaces(line)
	if len(wr) == 0 {
		return nil
	}
	ss := sliceIntoTokenStrings(wr)
	return convertToTokens(ss)
}

func removeWhiteSpaces(line string) string {
	//How do I work with strings in golang.. this will be a learning
	//exercise as well for me.
	line = strings.TrimSpace(line)

	//check if the string starts with //
	if strings.HasPrefix(line, "//") {
		return ""
	}

	splitOnSlashes := strings.Split(line, "//")

	aOrCInstructionWithWhiteSpace := strings.TrimSpace(splitOnSlashes[0])
	finalString := removeInsideWhiteSpaces(aOrCInstructionWithWhiteSpace)

	return finalString
}

func removeInsideWhiteSpaces(str string) string {
	reInsideWs := regexp.MustCompile(`[\s\p{Zs}]`)
	return reInsideWs.ReplaceAllString(str, "")
}

func sliceIntoTokenStrings(wr string) []string {
	// What should we do here.. we have the string without any ws
	// so the idea should be ok
	// At this point we knwo that each line that we get will need
	// to be split up into either a instruction or c instruction.

	switch c := isAInstruction(wr); c {
	case true:
		return []string{"@", wr[1:]}
	case false:
		return splitCInstruction(wr)
	}
	return nil
}

func splitCInstruction(wr string) []string {
	result = make([]string, 0)
	intermediate := strings.Split(wr, "=")
	append(result, intermediate[0])
	final := strings.Split(intermediate, ";")
	append(result, final[0])
	append(result, final[1])
	return result
}

func isAInstruction(wr string) bool {
	return strings.HasPrefix(wr, "@")
}

func convertToTokens(ss []string) []token.Token {
	return nil
}
