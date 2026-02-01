package compile

import (
	"strings"
	"testing"
)

func TestGroupByMainCategory_sortByName(t *testing.T) {
	info := &Info{} // no PositionOrder
	categories := []Category{{ID: "tools", Name: "Tools"}}
	items := map[string]*Item{
		"b": {Name: "B", MainCategory: "tools"},
		"a": {Name: "A", MainCategory: "tools"},
	}
	byCat := GroupByMainCategory(info, categories, items)
	list := byCat["tools"]
	if len(list) != 2 {
		t.Fatalf("expected 2 items, got %d", len(list))
	}
	if list[0].Name != "A" || list[1].Name != "B" {
		t.Errorf("expected order A, B; got %s, %s", list[0].Name, list[1].Name)
	}
}

func TestGroupByMainCategory_sortByPosition(t *testing.T) {
	info := &Info{PositionOrder: []string{"featured", "ordinary"}}
	categories := []Category{{ID: "tools", Name: "Tools"}}
	items := map[string]*Item{
		"a": {Name: "A", MainCategory: "tools", Position: "ordinary"},
		"b": {Name: "B", MainCategory: "tools", Position: "featured"},
	}
	byCat := GroupByMainCategory(info, categories, items)
	list := byCat["tools"]
	if len(list) != 2 {
		t.Fatalf("expected 2 items, got %d", len(list))
	}
	if list[0].Name != "B" || list[1].Name != "A" {
		t.Errorf("expected featured first (B, A); got %s, %s", list[0].Name, list[1].Name)
	}
}

func TestRender_containsTitleAndContents(t *testing.T) {
	info := &Info{Name: "Awesome List", Description: "A list."}
	categories := []Category{{ID: "tools", Name: "Tools"}}
	items := map[string]*Item{
		"foo": {
			Name: "Foo", URL: "https://foo.com", MainCategory: "tools",
		},
	}
	items["foo"].onelinerValue = "A tool."
	byCat := GroupByMainCategory(info, categories, items)
	md := Render(info, categories, byCat)
	if !strings.Contains(md, "# Awesome List") {
		t.Error("output should contain title")
	}
	if !strings.Contains(md, "## Contents") {
		t.Error("output should contain Contents")
	}
	if !strings.Contains(md, "## Tools") {
		t.Error("output should contain Tools section")
	}
	if !strings.Contains(md, "[Foo](https://foo.com)") {
		t.Error("output should contain item link")
	}
}

func TestEscapeMarkdownBracket(t *testing.T) {
	// escapeMarkdownBracket is unexported; test via Render with item name containing ]
	info := &Info{Name: "T", Description: "D"}
	categories := []Category{{ID: "x", Name: "X"}}
	items := map[string]*Item{
		"a": {Name: "Item [extra]", URL: "https://x.com", MainCategory: "x"},
	}
	items["a"].onelinerValue = "Desc."
	byCat := GroupByMainCategory(info, categories, items)
	md := Render(info, categories, byCat)
	if !strings.Contains(md, `\]`) {
		t.Error("expected bracket in name to be escaped")
	}
}
