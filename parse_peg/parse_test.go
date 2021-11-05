package parse_peg

import (
	"encoding/json"
	"fmt"
	"testing"
)

//go:generate pigeon -o grammar.peg.go ./grammar.peg

// TODO: complete test code
func TestParseArray(t *testing.T) {
	s := ` [ ] `
	j, err := ParseJson(s)
	if err != nil {
		t.Error("error not null string")
		return
	}
	t.Logf("xxx: %v", j)

	s = ` [ 1 ] `
	j, err = ParseJson(s)
	if err != nil {
		t.Error("error not null string")
		return
	}
	t.Logf("xxx: %v", j)

	s = ` [ 1 , "hello", null, true, false, [], [ 1 , "hello", null, true, false, [123]  ] ] `
	j, err = ParseJson(s)
	if err != nil {
		t.Error("error not null string")
		return
	}
	t.Logf("xxx: %v", j)

	str, err := json.Marshal(j)
	if err != nil {
		t.Error("error not null string")
		return
	}
	t.Logf("%s", str)

	// if j != s {
	// 	t.Errorf("error not string: %s", s)
	// 	return
	// }
	t.Logf("%v", j)
}

func TestParseNull(t *testing.T) {
	null, err := ParseJson("   null   ")
	if err != nil || null != nil {
		t.Error("parse null error:", err)
		return
	}
	t.Log("parse null success:", null)
}

func TestParseNumber(t *testing.T) {
	mockInteger := int64(1)
	integer, err := ParseJson(fmt.Sprint(mockInteger))
	if err != nil || integer != mockInteger {
		t.Error("parse integer error:", err)
		return
	}
	t.Log("parse integer success:", integer)
}

func TestParseBoolean(t *testing.T) {
	mock := false
	parsedVal, err := ParseJson(fmt.Sprint(mock))
	if err != nil || parsedVal != mock {
		t.Error("parse boolean error:", err)
		return
	}
	t.Log("parse boolean success:", parsedVal)
}

func TestParseString(t *testing.T) {
	mockStr := " {   \\\"hello world   } \\\""
	str, err := ParseJson(`   "` + mockStr + `"   `)
	if err != nil || str != mockStr {
		t.Error("parse string error:", err)
		return
	}
	t.Log("parse string success:", str)
}

func TestParseObject(t *testing.T) {
	obj, err := ParseJson(`{"key":"value"}`)
	if err != nil {
		t.Error("parse error:", err)
		return
	}
	t.Log("parse json success:", obj)
}
