package stringify

import (
	"testing"
)

func TestStringifyBasicObject(t *testing.T) {
	obj := map[string]interface{}{
		"name":   "xiongwei",
		"age":    18,
		"sex":    nil,
		"people": true,
		"weight": 199.9,
	}
	json, err := Stringify(obj)
	if err != nil {
		t.Error("stringify error:", err)
	}
	t.Log("stringify json success:", json)
}

func TestStringifyBasicArray(t *testing.T) {
	array := []interface{}{"xiongwei", 18, nil, true, 119.9}
	json, err := Stringify(array)
	if err != nil {
		t.Error("stringify error:", err)
	}
	t.Log("stringify json success:", json)
}

func TestStringifyComplexData(t *testing.T) {
	obj := map[string]interface{}{
		"name":   "xiongwei",
		"age":    18,
		"sex":    nil,
		"people": true,
		"weight": 199.9,
		"friends": []map[string]interface{}{
			{
				"name":    "xiongwei1",
				"age":     18,
				"sex":     nil,
				"people":  true,
				"weight":  199.9,
				"friends": []string{},
			},
			{
				"name":   "xiongwei2",
				"age":    18,
				"sex":    nil,
				"people": true,
				"weight": 199.9,
			},
		},
		"wifi": map[string]interface{}{
			"id":       0,
			"name":     "my wifi",
			"password": "123456",
		},
	}
	json, err := Stringify(obj)
	if err != nil {
		t.Error("stringify error:", err)
	}
	t.Log("stringify json success:", json)
}

func TestStringifyStruct(t *testing.T) {
	type Person struct {
		Name string
	}
	type Dog struct {
		Name   string
		Age    int
		Master Person
	}
	dog := Dog{
		Name:   "xx",
		Age:    8,
		Master: Person{Name: "xiongwei"},
	}
	json, err := Stringify(dog)
	if err != nil {
		t.Error("stringify error:", err)
	}
	t.Log("stringify json success:", json)
}

func TestStringifyPointer(t *testing.T) {
	type Person struct {
		Name string
	}
	listP := &[]interface{}{
		"hello",
		18,
		nil,
		false,
		&Person{Name: "xxx"},
	}
	json, err := Stringify(listP)
	if err != nil {
		t.Error("stringify error:", err)
	}
	t.Log("stringify json success:", json)
}
