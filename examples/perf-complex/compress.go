package main

import (
	"compress/gzip"
	"encoding/json"
	"os"
)

func writeLayout(path string, v any) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	z := gzip.NewWriter(f)
	err = json.NewEncoder(z).Encode(v)
	ze := z.Close()
	fe := f.Close()
	if err != nil {
		return err
	}
	if ze != nil {
		return ze
	}
	return fe
}
