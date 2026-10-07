//go:build windows

package collector

import "syscall"

// Windows API bindings. LazyDLL keeps the calls out of the startup path.

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")
	psapi    = syscall.NewLazyDLL("psapi.dll")

	procGetSystemTimes            = kernel32.NewProc("GetSystemTimes")
	procGlobalMemoryStatusEx      = kernel32.NewProc("GlobalMemoryStatusEx")
	procGetTickCount64            = kernel32.NewProc("GetTickCount64")
	procGetLogicalDrives          = kernel32.NewProc("GetLogicalDrives")
	procGetDriveTypeW             = kernel32.NewProc("GetDriveTypeW")
	procGetDiskFreeSpaceExW       = kernel32.NewProc("GetDiskFreeSpaceExW")
	procCreateToolhelp32Snapshot  = kernel32.NewProc("CreateToolhelp32Snapshot")
	procProcess32FirstW           = kernel32.NewProc("Process32FirstW")
	procProcess32NextW            = kernel32.NewProc("Process32NextW")
	procOpenProcess               = kernel32.NewProc("OpenProcess")
	procCloseHandle               = kernel32.NewProc("CloseHandle")
	procGetProcessTimes           = kernel32.NewProc("GetProcessTimes")
	procGetProcessIoCounters      = kernel32.NewProc("GetProcessIoCounters")
	procQueryFullProcessImageName = kernel32.NewProc("QueryFullProcessImageNameW")
	procGetProcessMemoryInfo      = psapi.NewProc("GetProcessMemoryInfo")
)

const (
	th32csSnapProcess               = 0x2
	processQueryLimitedInfo         = 0x1000
	driveRemovable                  = 2
	driveFixed                      = 3
	invalidHandle           uintptr = ^uintptr(0)
)
