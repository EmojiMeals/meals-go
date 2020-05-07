package meals

import (
	"testing"
)

func TestMealify(t *testing.T) {
	cookbook := NewCookbook("recipes.json")
	var tests = []struct {
		in       []string
		expected string
		err      error
	}{
		{[]string{"🍞", "🍅", "🧀"}, "🍕", nil},
		{[]string{"🍞", "🧀", "🍅"}, "🍕", nil},
		{[]string{"💧", "🍈"}, "🍉", nil},
		{[]string{"🍈", "💧"}, "🍉", nil},
		{[]string{"🍈", "🍳", "💧"}, "", ErrNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if out, err := cookbook.Mealify(tt.in...); out != tt.expected || err != tt.err {
				t.Fatalf("expected input %s to make %s with err \"%s\". Instead got \"%s\" with err \"%s\"\n", tt.in, tt.expected, tt.err, out, err)
			}
		})
	}
}
