//go:build windows

package reader

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows"
)

// pathIdentity names the file a path currently points at using the volume
// serial number and file index. The attributes-only handle is opened with
// every sharing flag so a rotating writer can rename or delete at any time.
func pathIdentity(path string) (string, error) {
	p, err := windows.UTF16PtrFromString(path)
	if err != nil {
		return "", err
	}

	h, err := windows.CreateFile(p, windows.FILE_READ_ATTRIBUTES,
		windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE|windows.FILE_SHARE_DELETE,
		nil, windows.OPEN_EXISTING, windows.FILE_FLAG_BACKUP_SEMANTICS, 0)
	if err != nil {
		return "", err
	}

	defer windows.CloseHandle(h)

	return fileID(h)
}

func handleIdentity(f *os.File) (string, error) {
	return fileID(windows.Handle(f.Fd()))
}

func fileID(h windows.Handle) (string, error) {
	var info windows.ByHandleFileInformation
	if err := windows.GetFileInformationByHandle(h, &info); err != nil {
		return "", err
	}

	return fmt.Sprintf("%08x-%04x%08x",
		info.VolumeSerialNumber, info.FileIndexHigh, info.FileIndexLow), nil
}

// openShared opens a file for reading and lets other processes write,
// rename, or delete it while the handle is open. That is what makes a tail
// survive rotation on Windows.
func openShared(path string) (*os.File, error) {
	p, err := windows.UTF16PtrFromString(path)
	if err != nil {
		return nil, err
	}

	h, err := windows.CreateFile(p, windows.GENERIC_READ,
		windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE|windows.FILE_SHARE_DELETE,
		nil, windows.OPEN_EXISTING,
		windows.FILE_ATTRIBUTE_NORMAL|windows.FILE_FLAG_SEQUENTIAL_SCAN, 0)
	if err != nil {
		return nil, err
	}

	return os.NewFile(uintptr(h), path), nil
}
