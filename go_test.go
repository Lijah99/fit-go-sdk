package fitgosdk

import "testing"

func TestGivesTrue(t *testing.T) {
	if !givesTrue() {
		t.Error("Expected givesTrue to return true, but it returned false")
	}
}
