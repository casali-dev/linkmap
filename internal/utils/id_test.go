package utils

import (
	"testing"
)

func TestGenerateID(t *testing.T) {
	t.Run("returns no error", func(t *testing.T) {
		_, err := GenerateID(6)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
	})

	t.Run("returns string of correct length", func(t *testing.T) {
		id, _ := GenerateID(6)
		if len(id) != 6 {
			t.Errorf("expected length 6, got: %d", len(id))
		}
	})

	t.Run("generates different values", func(t *testing.T) {
		id1, _ := GenerateID(6)
		id2, _ := GenerateID(6)
		if id1 == id2 {
			t.Error("expected different IDs")
		}
	})
}
