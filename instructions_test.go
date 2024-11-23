package polkavm

import (
	"fmt"
	"testing"
)

func TestDecodeImmediate(t *testing.T) {
	tests := []struct {
		Input  []byte
		Output uint32
	}{
		{
			Input:  []byte{},
			Output: 0,
		},
		{
			Input:  []byte{0},
			Output: 0,
		},
		{
			Input:  []byte{0x07},
			Output: 0x07,
		},
		{
			Input:  []byte{0xFF},
			Output: 0xFFFFFFFF,
		},
		{
			Input:  []byte{0xF0},
			Output: 0xFFFFFFF0,
		},
		{
			Input:  []byte{0x23, 0x7F},
			Output: 0x7F23,
		},
		{
			Input:  []byte{0x12, 0x80},
			Output: 0xFFFF8012,
		},
		{
			Input:  []byte{0x34, 0x12, 0x7F},
			Output: 0x7F1234,
		},
		{
			Input:  []byte{0x34, 0x12, 0x80},
			Output: 0xFF801234,
		},
		{
			Input:  []byte{0x12, 0x34, 0x56, 0x78},
			Output: 0x78563412,
		},
		{
			Input:  []byte{0x12, 0x34, 0x56, 0xFA},
			Output: 0xFA563412,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%x", test.Output), func(t *testing.T) {
			o := decodeImmediate(test.Input)
			if o != test.Output {
				t.Fatalf("expected %x actual %x", test.Output, o)
			}
		})
	}
}
