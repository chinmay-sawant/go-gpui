package transfer

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// PartialSuffix is appended to the destination name for the in-progress file.
const PartialSuffix = ".ownframe-part"

// PartialPath returns the partial-file path for a destination.
func PartialPath(dest string) string { return dest + PartialSuffix }

// reserved are Windows device names that cannot be used as file names.
var reserved = map[string]bool{
	"con": true, "prn": true, "aux": true, "nul": true,
	"com1": true, "com2": true, "com3": true, "com4": true,
	"com5": true, "com6": true, "com7": true, "com8": true, "com9": true,
	"lpt1": true, "lpt2": true, "lpt3": true, "lpt4": true,
	"lpt5": true, "lpt6": true, "lpt7": true, "lpt8": true, "lpt9": true,
}

// SafeName derives a bare file name from an untrusted URL path base. The
// result has no separators or control characters, is never a Windows device
// name, and never ends in a dot or space. It falls back to "download".
func SafeName(raw string) string {
	name := pathBase(raw)
	name = strings.Map(safeRune, name)
	name = strings.TrimRight(name, ". ")
	name = trimWindows(name)
	if name == "" || name == "." || name == ".." {
		name = "download"
	}

	return clipName(name)
}

// pathBase takes the last path element of a URL-ish string.
func pathBase(raw string) string {
	raw = strings.SplitN(raw, "?", 2)[0]
	raw = strings.SplitN(raw, "#", 2)[0]
	raw = strings.ReplaceAll(raw, "\\", "/")
	parts := strings.Split(raw, "/")
	for i := len(parts) - 1; i >= 0; i-- {
		if parts[i] != "" {
			return parts[i]
		}
	}

	return ""
}

// safeRune drops separators, control characters, and characters Windows
// forbids, and keeps Unicode letters as they are.
func safeRune(r rune) rune {
	if r < 0x20 || r == 0x7f {
		return -1
	}

	if strings.ContainsRune(`<>:"|?*/\`, r) {
		return '_'
	}

	return r
}

// trimWindows cuts the extension off a device name and rejects "con.txt".
func trimWindows(name string) string {
	base := name
	if i := strings.IndexByte(base, '.'); i >= 0 {
		base = base[:i]
	}

	if reserved[strings.ToLower(base)] {
		return "_" + name
	}

	return name
}

// clipName bounds a long name while keeping the extension.
func clipName(name string) string {
	const max = 120
	if len(name) <= max {
		return name
	}

	ext := filepath.Ext(name)
	if len(ext) > 16 {
		ext = ""
	}

	return name[:max-len(ext)] + ext
}

// UniqueDestination picks a never-existing path under dir for raw. It tries
// raw, then "name (1)", "name (2)", and so on, comparing case-folded names
// so a Windows directory cannot produce two colliding entries.
func UniqueDestination(dir, raw string) (string, error) {
	name := SafeName(raw)
	dir = filepath.Clean(dir)

	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}

	for i := 0; i < 1000; i++ {
		candidate := filepath.Join(dir, numbered(name, i))
		used, err := taken(dir, candidate)
		if err != nil {
			return "", err
		}

		if !used {
			return candidate, nil
		}
	}

	return "", fmt.Errorf("transfer: no free destination name for %q", raw)
}

// numbered inserts " (n)" before the extension. Index zero keeps the name.
func numbered(name string, n int) string {
	if n == 0 {
		return name
	}

	ext := filepath.Ext(name)

	return fmt.Sprintf("%s (%d)%s", name[:len(name)-len(ext)], n, ext)
}

// taken reports whether candidate or a case-folded twin exists in dir.
func taken(dir, candidate string) (bool, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false, err
	}

	want := strings.ToLower(filepath.Base(candidate))
	for _, entry := range entries {
		if strings.ToLower(entry.Name()) == want {
			return true, nil
		}
	}

	return false, nil
}

// HashFile returns the hex sha256 of path.
func HashFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	sum := sha256.New()
	if _, err := io.Copy(sum, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(sum.Sum(nil)), nil
}

// CheckChecksum verifies a "sha256:<hex>" checksum against path. An empty
// want always passes.
func CheckChecksum(path, want string) error {
	if want == "" {
		return nil
	}

	algo, hexWant, ok := strings.Cut(want, ":")
	if !ok || algo != "sha256" {
		return fmt.Errorf("%w: unsupported checksum %q", ErrChecksum, want)
	}

	got, err := HashFile(path)
	if err != nil {
		return err
	}

	if !strings.EqualFold(got, hexWant) {
		return fmt.Errorf("%w: sha256 %s != %s", ErrChecksum, got, hexWant)
	}

	return nil
}
