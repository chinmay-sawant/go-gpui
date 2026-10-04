package window

import (
	"reflect"
	"strconv"
)

// buildStruct renders exported fields in declaration order.
func (p *devJSONPrinter) buildStruct(v reflect.Value, key, path string) *devJSONNode {
	node := &devJSONNode{kind: 'o', key: key, path: path}
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.PkgPath != "" {
			continue
		}

		name, opts := devJSONTag(f.Tag.Get("json"))
		if name == "-" {
			continue
		}

		if name == "" {
			name = f.Name
		}

		fv := v.Field(i)
		if opts == "omitempty" && devJSONEmpty(fv) {
			continue
		}

		node.kids = append(node.kids, p.build(fv, name, devJSONJoin(path, name)))
	}

	return node
}

// buildArray renders elements in order, named by index.
func (p *devJSONPrinter) buildArray(v reflect.Value, key, path string) *devJSONNode {
	node := &devJSONNode{kind: 'a', key: key, path: path}

	for i := 0; i < v.Len(); i++ {
		ik := strconv.Itoa(i)
		node.kids = append(node.kids, p.build(v.Index(i), ik, devJSONJoin(path, ik)))
	}

	return node
}
