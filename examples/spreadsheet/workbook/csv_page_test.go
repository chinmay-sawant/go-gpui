package workbook

import (
	"strings"
	"testing"
)

func TestCSVPage(t *testing.T) {
	var b strings.Builder

	for i := 0; i < 120; i++ {
		b.WriteString("r")
		b.WriteString(itoa(i))
		b.WriteString("\n")
	}

	tab, err := ParseCSV([]byte(b.String()), DefaultCSVOptions())
	if err != nil {
		t.Fatal(err)
	}

	page := tab.Page(0, 0)
	if len(page.Rows) != 50 || page.Next != 50 || !page.More {
		t.Fatalf("page 1 = %+v", page)
	}

	page = tab.Page(page.Next, 50)
	if len(page.Rows) != 50 || page.Next != 100 || !page.More {
		t.Fatalf("page 2 = %+v", page)
	}

	page = tab.Page(page.Next, 50)
	if len(page.Rows) != 20 || page.More {
		t.Fatalf("page 3 = %+v", page)
	}
}
