//go:build linux

package collector

import "testing"

// diskFixture holds a whole disk, a partition, and a loop device.
const diskFixture = `   8       0 sda 100 0 2000 0 50 0 4000 0 0 0 0
   8       1 sda1 100 0 2000 0 50 0 4000 0 0 0 0
 259       0 nvme0n1 10 0 20 0 30 0 40 0 0 0 0
 259       1 nvme0n1p1 10 0 20 0 30 0 40 0 0 0 0
   7       0 loop0 1 0 2 0 3 0 4 0 0 0 0
`

// TestParseDiskStats checks the sector to byte conversion and the whole-disk
// filter.
func TestParseDiskStats(t *testing.T) {
	io := parseDiskStats([]byte(diskFixture))

	if len(io) != 2 {
		t.Fatalf("devices = %d: %+v", len(io), io)
	}
	if io["sda"].read != 2000*512 || io["sda"].write != 4000*512 {
		t.Fatalf("sda = %+v", io["sda"])
	}
	if io["nvme0n1"].read != 20*512 {
		t.Fatalf("nvme = %+v", io["nvme0n1"])
	}
	if _, ok := io["loop0"]; ok {
		t.Fatal("loop device was included")
	}
}

// TestWholeDiskName checks the name mapping.
func TestWholeDiskName(t *testing.T) {
	cases := map[string]string{
		"sda":       "sda",
		"sda1":      "sda",
		"nvme0n1":   "nvme0n1",
		"nvme0n1p2": "nvme0n1",
		"mmcblk0":   "mmcblk0",
		"mmcblk0p1": "mmcblk0",
		"loop0":     "",
		"dm-0":      "",
		"ram0":      "",
	}
	for in, want := range cases {
		if got := wholeDiskName(in); got != want {
			t.Fatalf("wholeDiskName(%q) = %q, want %q", in, got, want)
		}
	}
}
