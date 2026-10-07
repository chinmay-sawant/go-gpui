//go:build linux

package collector

import "testing"

// TestParseMounts checks the filesystem filter and octal escape decoding.
func TestParseMounts(t *testing.T) {
	data := []byte(`/dev/sda1 / ext4 rw,relatime 0 0
tmpfs /tmp tmpfs rw 0 0
/dev/nvme0n1p2 /home\040dir ext4 rw 0 0
proc /proc proc rw 0 0
`)

	mounts := parseMounts(data)
	if len(mounts) != 2 {
		t.Fatalf("mounts = %+v", mounts)
	}
	if mounts[1].mount != "/home dir" {
		t.Fatalf("escaped mount = %q", mounts[1].mount)
	}
}

// TestParseNetDev checks the interface counter positions and the header skip.
func TestParseNetDev(t *testing.T) {
	data := []byte(`Inter-|   Receive                                                |  Transmit
 face |bytes    packets errs drop fifo frame compressed multicast|bytes    packets errs drop fifo colls carrier compressed
    lo: 1000 10 0 0 0 0 0 0 1000 10 0 0 0 0 0 0
  eth0: 5000 50 1 0 0 0 0 0 6000 60 2 0 0 0 0 0
`)

	nets := parseNetDev(data)
	if len(nets) != 2 {
		t.Fatalf("interfaces = %+v", nets)
	}
	if nets[1].Name != "eth0" || nets[1].RXBytes != 5000 || nets[1].TXBytes != 6000 {
		t.Fatalf("eth0 = %+v", nets[1])
	}
	if nets[1].RXErrors != 1 || nets[1].TXErrors != 2 {
		t.Fatalf("eth0 errors = %+v", nets[1])
	}
}
