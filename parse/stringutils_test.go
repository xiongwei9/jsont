package parse

import "testing"

func TestPassSpace(t *testing.T) {
	str := "     hello world"
	cList := []rune(str)
	i := 0
	j := passSpace(cList, i)
	if j != 5 {
		t.Error("must be 5")
	}
	i = 10
	j = passSpace(cList, i)
	if j != 11 {
		t.Error("must be 11")
	}
	i = j
	j = passSpace(cList, i)
	if j != 11 {
		t.Error("must be 11")
	}
}

func TestGetQuoteContent(t *testing.T) {
	str := `" _key__"`
	content, _ := getQuoteContent([]rune(str), 0)
	if string(content) != " _key__" {
		t.Error("not match")
	}
}
