package parse

import (
	"fmt"
	"testing"
)

func TestParseNumber(t *testing.T) {
	mockInteger := 1
	integer, err := Parse(fmt.Sprint(mockInteger))
	if err != nil || integer != mockInteger {
		t.Error("parse integer error:", err)
		return
	}
	t.Log("parse integer success:", integer)
}

func TestParseString(t *testing.T) {
	mockStr := " {   hello world   } "
	str, err := Parse(`   "` + mockStr + `"   `)
	if err != nil || str != mockStr {
		t.Error("parse string error:", err)
		return
	}
	t.Log("parse string success:", str)
}

func TestParseObject(t *testing.T) {
	obj, err := Parse(`{"key":"value"}`)
	if err != nil {
		t.Error("parse error:", err)
		return
	}
	t.Log("parse json success:", obj)
}
