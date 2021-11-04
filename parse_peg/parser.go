package parse_peg

//go:generate pigeon -o grammar.peg.go ./grammar.peg

import (
	"io"
	"io/ioutil"
	"strings"
)

type namedReader interface {
	Name() string
}

func ParseJsonFile(r io.Reader) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	name := "<reader>"
	if named, ok := r.(namedReader); ok {
		name = named.Name()
	}
	t, err := Parse(name, b)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// TODO: return a pointer for better performance?
func ParseJson(contents string) (interface{}, error) {
	return ParseJsonFile(strings.NewReader(contents))
}
