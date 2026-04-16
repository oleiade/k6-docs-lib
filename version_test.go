package k6docslib

import (
	"runtime/debug"
	"testing"
)

func TestMapToWildcard(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected string
	}{
		{"v1.5.0", "v1.5.x"},
		{"v1.6.0", "v1.6.x"},
		{"v0.55.2-rc.1", "v0.55.x"},
		{"1.5.0", "v1.5.x"},
		{"v1.5.0-beta+build", "v1.5.x"},
		{"", ""},
		{"v1", "v1"},
		{"v1.5", "v1.5"},
		{"invalid", "invalid"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			got := MapToWildcard(tt.input)
			if got != tt.expected {
				t.Errorf("MapToWildcard(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestDetectK6Version(t *testing.T) {
	t.Parallel()

	t.Run("k6 dependency found", func(t *testing.T) {
		t.Parallel()
		readBuildInfo := func() (*debug.BuildInfo, bool) {
			return &debug.BuildInfo{
				Deps: []*debug.Module{
					{Path: "go.k6.io/k6", Version: "v1.6.0"},
				},
			}, true
		}
		got, err := detectK6Version(readBuildInfo)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != "v1.6.x" {
			t.Errorf("got %q, want %q", got, "v1.6.x")
		}
	})

	t.Run("k6 dependency not found", func(t *testing.T) {
		t.Parallel()
		readBuildInfo := func() (*debug.BuildInfo, bool) {
			return &debug.BuildInfo{
				Deps: []*debug.Module{
					{Path: "some/other/module", Version: "v1.0.0"},
				},
			}, true
		}
		_, err := detectK6Version(readBuildInfo)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("build info unavailable", func(t *testing.T) {
		t.Parallel()
		readBuildInfo := func() (*debug.BuildInfo, bool) {
			return nil, false
		}
		_, err := detectK6Version(readBuildInfo)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("k6 pre-release version", func(t *testing.T) {
		t.Parallel()
		readBuildInfo := func() (*debug.BuildInfo, bool) {
			return &debug.BuildInfo{
				Deps: []*debug.Module{
					{Path: "go.k6.io/k6", Version: "v0.55.2-rc.1"},
				},
			}, true
		}
		got, err := detectK6Version(readBuildInfo)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != "v0.55.x" {
			t.Errorf("got %q, want %q", got, "v0.55.x")
		}
	})
}
