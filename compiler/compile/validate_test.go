package compile

import (
	"testing"
)

func TestValidate_emptyInfo(t *testing.T) {
	info := &Info{}
	categories := []Category{{ID: "tools", Name: "Tools"}}
	items := map[string]*Item{}
	errs := Validate(info, categories, items)
	if len(errs) < 2 {
		t.Fatalf("expected at least 2 errors for empty info, got %d", len(errs))
	}
}

func TestValidate_valid(t *testing.T) {
	info := &Info{Name: "Test", Description: "Desc"}
	categories := []Category{{ID: "tools", Name: "Tools"}}
	item := &Item{
		Name: "Foo", Slug: "foo", URL: "https://foo.com",
		MainCategory: "tools",
	}
	item.onelinerValue = "A tool."
	items := map[string]*Item{"foo": item}
	errs := Validate(info, categories, items)
	if len(errs) != 0 {
		t.Fatalf("expected no errors, got %v", errs)
	}
}

func TestValidate_slugMismatch(t *testing.T) {
	info := &Info{Name: "Test", Description: "Desc"}
	categories := []Category{{ID: "tools", Name: "Tools"}}
	item := &Item{
		Name: "Foo", Slug: "wrong", URL: "https://foo.com",
		MainCategory: "tools",
	}
	item.onelinerValue = "A tool."
	items := map[string]*Item{"foo": item}
	errs := Validate(info, categories, items)
	if len(errs) == 0 {
		t.Fatal("expected error for slug mismatch")
	}
}

func TestValidate_unknownCategory(t *testing.T) {
	info := &Info{Name: "Test", Description: "Desc"}
	categories := []Category{{ID: "tools", Name: "Tools"}}
	item := &Item{
		Name: "Foo", Slug: "foo", URL: "https://foo.com",
		MainCategory: "unknown",
	}
	item.onelinerValue = "A tool."
	items := map[string]*Item{"foo": item}
	errs := Validate(info, categories, items)
	if len(errs) == 0 {
		t.Fatal("expected error for unknown main_category")
	}
}

func TestOnelinerSuffix(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"hello", "hello."},
		{"hello.", "hello."},
		{"  hi  ", "hi."},
		{"", ""},
	}
	for _, tt := range tests {
		got := OnelinerSuffix(tt.in)
		if got != tt.want {
			t.Errorf("OnelinerSuffix(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
