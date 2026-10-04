package window

import (
	"fmt"
	"reflect"
	"sort"
)

// devJSONPair is one map entry waiting for its sorted place.
type devJSONPair struct {
	label string
	value reflect.Value
}

// buildMap renders entries sorted by their member name.
func (p *devJSONPrinter) buildMap(v reflect.Value, key, path string) *devJSONNode {
	node := &devJSONNode{kind: 'o', key: key, path: path}
	pairs := make([]devJSONPair, 0, v.Len())

	for _, k := range v.MapKeys() {
		pairs = append(pairs, devJSONPair{label: devJSONMapKey(k), value: v.MapIndex(k)})
	}

	sort.Slice(pairs, func(i, j int) bool { return pairs[i].label < pairs[j].label })

	for _, pair := range pairs {
		node.kids = append(node.kids, p.build(pair.value, pair.label, devJSONJoin(path, pair.label)))
	}

	return node
}

// devJSONMapKey is the member name for one map key.
func devJSONMapKey(k reflect.Value) string {
	if k.Kind() == reflect.String {
		return k.String()
	}

	return fmt.Sprint(k.Interface())
}
