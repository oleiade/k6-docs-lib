package k6docslib

import (
	"strings"
	"testing"
)

func TestTransform(t *testing.T) {
	t.Parallel()

	t.Run("strips code tags", func(t *testing.T) {
		t.Parallel()
		input := "before\n{{< code >}}\ncontent\n{{< /code >}}\nafter"
		got := Transform(input, "v1.7.x")
		if strings.Contains(got, "{{< code >}}") {
			t.Error("code tags not stripped")
		}
		if !strings.Contains(got, "content") {
			t.Error("content between code tags should be preserved")
		}
	})

	t.Run("converts admonitions to blockquotes", func(t *testing.T) {
		t.Parallel()
		input := "{{< admonition type=\"warning\" >}}\nBe careful here.\n{{< /admonition >}}"
		got := Transform(input, "v1.7.x")
		if !strings.Contains(got, "> **Warning:** Be careful here.") {
			t.Errorf("admonition not converted, got: %q", got)
		}
	})

	t.Run("strips section tags", func(t *testing.T) {
		t.Parallel()
		input := "{{< section >}}\ncontent\n{{< /section >}}"
		got := Transform(input, "v1.7.x")
		if strings.Contains(got, "{{< section >}}") {
			t.Error("section tags not stripped")
		}
	})

	t.Run("strips remaining shortcodes", func(t *testing.T) {
		t.Parallel()
		input := "{{< custom-shortcode param=\"val\" >}}"
		got := Transform(input, "v1.7.x")
		if strings.Contains(got, "{{<") {
			t.Errorf("shortcode not stripped, got: %q", got)
		}
	})

	t.Run("strips component tags", func(t *testing.T) {
		t.Parallel()
		input := "<Glossary>term</Glossary>"
		got := Transform(input, "v1.7.x")
		if strings.Contains(got, "<Glossary>") {
			t.Error("component tags not stripped")
		}
		if !strings.Contains(got, "term") {
			t.Error("content inside component should be preserved")
		}
	})

	t.Run("strips br tags", func(t *testing.T) {
		t.Parallel()
		input := "line1<br/>line2<br />line3"
		got := Transform(input, "v1.7.x")
		if strings.Contains(got, "<br") {
			t.Errorf("br tags not stripped, got: %q", got)
		}
	})

	t.Run("replaces K6_VERSION", func(t *testing.T) {
		t.Parallel()
		input := "Running k6 <K6_VERSION>."
		got := Transform(input, "v1.7.x")
		if !strings.Contains(got, "v1.7.x") {
			t.Errorf("version not replaced, got: %q", got)
		}
	})

	t.Run("strips image links keeping alt text", func(t *testing.T) {
		t.Parallel()
		input := "![screenshot](http://example.com/img.png)"
		got := Transform(input, "v1.7.x")
		if strings.Contains(got, "http://") {
			t.Error("image URL not stripped")
		}
		if !strings.Contains(got, "screenshot") {
			t.Error("alt text should be preserved")
		}
	})

	t.Run("strips markdown links keeping text", func(t *testing.T) {
		t.Parallel()
		input := "See [the docs](https://example.com) for details."
		got := Transform(input, "v1.7.x")
		if strings.Contains(got, "https://") {
			t.Error("URL not stripped")
		}
		if !strings.Contains(got, "the docs") {
			t.Error("link text should be preserved")
		}
	})

	t.Run("strips HTML comments", func(t *testing.T) {
		t.Parallel()
		input := "before<!-- hidden -->after"
		got := Transform(input, "v1.7.x")
		if strings.Contains(got, "hidden") {
			t.Error("HTML comment not stripped")
		}
	})

	t.Run("strips frontmatter", func(t *testing.T) {
		t.Parallel()
		input := "---\ntitle: Test\n---\n# Hello"
		got := Transform(input, "v1.7.x")
		if strings.Contains(got, "title: Test") {
			t.Error("frontmatter not stripped")
		}
		if !strings.Contains(got, "# Hello") {
			t.Error("body content should be preserved")
		}
	})

	t.Run("normalizes whitespace", func(t *testing.T) {
		t.Parallel()
		input := "a\n\n\n\n\nb"
		got := Transform(input, "v1.7.x")
		if strings.Contains(got, "\n\n\n") {
			t.Error("excessive newlines not collapsed")
		}
	})

	t.Run("empty string returns empty", func(t *testing.T) {
		t.Parallel()
		if got := Transform("", "v1.7.x"); got != "" {
			t.Errorf("expected empty, got %q", got)
		}
	})
}

func TestSplitFrontmatter(t *testing.T) {
	t.Parallel()

	t.Run("valid frontmatter", func(t *testing.T) {
		t.Parallel()
		input := "---\ntitle: Test\nweight: 1\n---\nBody text"
		yaml, body, ok := SplitFrontmatter(input)
		if !ok {
			t.Fatal("expected ok=true")
		}
		if yaml != "title: Test\nweight: 1" {
			t.Errorf("yaml = %q", yaml)
		}
		if body != "Body text" {
			t.Errorf("body = %q", body)
		}
	})

	t.Run("no frontmatter", func(t *testing.T) {
		t.Parallel()
		input := "Just some text"
		_, body, ok := SplitFrontmatter(input)
		if ok {
			t.Fatal("expected ok=false")
		}
		if body != input {
			t.Errorf("body = %q, want original", body)
		}
	})
}

func TestPrepareTransform(t *testing.T) {
	t.Parallel()

	t.Run("resolves shared shortcodes", func(t *testing.T) {
		t.Parallel()
		content := `Before {{< docs/shared source="k6" lookup="shared/intro.md" >}} after.`
		shared := map[string]string{
			"shared/intro.md": "---\ntitle: Intro\n---\nShared content here.",
		}
		got := PrepareTransform(content, shared)
		if !strings.Contains(got, "Shared content here.") {
			t.Errorf("shared content not resolved, got: %q", got)
		}
		if strings.Contains(got, "{{<") {
			t.Error("shortcode not replaced")
		}
	})

	t.Run("empty content returns empty", func(t *testing.T) {
		t.Parallel()
		if got := PrepareTransform("", nil); got != "" {
			t.Errorf("expected empty, got %q", got)
		}
	})

	t.Run("missing shared key returns empty replacement", func(t *testing.T) {
		t.Parallel()
		content := `{{< docs/shared source="k6" lookup="missing.md" >}}`
		got := PrepareTransform(content, map[string]string{})
		if strings.Contains(got, "{{<") {
			t.Error("shortcode should be replaced even when key is missing")
		}
	})
}
