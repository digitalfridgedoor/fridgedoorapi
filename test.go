package fridgedoorapi

import (
	"fmt"
	"testing"
)

// ConnectOrSkip attempts to connect to mongodb, or skips the test
func ConnectOrSkip(t *testing.T) {
	connected := Connect()
	if !connected {
		fmt.Println("Cannot connect, skipping test")
		t.SkipNow()
	}
}
