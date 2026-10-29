package guest

import "testing"

func TestInfoZeroValue(t *testing.T) {
	var info Info
	if info.OS != "" {
		t.Fatal("unexpected OS")
	}
}
