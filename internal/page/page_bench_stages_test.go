package page_test

import (
	"bytes"
	"context"
	"html/template"
	"testing"

	"github.com/chinmay-sawant/gowkhtmltopdf/css"
	"github.com/chinmay-sawant/gowkhtmltopdf/html"
	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// BenchmarkRedrawStages breaks Redraw into its pipeline stages so the cost of
// each is readable on its own.
func BenchmarkRedrawStages(b *testing.B) {
	ctx := context.Background()
	tpl := template.Must(template.New("bench").Parse(benchHTML))

	execute := func() string {
		var body bytes.Buffer
		if err := tpl.Execute(&body, nil); err != nil {
			b.Fatal(err)
		}

		return body.String()
	}

	parse := func(source string) *html.Document {
		doc, err := html.Parse([]byte(source))
		if err != nil {
			b.Fatal(err)
		}

		return doc
	}

	apply := func(doc *html.Document) *css.Document {
		styled, err := css.Apply(ctx, doc, css.Options{
			WidthPx:  benchWidth,
			HeightPx: benchHeight,
			Media:    "screen",
		})
		if err != nil {
			b.Fatal(err)
		}

		return styled
	}

	b.Run("TemplateExecute", func(b *testing.B) {
		for b.Loop() {
			execute()
		}
	})

	b.Run("HTMLParse", func(b *testing.B) {
		source := execute()

		for b.Loop() {
			parse(source)
		}
	})

	b.Run("CSSApply", func(b *testing.B) {
		doc := parse(execute())

		for b.Loop() {
			apply(doc)
		}
	})

	b.Run("DisplayList", func(b *testing.B) {
		styled := apply(parse(execute()))

		for b.Loop() {
			if _, err := layout.DisplayListOptions(ctx, styled, layout.Options{}); err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("Lay", func(b *testing.B) {
		styled := apply(parse(execute()))

		for b.Loop() {
			if _, err := layout.LayOptions(ctx, styled, layout.Options{}); err != nil {
				b.Fatal(err)
			}
		}
	})
}
