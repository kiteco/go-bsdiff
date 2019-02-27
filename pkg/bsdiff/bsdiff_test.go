package bsdiff

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestDiff(t *testing.T) {
	oldbs := []byte{0xFF, 0xFA, 0xB7, 0xDD}
	newbs := []byte{0xFF, 0xFA, 0x90, 0xB7, 0xDD, 0xFE}
	var diffbs []byte
	var err error
	if diffbs, err = Bytes(oldbs, newbs); err != nil {
		t.Fatal(err)
	}
	z := []byte{
		66, 83, 68, 73, 70, 70, 52, 48, 45, 0, 0, 0, 0, 0, 0, 0,
		37, 0, 0, 0, 0, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0,
		66, 90, 104, 57, 49, 65, 89, 38, 83, 89, 201, 157, 29, 51, 0, 0,
		6, 192, 64, 92, 0, 64, 0, 32, 0, 33, 140, 160, 96, 108, 226, 200,
		241, 71, 197, 220, 145, 78, 20, 36, 50, 103, 71, 76, 192, 66, 90, 104,
		57, 49, 65, 89, 38, 83, 89, 255, 72, 155, 130, 0, 0, 0, 192, 0,
		64, 0, 32, 0, 33, 24, 70, 194, 238, 72, 167, 10, 18, 31, 233, 19,
		112, 64, 66, 90, 104, 57, 49, 65, 89, 38, 83, 89, 221, 19, 191, 92, 0,
		0, 0, 0, 42, 192, 0, 0, 128, 0, 2, 0, 1, 32, 0,
	}
	if !bytes.Equal(diffbs[:len(z)], z) {
		t.Fatal(diffbs[:len(z)], "!=", z)
	}
}

func TestOfftout(t *testing.T) {
	buf := make([]byte, 8)
	offtout(9001, buf)
	n := binary.LittleEndian.Uint64(buf)
	if n != 9001 {
		t.Fatal(n, "!=", 9001)
	}
	//
	offtout(9002, buf)
	n = binary.LittleEndian.Uint64(buf)
	if n != 9002 {
		t.Fatal(n, "!=", 9002)
	}
}
