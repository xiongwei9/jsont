package parse

import (
	"testing"
)

func TestParse(t *testing.T) {
	_, err := Parse("")
	if err != nil {
		t.Error("parse error")
	}
}
