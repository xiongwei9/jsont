package parse

import (
	"testing"
)

func TestParse(t *testing.T) {
	obj, err := Parse(`{"key":"value"}`)
	if err != nil {
		t.Error("parse error:", err)
	}
	t.Log("parse json success:", obj)
}
