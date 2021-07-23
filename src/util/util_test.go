package util

import "testing"

func TestStringContainsTrue(t *testing.T) {
	matching := StringContains("Hello, World!", "hello")

	if !matching {
		t.Errorf("StringContains failed, got: %t, want: %t", matching, true)
	}
}
