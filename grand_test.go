package grand

import (
	"regexp"
	"testing"
)

func TestGeneratesDifferentStrings(t *testing.T) {
	randStr1 := String(20)

	if len(randStr1) != 20 {
		t.Fatal()
	}

	randStr2 := String(20)
	if randStr1 == randStr2 {
		t.Fatal()
	}

	randStr3 := String(20)
	if randStr2 == randStr3 {
		t.Fatal()
	}
}

func TestSetDefaultCharSet(t *testing.T) {
	randStr1 := String(20)

	// There should not be any numbers in the string
	res := regexp.MustCompile("\\d").FindAllString(randStr1, -1)
	if len(res) != 0 {
		t.Fatal()
	}
}

func TestLcString(t *testing.T) {
	randStr1 := LcString(20)

	// There should not be any uppercase in the string
	res := regexp.MustCompile("[A-Z]").FindAllString(randStr1, -1)
	if len(res) != 0 {
		t.Fatal()
	}
}

func TestSetDifferentCharSet(t *testing.T) {
	gen := New(CharSetBase62)

	randStr1 := gen.String(20)

	// There should be at least one number in generated string
	res := regexp.MustCompile("\\d").FindAllString(randStr1, -1)
	if len(res) == 0 {
		t.Fatal()
	}
}
