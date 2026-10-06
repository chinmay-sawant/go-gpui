package ui

import "testing"

// BenchmarkStressPage measures one refresh and page slice over the 10,000
// process fixture, which is the worst case the table renders.
func BenchmarkStressPage(b *testing.B) {
	tb := newTable()
	tb.offer(snapshotOf(manyProcs(10000)...))

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tb.refresh()
		_ = tb.pageView()
	}
}
