package closure

import (
	"testing"
)

func Test_NewIDGenerator(t *testing.T) {
	idGen := NewIDGenerator()
	if id := idGen(); id != 102 {
		t.Errorf("Expected 102, got %d",
			id)
	}
	if id := idGen(); id != 103 {
		t.Errorf("Expected 103, got %d",
			id)
	}
	if id := idGen(); id != 104 {
		t.Errorf("Expected 104, got %d",
			id)
	}
}

// Test concurrent calls to the ID generator
func Test_NewIDGenerator_Concurrent(
	t *testing.T) {
	idGen := NewIDGenerator()
	const N = 3
	results := make(chan int, N)
	for range N {
		go func() {
			results <- idGen()
		}()
	}
	for range N {
		<-results
	}
}
