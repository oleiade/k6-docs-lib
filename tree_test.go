package k6docslib

import (
	"testing"
)

func testSectionsForTree() []Section {
	return []Section{
		{Slug: "using-k6", Title: "Using k6", Category: "using-k6", Weight: 10, IsIndex: true},
		{Slug: "using-k6/scenarios", Title: "Scenarios", Category: "using-k6", Weight: 20, IsIndex: true},
		{Slug: "using-k6/scenarios/open-model", Title: "Open Model", Category: "using-k6", Weight: 30},
		{Slug: "using-k6/scenarios/closed-model", Title: "Closed Model", Category: "using-k6", Weight: 40},
		{Slug: "javascript-api", Title: "JavaScript API", Category: "javascript-api", Weight: 50, IsIndex: true},
	}
}

func TestBuildSectionTree(t *testing.T) {
	t.Parallel()

	sections := testSectionsForTree()

	t.Run("top level depth 1", func(t *testing.T) {
		t.Parallel()
		nodes, err := BuildSectionTree(sections, "", 1)
		if err != nil {
			t.Fatal(err)
		}
		if len(nodes) != 2 {
			t.Fatalf("expected 2 top-level nodes, got %d", len(nodes))
		}
		if nodes[0].Slug != "using-k6" {
			t.Errorf("first node slug = %q", nodes[0].Slug)
		}
		if !nodes[0].HasMoreChildren {
			t.Error("using-k6 should have HasMoreChildren at depth 1")
		}
		if len(nodes[0].Children) != 0 {
			t.Error("children should not be expanded at depth 1")
		}
	})

	t.Run("top level depth 2", func(t *testing.T) {
		t.Parallel()
		nodes, err := BuildSectionTree(sections, "", 2)
		if err != nil {
			t.Fatal(err)
		}
		if len(nodes[0].Children) != 1 {
			t.Fatalf("using-k6 should have 1 child at depth 2, got %d", len(nodes[0].Children))
		}
		child := nodes[0].Children[0]
		if child.Slug != "using-k6/scenarios" {
			t.Errorf("child slug = %q", child.Slug)
		}
		if !child.HasMoreChildren {
			t.Error("scenarios should have HasMoreChildren at depth 2")
		}
	})

	t.Run("rooted tree", func(t *testing.T) {
		t.Parallel()
		nodes, err := BuildSectionTree(sections, "using-k6/scenarios", 1)
		if err != nil {
			t.Fatal(err)
		}
		if len(nodes) != 2 {
			t.Fatalf("expected 2 children of scenarios, got %d", len(nodes))
		}
		if nodes[0].Slug != "using-k6/scenarios/open-model" {
			t.Errorf("first child slug = %q", nodes[0].Slug)
		}
	})

	t.Run("invalid depth", func(t *testing.T) {
		t.Parallel()
		_, err := BuildSectionTree(sections, "", 0)
		if err == nil {
			t.Error("expected error for depth 0")
		}
	})

	t.Run("unknown root slug", func(t *testing.T) {
		t.Parallel()
		_, err := BuildSectionTree(sections, "nonexistent", 1)
		if err == nil {
			t.Error("expected error for unknown root slug")
		}
	})
}

func TestNodesToDTO(t *testing.T) {
	t.Parallel()

	sections := testSectionsForTree()
	nodes, err := BuildSectionTree(sections, "", 2)
	if err != nil {
		t.Fatal(err)
	}

	dtos := NodesToDTO(nodes)
	if len(dtos) != 2 {
		t.Fatalf("expected 2 DTOs, got %d", len(dtos))
	}
	if dtos[0].Slug != "using-k6" {
		t.Errorf("first DTO slug = %q", dtos[0].Slug)
	}
	if dtos[0].ChildCount != 1 {
		t.Errorf("child count = %d, want 1", dtos[0].ChildCount)
	}
	if len(dtos[0].Children) != 1 {
		t.Fatalf("expected 1 child DTO, got %d", len(dtos[0].Children))
	}
	if dtos[0].Children[0].Slug != "using-k6/scenarios" {
		t.Errorf("child DTO slug = %q", dtos[0].Children[0].Slug)
	}
}

func TestNodesToDTOEmpty(t *testing.T) {
	t.Parallel()

	dtos := NodesToDTO(nil)
	if len(dtos) != 0 {
		t.Errorf("expected empty DTOs, got %d", len(dtos))
	}
}
