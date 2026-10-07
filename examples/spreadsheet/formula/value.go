package formula

// Value is a computed cell value.
type Value struct {
	Kind   Kind
	Num    float64
	Str    string
	Code   string
	Detail string
}

// Number wraps a number.
func Number(n float64) Value { return Value{Kind: KindNumber, Num: n} }

// Text wraps text.
func Text(s string) Value { return Value{Kind: KindText, Str: s} }

// Err builds an error value.
func Err(code string) Value { return Value{Kind: KindError, Code: code} }

// ErrDetail builds an error value with a human-readable detail.
func ErrDetail(code, detail string) Value {
	return Value{Kind: KindError, Code: code, Detail: detail}
}

// IsError reports whether the value is an error.
func (v Value) IsError() bool { return v.Kind == KindError }

// IsBlank reports whether the value came from an empty cell.
func (v Value) IsBlank() bool { return v.Kind == KindBlank }

// Display renders the value the way a cell shows it.
func (v Value) Display() string {
	switch v.Kind {
	case KindNumber:
		return FormatNumber(v.Num)
	case KindText:
		return v.Str
	case KindError:
		return v.Code
	}

	return ""
}
