package k6docslib

import (
	"testing"
)

func newTestMultiIndex() *MultiIndex {
	mi := NewMultiIndex()

	idx17 := newTestIndex() // v1.7.x from sections_test.go
	mi.Add("v1.7.x", idx17)

	idx16 := &Index{
		Version: "v1.6.x",
		Sections: []Section{
			{Slug: "using-k6", Title: "Using k6", Category: "using-k6", Weight: 10, IsIndex: true, Children: []string{}},
			{Slug: "results-output", Title: "Results Output", Category: "results-output", Weight: 20, IsIndex: true, Children: []string{}},
		},
	}
	idx16.bySlug = make(map[string]*Section, len(idx16.Sections))
	for i := range idx16.Sections {
		idx16.bySlug[idx16.Sections[i].Slug] = &idx16.Sections[i]
	}
	mi.Add("v1.6.x", idx16)

	return mi
}

func TestMultiIndexAdd(t *testing.T) {
	t.Parallel()

	mi := newTestMultiIndex()
	versions := mi.GetVersions()
	if len(versions) != 2 {
		t.Fatalf("expected 2 versions, got %d", len(versions))
	}
	if versions[0] != "v1.7.x" || versions[1] != "v1.6.x" {
		t.Errorf("versions = %v", versions)
	}
}

func TestMultiIndexLatest(t *testing.T) {
	t.Parallel()

	mi := newTestMultiIndex()
	if mi.GetLatestVersion() != "v1.7.x" {
		t.Errorf("latest = %q, want v1.7.x", mi.GetLatestVersion())
	}

	mi.SetLatest("v1.6.x")
	if mi.GetLatestVersion() != "v1.6.x" {
		t.Errorf("latest after SetLatest = %q, want v1.6.x", mi.GetLatestVersion())
	}
}

func TestMultiIndexLookup(t *testing.T) {
	t.Parallel()

	mi := newTestMultiIndex()

	// Lookup in v1.7.x.
	sec, ok := mi.Lookup("using-k6/scenarios", "v1.7.x")
	if !ok {
		t.Fatal("expected to find using-k6/scenarios in v1.7.x")
	}
	if sec.Title != "Scenarios" {
		t.Errorf("title = %q", sec.Title)
	}

	// Same slug not in v1.6.x.
	_, ok = mi.Lookup("using-k6/scenarios", "v1.6.x")
	if ok {
		t.Error("using-k6/scenarios should not exist in v1.6.x")
	}

	// Empty version resolves to latest.
	sec, ok = mi.Lookup("using-k6/scenarios", "")
	if !ok {
		t.Fatal("expected lookup with empty version to use latest")
	}
	if sec.Title != "Scenarios" {
		t.Errorf("title = %q", sec.Title)
	}

	// Case-insensitive.
	_, ok = mi.Lookup("USING-K6", "v1.7.x")
	if !ok {
		t.Error("expected case-insensitive lookup")
	}
}

func TestMultiIndexGetAll(t *testing.T) {
	t.Parallel()

	mi := newTestMultiIndex()
	all := mi.GetAll("v1.6.x")
	if len(all) != 2 {
		t.Errorf("expected 2 sections in v1.6.x, got %d", len(all))
	}

	// Unknown version.
	all = mi.GetAll("v1.0.x")
	if all != nil {
		t.Errorf("expected nil for unknown version, got %v", all)
	}
}

func TestMultiIndexGetByCategory(t *testing.T) {
	t.Parallel()

	mi := newTestMultiIndex()
	sections := mi.GetByCategory("javascript-api", "v1.7.x")
	if len(sections) != 3 {
		t.Errorf("expected 3 javascript-api sections, got %d", len(sections))
	}
}

func TestMultiIndexGetCategories(t *testing.T) {
	t.Parallel()

	mi := newTestMultiIndex()
	cats := mi.GetCategories("v1.6.x")
	if len(cats) != 2 {
		t.Fatalf("expected 2 categories, got %d: %v", len(cats), cats)
	}
}

func TestMultiIndexMatchVersion(t *testing.T) {
	t.Parallel()

	mi := newTestMultiIndex()

	t.Run("direct match", func(t *testing.T) {
		t.Parallel()
		v, err := mi.MatchVersion("v1.7.x")
		if err != nil {
			t.Fatal(err)
		}
		if v != "v1.7.x" {
			t.Errorf("got %q", v)
		}
	})

	t.Run("semantic to wildcard", func(t *testing.T) {
		t.Parallel()
		v, err := mi.MatchVersion("v1.6.0")
		if err != nil {
			t.Fatal(err)
		}
		if v != "v1.6.x" {
			t.Errorf("got %q", v)
		}
	})

	t.Run("no match returns latest", func(t *testing.T) {
		t.Parallel()
		v, err := mi.MatchVersion("v9.9.9")
		if err == nil {
			t.Fatal("expected error for unmatched version")
		}
		if v != "v1.7.x" {
			t.Errorf("got %q, want latest", v)
		}
	})

	t.Run("empty returns latest", func(t *testing.T) {
		t.Parallel()
		v, err := mi.MatchVersion("")
		if err != nil {
			t.Fatal(err)
		}
		if v != "v1.7.x" {
			t.Errorf("got %q", v)
		}
	})
}

func TestMultiIndexHasVersion(t *testing.T) {
	t.Parallel()

	mi := newTestMultiIndex()
	if !mi.HasVersion("v1.7.x") {
		t.Error("expected HasVersion true for v1.7.x")
	}
	if mi.HasVersion("v1.0.x") {
		t.Error("expected HasVersion false for v1.0.x")
	}
}
