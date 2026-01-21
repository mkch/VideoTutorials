package closure

import "testing"

func Test_NewIDGenerator2(t *testing.T) {
	idGen := NewIDGenerator2()
	if id := idGen.Call(); id != 102 {
		t.Errorf("Expected 102, got %d", id)
	}
	if id := idGen.Call(); id != 103 {
		t.Errorf("Expected 103, got %d", id)
	}
	if id := idGen.Call(); id != 104 {
		t.Errorf("Expected 104, got %d", id)
	}
}
