package window

import (
	"encoding/json"
	"reflect"
	"strings"
)

// devJSONMarshal is json.Marshal with a null fallback for values it refuses.
func devJSONMarshal(v any) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		return []byte("null")
	}

	return b
}

// devJSONTag splits a json struct tag into its name and options.
func devJSONTag(tag string) (string, string) {
	if i := strings.IndexByte(tag, ','); i >= 0 {
		return tag[:i], tag[i+1:]
	}

	return tag, ""
}

// devJSONEmpty reports the values omitempty drops.
func devJSONEmpty(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Pointer:
		return v.IsNil()
	}

	return false
}

// devJSONJoin names a child node under its parent.
func devJSONJoin(path, key string) string {
	if path == "" {
		return key
	}

	return path + "." + key
}
