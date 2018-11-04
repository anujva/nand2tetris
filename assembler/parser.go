package assembler

import (
	token "nand2tetris/token"
	"regexp"
	"strings"
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
	return nil
}

func convertToTokens(ss []string) []token.Token {
	return nil
}
