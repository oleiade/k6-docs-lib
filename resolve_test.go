package k6docslib

import (
	"testing"
)

func TestNormalizeArgs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{"splits slashes", []string{"http/get"}, []string{"http", "get"}},
		{"multiple args", []string{"http", "get"}, []string{"http", "get"}},
		{"mixed", []string{"mod/child", "extra"}, []string{"mod", "child", "extra"}},
		{"empty parts filtered", []string{"a//b"}, []string{"a", "b"}},
		{"empty input", []string{}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NormalizeArgs(tt.input)
			if len(got) != len(tt.expected) {
				t.Fatalf("len = %d, want %d; got %v", len(got), len(tt.expected), got)
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("got[%d] = %q, want %q", i, got[i], tt.expected[i])
				}
			}
		})
	}
}

func TestResolveWithLookup(t *testing.T) {
	t.Parallel()

	known := map[string]bool{
		"using-k6":                          true,
		"using-k6/scenarios":                true,
		"javascript-api":                    true,
		"javascript-api/k6-http":            true,
		"javascript-api/k6-http/get":        true,
		"javascript-api/k6-http/k6-http-ok": true,
	}
	exists := func(slug string) bool { return known[slug] }

	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{"direct category match", []string{"using-k6", "scenarios"}, "using-k6/scenarios"},
		{"direct slug match", []string{"using-k6"}, "using-k6"},
		{"js api shorthand", []string{"http", "get"}, "javascript-api/k6-http/get"},
		{"full js api slug", []string{"javascript-api/k6-http/get"}, "javascript-api/k6-http/get"},
		{"k6 prefix fallback", []string{"http"}, "javascript-api/k6-http"},
		{"parent prefix fallback", []string{"javascript-api/k6-http/ok"}, "javascript-api/k6-http/k6-http-ok"},
		{"empty args", []string{}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := ResolveWithLookup(tt.args, exists)
			if got != tt.expected {
				t.Errorf("ResolveWithLookup(%v) = %q, want %q", tt.args, got, tt.expected)
			}
		})
	}
}

func TestResolveWithLookup_NilExists(t *testing.T) {
	t.Parallel()

	// Without a lookup function, JS API modules get k6- prefix by default.
	got := ResolveWithLookup([]string{"http", "get"}, nil)
	if got != "javascript-api/k6-http/get" {
		t.Errorf("got %q, want %q", got, "javascript-api/k6-http/get")
	}
}
