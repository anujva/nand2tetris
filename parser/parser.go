package parser

import (
	"regexp"
	"strconv"
	"strings"

	token "github.com/anujva/nand2tetris/token"
)

//Parser for parsing the hack program
type Parser interface {
	Parse(line string) []token.Token
}

// New returns an implementation of Parser
func New() Parser {
	return &HackParser{}
}

//HackParser is an implementation of Parser
//and knows how to parse for hack language specs
type HackParser struct {
}

// Parse the line into Hack Tokens
func (h *HackParser) Parse(line string) []token.Token {
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

// splitCInstruction will split the instruction into three parts
// DEST=CMP;JMP
// Apart from the normal values, it is possible for the slice of
// string to return null, CMP, null
func splitCInstruction(wr string) []string {
	result := make([]string, 0)
	intermediate := strings.Split(wr, "=")
	if len(intermediate) != 2 {
		//Probably destination is missing the instruction
		intermediate = append([]string{"null"}, intermediate...)
	}
	result = append(result, intermediate[0])
	final := strings.Split(intermediate[1], ";")
	result = append(result, final[0])
	if len(final) == 2 && final[1] != "" {
		result = append(result, final[1])
	} else {
		result = append(result, "null")
	}
	return result
}

func isAInstruction(wr string) bool {
	return strings.HasPrefix(wr, "@")
}

func convertToTokens(ss []string) []token.Token {
	if ss[0] == "@" {
		if _, err := strconv.Atoi(ss[1]); err == nil {
			return []token.Token{
				token.Token{
					token.ADDRESS,
					ss[1],
				},
			}
		}
		return []token.Token{
			token.Token{
				token.SYMBOL,
				ss[1],
			},
		}
	}

	// this is a c instruction
	return []token.Token{
		token.Token{
			token.DEST,
			ss[0],
		},
		token.Token{
			token.COMP,
			ss[1],
		},
		token.Token{
			token.JUMP,
			ss[2],
		},
	}
}
