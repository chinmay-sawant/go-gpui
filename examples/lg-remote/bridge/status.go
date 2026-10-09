package bridge

// BluetoothStatus includes a revision so repeated results reach the page.
func BluetoothStatus() (string, uint64) {
	mu.Lock()
	defer mu.Unlock()
	return bt, btRevision
}
