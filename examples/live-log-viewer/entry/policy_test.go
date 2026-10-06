package entry

import "testing"

func TestDefaultPolicyValidates(t *testing.T) {
	if err := DefaultPolicy().Validate(); err != nil {
		t.Fatalf("default policy invalid: %v", err)
	}

	bad := []Policy{
		{},
		{BatchRecords: 1, BatchBytes: 1, MaxRecord: 1, MultilineLines: 1,
			MultilineBytes: 1, Overflow: Overflow(9)},
		{BatchRecords: 1, BatchBytes: 1, MaxRecord: 1, MultilineLines: 1,
			MultilineBytes: 1, Overflow: OverflowPause, MultilineHold: -1},
	}

	for i, p := range bad {
		if err := p.Validate(); err == nil {
			t.Errorf("policy %d validated", i)
		}
	}
}

func TestForStreamCountsLoss(t *testing.T) {
	p := DefaultPolicy()
	if p.Overflow != OverflowPause {
		t.Fatalf("default overflow = %v", p.Overflow)
	}

	s := p.ForStream()
	if s.Overflow != OverflowCountLoss {
		t.Fatalf("stream overflow = %v", s.Overflow)
	}

	if s.BatchRecords != p.BatchRecords {
		t.Fatal("ForStream changed the batch limits")
	}
}
