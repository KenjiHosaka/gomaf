package gomaf

import (
	"errors"
	"reflect"
	"unicode"
)

const (
	tag = "gomaf"
)

func Diff(old, new any) (map[string]any, error) {
	result := make(map[string]any)

	oldVal := reflect.ValueOf(old)
	newVal := reflect.ValueOf(new)

	if oldVal.Type() != newVal.Type() {
		return nil, errors.New("different type")
	}

	for i := 0; i < oldVal.NumField(); i++ {
		oldFieldValue := oldVal.Field(i)
		newFieldValue := newVal.Field(i)

		key := oldVal.Type().Field(i).Tag.Get(tag)
		if key == "" {
			key = toSnakeCase(oldVal.Type().Field(i).Name)
		}

		if !reflect.DeepEqual(oldFieldValue.Interface(), newFieldValue.Interface()) {
			result[key] = newFieldValue.Interface()
		}
	}

	return result, nil
}

func toSnakeCase(str string) string {
	result := make([]rune, 0, len(str)*2)
	for i, char := range str {
		if unicode.IsUpper(char) {
			if i > 0 {
				result = append(result, '_')
			}
			char = unicode.ToLower(char)
		}
		result = append(result, char)
	}
	return string(result)
}
