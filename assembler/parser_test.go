package assembler

import (
	"strings"
	"testing"
)

func TestRemoveInsideWhiteSpaces(t *testing.T) {
	str1 := removeWhiteSpaces("abc d")
	if strings.Compare(str1, "abcd") != 0 {
		t.Fail()
	}
	str2 := removeWhiteSpaces("   a   b    c   d    ")
	if strings.Compare(str2, "abcd") != 0 {
		t.Fail()
	}
}
