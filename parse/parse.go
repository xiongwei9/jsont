package parse

import (
	"errors"
	"fmt"
)

type parser struct {
	index int
	str   []rune
}

func (p *parser) parseString() (string, error) {
	str, i := getQuoteContent(p.str, p.index)
	if i < 0 {
		return "", errors.New("parseString error")
	}
	p.index = i + 1
	return string(str), nil
}

func (p *parser) parseObject() (map[string]interface{}, error) {
	charList := p.str
	i := p.index
	if charList[i] != '{' {
		return nil, errors.New("parseObject error")
	}
	obj := make(map[string]interface{})
	for {
		j := passSpace(charList, i+1)
		key, j := getQuoteContent(charList, j)
		if j < 0 {
			return nil, errors.New("parseObject error: not key")
		}
		j = passSpace(charList, j+1)
		if charList[j] != ':' {
			return nil, errors.New("parseObject error: get colon error")
		}
		j = passSpace(charList, j+1)
		value, err := p.parse()
		if err != nil {
			return nil, err
		}
		obj[string(key)] = value

		if charList[j] == '}' {
			break
		}
	}
	return obj, nil
}

func (p *parser) parse() (interface{}, error) {
	var jsonVal interface{}
	charList := p.str[p.index:]
	for i := 0; i < len(charList); i++ {
		char := charList[i]
		switch char {
		case ' ':
			i = passSpace(charList, i)
			if i > 0 {
				i = i - 1
			}
		case '"':
			return p.parseString()
		case '{':
			return p.parseObject()
		case '[':
			fmt.Println("Array start")
		}
	}

	return jsonVal, nil
}

func Parse(jsonStr string) (interface{}, error) {
	p := parser{
		str:   []rune(jsonStr),
		index: 0,
	}
	return p.parse()
}
