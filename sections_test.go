package k6docslib

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"

	"go.k6.io/k6/lib/fsext"
)

func newTestIndex() *Index {
	idx := &Index{
		Version: "v1.7.x",
		Sections: []Section{
			{Slug: "using-k6", Title: "Using k6", Category: "using-k6", Weight: 10, IsIndex: true, Children: []string{"using-k6/scenarios"}},
			{Slug: "using-k6/scenarios", Title: "Scenarios", Category: "using-k6", Weight: 20, Children: []string{}},
			{Slug: "javascript-api", Title: "JavaScript API", Category: "javascript-api", Weight: 30, IsIndex: true, Children: []string{"javascript-api/k6-http"}},
			{Slug: "javascript-api/k6-http", Title: "k6/http", Category: "javascript-api", Weight: 40, IsIndex: true, Children: []string{"javascript-api/k6-http/get"}},
			{Slug: "javascript-api/k6-http/get", Title: "get(url, [params])", Description: "Make HTTP GET requests", Category: "javascript-api", Weight: 50, Children: []string{}},
		},
	}
	idx.bySlug = make(map[string]*Section, len(idx.Sections))
	for i := range idx.Sections {
		idx.bySlug[strings.ToLower(idx.Sections[i].Slug)] = &idx.Sections[i]
	}
	return idx
}

func TestLoadIndex(t *testing.T) {
	t.Parallel()

	idx := newTestIndex()
	data, err := json.Marshal(idx)
	if err != nil {
		t.Fatal(err)
	}

	afs := fsext.NewMemMapFs()
	dir := "/cache/v1.7.x"
	if err := afs.MkdirAll(dir, 0o750); err != nil {
		t.Fatal(err)
	}
	if err := fsext.WriteFile(afs, filepath.Join(dir, "sections.json"), data, 0o640); err != nil {
		t.Fatal(err)
	}

	loaded, err := LoadIndex(afs, dir)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.Version != "v1.7.x" {
		t.Errorf("version = %q, want %q", loaded.Version, "v1.7.x")
	}
	if len(loaded.Sections) != 5 {
		t.Errorf("sections count = %d, want 5", len(loaded.Sections))
	}
}

func TestLookup(t *testing.T) {
	t.Parallel()

	idx := newTestIndex()

	sec, ok := idx.Lookup("using-k6")
	if !ok {
		t.Fatal("expected to find using-k6")
	}
	if sec.Title != "Using k6" {
		t.Errorf("title = %q, want %q", sec.Title, "Using k6")
	}

	// Case-insensitive lookup.
	sec, ok = idx.Lookup("USING-K6")
	if !ok {
		t.Fatal("expected case-insensitive lookup to find USING-K6")
	}
	if sec.Slug != "using-k6" {
		t.Errorf("slug = %q, want %q", sec.Slug, "using-k6")
	}

	// Missing slug.
	_, ok = idx.Lookup("nonexistent")
	if ok {
		t.Error("expected nonexistent slug to not be found")
	}
}

func TestSearch(t *testing.T) {
	t.Parallel()

	idx := newTestIndex()

	results := idx.Search("scenarios", nil)
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if results[0].Slug != "using-k6/scenarios" {
		t.Errorf("slug = %q, want %q", results[0].Slug, "using-k6/scenarios")
	}

	// Fuzzy search via description.
	results = idx.Search("HTTP GET", nil)
	if len(results) != 1 {
		t.Fatalf("expected 1 result for 'HTTP GET', got %d", len(results))
	}
	if results[0].Slug != "javascript-api/k6-http/get" {
		t.Errorf("slug = %q, want %q", results[0].Slug, "javascript-api/k6-http/get")
	}

	// Empty term.
	results = idx.Search("", nil)
	if results != nil {
		t.Errorf("expected nil for empty term, got %v", results)
	}
}

func TestChildren(t *testing.T) {
	t.Parallel()

	idx := newTestIndex()

	children := idx.Children("javascript-api")
	if len(children) != 1 {
		t.Fatalf("expected 1 child, got %d", len(children))
	}
	if children[0].Slug != "javascript-api/k6-http" {
		t.Errorf("child slug = %q, want %q", children[0].Slug, "javascript-api/k6-http")
	}

	// Nonexistent slug.
	children = idx.Children("nonexistent")
	if children != nil {
		t.Errorf("expected nil for nonexistent slug, got %v", children)
	}
}

func TestTopLevel(t *testing.T) {
	t.Parallel()

	idx := newTestIndex()

	top := idx.TopLevel()
	if len(top) != 2 {
		t.Fatalf("expected 2 top-level sections, got %d", len(top))
	}
	if top[0].Slug != "using-k6" {
		t.Errorf("first top-level slug = %q, want %q", top[0].Slug, "using-k6")
	}
	if top[1].Slug != "javascript-api" {
		t.Errorf("second top-level slug = %q, want %q", top[1].Slug, "javascript-api")
	}
}
