package main

import (
	"fmt"
	"reflect"
)

func main() {
	//imap := make(map[string]string)
	//var input interface{}
	//input = imap
	//m := reflect.ValueOf(input)
	//if m.Kind() == reflect.Map {
	//	res := reflect.MakeMap(m.Type())
	//	keys := m.MapKeys()
	//	for _, k := range keys {
	//		key := k.Convert(res.Type().Key()) //.Convert(m.Type().Key())
	//		value := m.MapIndex(key)
	//		res.SetMapIndex(key, value)
	//	}
	//}
	mi := map[string]string{
		"a": "this is a",
		"b": "this is b",
	}

	var input interface{}
	input = mi
	m := reflect.ValueOf(input)
	if m.Kind() == reflect.Map {
		newInstance := reflect.MakeMap(m.Type())
		keys := m.MapKeys()
		for _, k := range keys {
			key := k.Convert(newInstance.Type().Key())
			value := m.MapIndex(key)
			newInstance.SetMapIndex(key, value)
		}
		fmt.Println(newInstance)
	}

}
