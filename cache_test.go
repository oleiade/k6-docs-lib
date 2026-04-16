package k6docslib

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/klauspost/compress/zstd"
	"go.k6.io/k6/lib/fsext"
)

func TestCacheDir(t *testing.T) {
	t.Parallel()

	t.Run("default path", func(t *testing.T) {
		t.Parallel()
		env := map[string]string{"HOME": "/home/user"}
		dir, err := CacheDir(env, "v1.7.x")
		if err != nil {
			t.Fatal(err)
		}
		want := filepath.Join("/home/user", ".local", "share", "k6", "docs", "v1.7.x")
		if dir != want {
			t.Errorf("got %q, want %q", dir, want)
		}
	})

	t.Run("override via K6_DOCS_CACHE_DIR", func(t *testing.T) {
		t.Parallel()
		env := map[string]string{"HOME": "/home/user", "K6_DOCS_CACHE_DIR": "/custom/cache"}
		dir, err := CacheDir(env, "v1.7.x")
		if err != nil {
			t.Fatal(err)
		}
		want := filepath.Join("/custom/cache", "v1.7.x")
		if dir != want {
			t.Errorf("got %q, want %q", dir, want)
		}
	})

	t.Run("no home dir", func(t *testing.T) {
		t.Parallel()
		env := map[string]string{}
		_, err := CacheDir(env, "v1.7.x")
		if err == nil {
			t.Fatal("expected error when HOME is not set")
		}
	})
}

func TestIsCached(t *testing.T) {
	t.Parallel()

	afs := fsext.NewMemMapFs()
	env := map[string]string{"HOME": "/home/user"}

	if IsCached(afs, env, "v1.7.x") {
		t.Error("expected not cached before creating dir")
	}

	dir, _ := CacheDir(env, "v1.7.x")
	if err := afs.MkdirAll(dir, 0o750); err != nil {
		t.Fatal(err)
	}

	if !IsCached(afs, env, "v1.7.x") {
		t.Error("expected cached after creating dir")
	}
}

func TestIsValidVersion(t *testing.T) {
	t.Parallel()

	tests := []struct {
		version string
		valid   bool
	}{
		{"v1.7.x", true},
		{"v0.55.x", true},
		{"1.5.0-rc.1", true},
		{"", false},
		{".", false},
		{"..", false},
		{"../escape", false},
		{"v1.7.x/../../etc/passwd", false},
	}

	for _, tt := range tests {
		t.Run(tt.version, func(t *testing.T) {
			t.Parallel()
			got := isValidVersion(tt.version)
			if got != tt.valid {
				t.Errorf("isValidVersion(%q) = %v, want %v", tt.version, got, tt.valid)
			}
		})
	}
}

func TestBundleURL(t *testing.T) {
	t.Parallel()

	t.Run("default URL", func(t *testing.T) {
		t.Parallel()
		url := bundleURL(map[string]string{}, "v1.7.x")
		want := "https://github.com/grafana/k6-docs-lib/releases/download/doc-bundles/docs-v1.7.x.tar.zst"
		if url != want {
			t.Errorf("got %q, want %q", url, want)
		}
	})

	t.Run("override URL", func(t *testing.T) {
		t.Parallel()
		env := map[string]string{"K6_DOCS_BUNDLE_URL": "http://example.com/bundle.tar.zst"}
		url := bundleURL(env, "v1.7.x")
		if url != "http://example.com/bundle.tar.zst" {
			t.Errorf("got %q, want override URL", url)
		}
	})
}

// buildTestBundle creates a tar.zst bundle in memory containing a sections.json
// and a markdown file for testing.
func buildTestBundle(t *testing.T) []byte {
	t.Helper()

	idx := Index{
		Version: "v1.7.x",
		Sections: []Section{
			{
				Slug:     "getting-started",
				RelPath:  "getting-started/_index.md",
				Title:    "Getting Started",
				Category: "getting-started",
				IsIndex:  true,
				Children: []string{},
			},
		},
	}
	sectionsJSON, err := json.Marshal(idx)
	if err != nil {
		t.Fatal(err)
	}

	var tarBuf bytes.Buffer
	tw := tar.NewWriter(&tarBuf)

	// sections.json
	if err := tw.WriteHeader(&tar.Header{
		Name: "sections.json",
		Size: int64(len(sectionsJSON)),
		Mode: 0o640,
	}); err != nil {
		t.Fatal(err)
	}
	if _, err := tw.Write(sectionsJSON); err != nil {
		t.Fatal(err)
	}

	// markdown dir
	if err := tw.WriteHeader(&tar.Header{
		Typeflag: tar.TypeDir,
		Name:     "markdown/",
		Mode:     0o750,
	}); err != nil {
		t.Fatal(err)
	}

	// markdown file
	mdContent := []byte("# Getting Started\nWelcome to k6.")
	if err := tw.WriteHeader(&tar.Header{
		Name: "markdown/getting-started/_index.md",
		Size: int64(len(mdContent)),
		Mode: 0o640,
	}); err != nil {
		t.Fatal(err)
	}
	if _, err := tw.Write(mdContent); err != nil {
		t.Fatal(err)
	}

	if err := tw.Close(); err != nil {
		t.Fatal(err)
	}

	// Compress with zstd.
	var zstBuf bytes.Buffer
	zw, err := zstd.NewWriter(&zstBuf)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := zw.Write(tarBuf.Bytes()); err != nil {
		t.Fatal(err)
	}
	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}

	return zstBuf.Bytes()
}

func TestEnsureDocsRoundTrip(t *testing.T) {
	t.Parallel()

	bundle := buildTestBundle(t)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("ETag", `"test-etag-1"`)
		_, _ = w.Write(bundle)
	}))
	defer srv.Close()

	afs := fsext.NewMemMapFs()
	env := map[string]string{
		"HOME":               "/home/user",
		"K6_DOCS_BUNDLE_URL": srv.URL + "/docs-v1.7.x.tar.zst",
	}

	ctx := context.Background()
	dir, err := EnsureDocs(ctx, afs, env, "v1.7.x", srv.Client())
	if err != nil {
		t.Fatalf("EnsureDocs: %v", err)
	}

	// Verify sections.json was extracted and is loadable.
	idx, err := LoadIndex(afs, dir)
	if err != nil {
		t.Fatalf("LoadIndex: %v", err)
	}
	if idx.Version != "v1.7.x" {
		t.Errorf("version = %q, want %q", idx.Version, "v1.7.x")
	}

	sec, ok := idx.Lookup("getting-started")
	if !ok {
		t.Fatal("expected to find getting-started section")
	}
	if sec.Title != "Getting Started" {
		t.Errorf("title = %q, want %q", sec.Title, "Getting Started")
	}

	// Verify markdown file was extracted.
	mdPath := filepath.Join(dir, "markdown", "getting-started", "_index.md")
	data, err := fsext.ReadFile(afs, mdPath)
	if err != nil {
		t.Fatalf("read markdown: %v", err)
	}
	if !bytes.Contains(data, []byte("Welcome to k6")) {
		t.Error("markdown content missing expected text")
	}

	// Second call should use cache (no download needed).
	dir2, err := EnsureDocs(ctx, afs, env, "v1.7.x", srv.Client())
	if err != nil {
		t.Fatalf("second EnsureDocs: %v", err)
	}
	if dir2 != dir {
		t.Errorf("second call returned different dir: %q vs %q", dir2, dir)
	}
}

func TestExtractPathTraversal(t *testing.T) {
	t.Parallel()

	// Build a tar with a path traversal entry.
	var tarBuf bytes.Buffer
	tw := tar.NewWriter(&tarBuf)
	if err := tw.WriteHeader(&tar.Header{
		Name: "../etc/passwd",
		Size: 4,
		Mode: 0o640,
	}); err != nil {
		t.Fatal(err)
	}
	if _, err := tw.Write([]byte("evil")); err != nil {
		t.Fatal(err)
	}
	if err := tw.Close(); err != nil {
		t.Fatal(err)
	}

	// Compress.
	var zstBuf bytes.Buffer
	zw, err := zstd.NewWriter(&zstBuf)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := zw.Write(tarBuf.Bytes()); err != nil {
		t.Fatal(err)
	}
	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}

	afs := fsext.NewMemMapFs()
	err = extract(afs, bytes.NewReader(zstBuf.Bytes()), "/cache")
	if err == nil {
		t.Fatal("expected error for path traversal")
	}
}
