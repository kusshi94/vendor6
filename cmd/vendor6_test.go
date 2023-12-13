package cmd_test

import (
	"fmt"
	"testing"
)

func TestVendor6Command(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    string
	}{
		{
			description: "",
			input: `fe80::9e50:ffc3:85b6:be65
			fe80::1a51:6f53:c20f:e2fb
			fe80::809c:cf3b:25a0:c3b4
			fe80::a837:14a4:1069:7ae4
			fe80::6800:b8a1:dc84:4b78
			fe80::d0da:cb6e:5125:ddff
			fe80::5474:3fa5:9fca:99f3
			fe80::6690:fb06:8824:fbf1`,
			expected: ``,
		},
	}
	for _, tt := range tests {
		fmt.Println(tt.description)
	}
}
