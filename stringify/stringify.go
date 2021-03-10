package stringify

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func stringifyStruct(structVal reflect.Value, structType reflect.Type) (string, error) {
	if structVal.Kind() != reflect.Struct {
		return "", errors.New("value is not struct")
	}
	fieldNum := structVal.NumField()
	strList := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		if !structVal.Field(i).CanInterface() {
			continue
		}
		key := structType.Field(i).Name
		value := structVal.Field(i).Interface()
		valueString, err := stringify(value)
		if err != nil {
			return "", err
		}
		str := fmt.Sprintf(`"%s":%s`, key, valueString)
		strList = append(strList, str)
	}
	return "{" + strings.Join(strList, ",") + "}", nil
}

func stringifyMap(mapVal reflect.Value) (string, error) {
	if mapVal.Kind() != reflect.Map {
		return "", errors.New("value is not array or slice")
	}
	mapLen := mapVal.Len()
	strList := make([]string, 0, mapLen)
	it := mapVal.MapRange()
	for it.Next() {
		key := it.Key().String()
		value := it.Value().Interface()
		valueString, err := stringify(value)
		if err != nil {
			return "", err
		}
		str := fmt.Sprintf(`"%s":%s`, key, valueString)
		strList = append(strList, str)
	}

	return "{" + strings.Join(strList, ",") + "}", nil
}

func stringifyArray(arrayVal reflect.Value) (string, error) {
	if arrayVal.Kind() != reflect.Array && arrayVal.Kind() != reflect.Slice {
		return "", errors.New("value is not array or slice")
	}
	arrayLen := arrayVal.Len()
	strList := make([]string, 0, arrayLen)
	for i := 0; i < arrayLen; i++ {
		ele := arrayVal.Index(i).Interface()

		str, err := stringify(ele)
		if err != nil {
			return "", err
		}
		strList = append(strList, str)
	}
	return "[" + strings.Join(strList, ",") + "]", nil
}

func stringify(jsonObj interface{}) (string, error) {
	if jsonObj == nil {
		return "null", nil
	}

	switch jsonObj.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", jsonObj), nil
	case float32, float64:
		return fmt.Sprintf("%f", jsonObj), nil
	case string:
		return fmt.Sprintf(`"%s"`, jsonObj), nil
	case bool:
		return strconv.FormatBool(jsonObj.(bool)), nil
	default:
		reflectVal := reflect.ValueOf(jsonObj)

		switch reflectVal.Kind() {
		case reflect.Ptr:
			return stringify(reflectVal.Elem().Interface())

		case reflect.Struct:
			reflectType := reflect.TypeOf(jsonObj)
			return stringifyStruct(reflectVal, reflectType)

		case reflect.Map:
			return stringifyMap(reflectVal)

		case reflect.Slice, reflect.Array:
			return stringifyArray(reflectVal)
		}

		return "", errors.New(fmt.Sprintf("not support this type:%s", reflect.TypeOf(jsonObj)))
		//panic(fmt.Sprintf("not support this type:%s", reflect.TypeOf(jsonObj)))
	}
}

func Stringify(jsonObj interface{}) (string, error) {
	return stringify(jsonObj)
}
