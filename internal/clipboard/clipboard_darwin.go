//go:build darwin && !ios && cgo

package clipboard

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework AppKit
#import <AppKit/AppKit.h>
#import <stdlib.h>

static void gpuiPut(const void *p, int n) {
	@autoreleasepool {
		NSPasteboard *pb = [NSPasteboard generalPasteboard];
		[pb clearContents];
		NSString *s = [[NSString alloc] initWithBytes:p
			length:(NSUInteger)n
			encoding:NSUTF8StringEncoding];
		if (s == nil) {
			[pb setString:@"" forType:NSPasteboardTypeString];
			return;
		}
		[pb setString:s forType:NSPasteboardTypeString];
		[s release];
	}
}

static char *gpuiGet(void) {
	@autoreleasepool {
		NSString *s = [[NSPasteboard generalPasteboard]
			stringForType:NSPasteboardTypeString];
		if (s == nil || [s UTF8String] == NULL) {
			return NULL;
		}
		return strdup([s UTF8String]);
	}
}
*/
import "C"
import "unsafe"

func writeOS(text string) {
	if text == "" {
		C.gpuiPut(nil, 0)

		return
	}

	p := unsafe.Pointer(unsafe.StringData(text))
	C.gpuiPut(p, C.int(len(text)))
}

func readOS() (string, bool) {
	p := C.gpuiGet()
	if p == nil {
		return "", false
	}

	defer C.free(unsafe.Pointer(p))

	return C.GoString(p), true
}
