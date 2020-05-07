package meals

import (
	"testing"
)

func TestMealify(t *testing.T) {
	var cookbook = []struct {
		in       string
		expected string
		err      error
	}{
		{"🍞🍅🧀", "🍕", nil},
		{"🍞🧀🍅", "🍕", nil},
		{"🧀🍞🍅", "🍕", nil},
		{"🧀🍅🍞", "🍕", nil},
		{"🍅🍞🧀", "🍕", nil},
		{"🍅🧀🍞", "🍕", nil},
		{"💧🍈", "🍉", nil},
		{"🍈💧", "🍉", nil},
		{"🍈🍳💧", "🍉", ErrNotFound},
	}

	for _, tt := range cookbook {
		t.Run(tt.in, func(t *testing.T) {
			if out, err := Mealify(tt.in); out != tt.expected && err != tt.err {
				t.Fatalf("expected input %s to make %s. Got \"%s\" instead\n", tt.in, tt.expected, out)
			}
		})
	}
}
