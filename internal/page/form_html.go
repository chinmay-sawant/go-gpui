package page

const formCSS = `<style>[data-gpui-field]{display:inline-block;border:1px solid #c8c2b4;padding:4px 6px;min-width:10em;min-height:1.6em;white-space:pre-wrap}[data-gpui-field][data-gpui-focus="1"]{border:2px solid #1a56db;padding:3px 5px}[data-gpui-field][data-gpui-selected="1"]{background:#d6e2ff}[data-gpui-selection]{background:#d6e2ff}[data-gpui-field][data-gpui-placeholder="1"]{color:#6b7280}[data-gpui-caret]{display:inline-block;width:1px;height:1em;background:#1c1915;margin-right:-1px}</style>`

// rewriteControls copies source, swapping each kept control for paintable
// HTML. It paints a select-all when selected is true and puts the caret at
// the end of the value. The page uses rewriteControlsState, which carries
// the caret and the range.
func rewriteControls(source string, spans []controlSpan, live map[string]Control, focusID string, selected bool) string {
	st := caretState{caret: -1}
	if focusID != "" {
		st.focus = focusID
		st.all = selected
	}

	return rewriteControlsState(source, spans, live, st)
}
