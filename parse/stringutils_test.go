package parse

import "testing"

func TestPassSpace(t *testing.T) {
	str := "     hello world"
	cList := []rune(str)
	i := 0
	j := passSpace(cList, i)
	if j != 5 {
		t.Error("must be 5")
		return
	}
	i = 10
	j = passSpace(cList, i)
	if j != 11 {
		t.Error("must be 11")
		return
	}
	i = j
	j = passSpace(cList, i)
	if j != 11 {
		t.Error("must be 11")
		return
	}
}

func TestGetQuoteContent(t *testing.T) {
	str := `" _key__"`
	content, i := getQuoteContent([]rune(str), 0)
	if string(content) != " _key__" || i < 0 {
		t.Error("not match")
		return
	}

	str = `" _key`
	_, i = getQuoteContent([]rune(str), 0)
	if i > 0 {
		t.Error("not match")
		return
	}
}
