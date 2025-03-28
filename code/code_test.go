package code

import (
	"testing"
)

func TestMake(t *testing.T) {
	tests := []struct {
		op       OpCode
		operands []int
		expected []byte
	}{
		{OpConstant, []int{65534}, []byte{byte(OpConstant), 255, 254}},
	}

	for _, test := range tests {
		instruction := Make(test.op, test.operands...)

		if len(instruction) != len(test.expected) {
			t.Errorf("instruction has wrong length. want: %d, got: %d",
				len(test.expected), len(instruction))
		}

		for i, b := range instruction {
			if b != test.expected[i] {
				t.Errorf("instruction byte %d is wrong. want: %d, got: %d",
					i, test.expected[i], b)
			}
		}
	}
}
