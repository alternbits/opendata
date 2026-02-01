package compile

import (
	"fmt"
	"sort"
	"strings"
)

// Render builds the full readme markdown.
func Render(info *Info, categories []Category, itemsByCategory map[string][]*Item) string {
	var b strings.Builder
	// Title
	b.WriteString("# ")
	b.WriteString(info.Name)
	b.WriteString("\n\n")
	// Badge
	if info.BadgeURL != "" && info.BadgeLink != "" {
		b.WriteString("[![Awesome](")
		b.WriteString(info.BadgeURL)
		b.WriteString(")](")
		b.WriteString(info.BadgeLink)
		b.WriteString(")\n\n")
	}
	// Intro
	b.WriteString(strings.TrimSpace(info.Description))
	b.WriteString("\n\n")
	// Table of contents
	b.WriteString("## Contents\n\n")
	for _, c := range categories {
		anchor := "#" + strings.ToLower(c.ID)
		b.WriteString(fmt.Sprintf("- [%s](%s)\n", c.Name, anchor))
	}
	b.WriteString("\n")
	// Sections
	for _, c := range categories {
		items := itemsByCategory[c.ID]
		if len(items) == 0 {
			continue
		}
		b.WriteString("## ")
		b.WriteString(c.Name)
		b.WriteString("\n\n")
		for _, item := range items {
			b.WriteString("- [")
			b.WriteString(escapeMarkdownBracket(item.Name))
			b.WriteString("](")
			b.WriteString(item.URL)
			b.WriteString(") - ")
			b.WriteString(OnelinerSuffix(item.OnelinerValue()))
			b.WriteString("\n")
		}
		b.WriteString("\n")
	}
	// Footer
	if info.License != "" || info.Contribute != "" {
		b.WriteString("---\n\n")
		if info.License != "" {
			b.WriteString("**License**: ")
			b.WriteString(info.License)
			b.WriteString("\n\n")
		}
		if info.Contribute != "" {
			b.WriteString("See [")
			b.WriteString(info.Contribute)
			b.WriteString("](")
			b.WriteString(info.Contribute)
			b.WriteString(") for contribution guidelines.\n")
		}
	}
	return b.String()
}

func escapeMarkdownBracket(s string) string {
	return strings.ReplaceAll(s, "]", "\\]")
}

// GroupByMainCategory builds category -> sorted items from full item set.
func GroupByMainCategory(categories []Category, items map[string]*Item) map[string][]*Item {
	out := make(map[string][]*Item)
	for _, item := range items {
		out[item.MainCategory] = append(out[item.MainCategory], item)
	}
	for _, list := range out {
		sort.Slice(list, func(i, j int) bool {
			return strings.ToLower(list[i].Name) < strings.ToLower(list[j].Name)
		})
	}
	return out
}
