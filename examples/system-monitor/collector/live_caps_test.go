//go:build linux

package collector

import "testing"

// TestProcSourceCapabilities checks the adapter's advertised surface.
func TestProcSourceCapabilities(t *testing.T) {
	s := newProcSource()
	caps := s.Capabilities()

	if !caps.PerCoreCPU || !caps.Processes || !caps.ProcessDetail {
		t.Fatalf("caps = %+v", caps)
	}
	if caps.Handles {
		t.Fatal("handles should be unavailable on Linux")
	}
	if s.Name() != "procfs" {
		t.Fatalf("name = %q", s.Name())
	}
}
