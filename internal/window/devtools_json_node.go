package window

import (
	"encoding/json"
	"image/color"
	"reflect"
)

// devJSONNode is one value in the rendered document: an object, an array,
// or a scalar, plus the key and dotted path that name it.
type devJSONNode struct {
	kind   byte
	key    string
	path   string
	scalar string
	ink    color.RGBA
	kids   []*devJSONNode
}

// devJSONScalar builds a leaf node.
func devJSONScalar(text string, ink color.RGBA, key, path string) *devJSONNode {
	return &devJSONNode{kind: 's', key: key, path: path, scalar: text, ink: ink}
}

// devJSONNumber reads a json.Number through its string kind.
func devJSONNumber(v reflect.Value) (string, bool) {
	if v.Kind() != reflect.String {
		return "", false
	}

	n, ok := v.Interface().(json.Number)

	return n.String(), ok
}

// build turns one reflected value into a node. key is the member name and
// path the dotted location in the document.
func (p *devJSONPrinter) build(v reflect.Value, key, path string) *devJSONNode {
	if !v.IsValid() {
		return devJSONScalar("null", devBoolInk, key, path)
	}

	if n, ok := devJSONNumber(v); ok {
		return devJSONScalar(n, devNumInk, key, path)
	}

	switch v.Kind() {
	case reflect.Interface, reflect.Pointer:
		if v.IsNil() {
			return devJSONScalar("null", devBoolInk, key, path)
		}

		return p.build(v.Elem(), key, path)
	case reflect.Struct:
		return p.buildStruct(v, key, path)
	case reflect.Slice, reflect.Array:
		if v.Kind() == reflect.Slice && v.Type().Elem().Kind() == reflect.Uint8 {
			return devJSONScalar(string(devJSONMarshal(v.Interface())), devStrInk, key, path)
		}

		return p.buildArray(v, key, path)
	case reflect.Map:
		return p.buildMap(v, key, path)
	}

	return devJSONLeaf(v, key, path)
}
