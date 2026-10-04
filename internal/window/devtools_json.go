package window

import "reflect"

// devJSONRows renders v as pretty-printed JSON rows with syntax colours.
// Structs keep declaration order, maps render keys sorted, slices render in
// order, and pointers or interfaces render their contents. An object or
// array opens on a row carrying devActJSON and the node's dotted key path,
// so a click can flip collapsed[key] to fold the node. A folded node renders
// as {...} or [...] on one row.
func devJSONRows(v any, collapsed map[string]bool) []devRow {
	p := &devJSONPrinter{collapsed: collapsed}
	p.emit(p.build(reflect.ValueOf(v), "", ""), 0, devLine{}, "", true)

	return p.rows
}
