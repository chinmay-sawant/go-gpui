package window

import (
	"reflect"
	"strconv"
)

// devJSONLeaf renders one scalar value in its JSON form.
func devJSONLeaf(v reflect.Value, key, path string) *devJSONNode {
	switch v.Kind() {
	case reflect.Bool:
		return devJSONScalar(strconv.FormatBool(v.Bool()), devBoolInk, key, path)
	case reflect.String:
		return devJSONScalar(string(devJSONMarshal(v.String())), devStrInk, key, path)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return devJSONScalar(strconv.FormatInt(v.Int(), 10), devNumInk, key, path)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return devJSONScalar(strconv.FormatUint(v.Uint(), 10), devNumInk, key, path)
	case reflect.Float32, reflect.Float64:
		return devJSONScalar(string(devJSONMarshal(v.Interface())), devNumInk, key, path)
	}

	return devJSONScalar("null", devBoolInk, key, path)
}
