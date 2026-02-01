package compile

import (
	"fmt"
	"sort"
	"strings"
)

// Render builds the full readme markdown.
// When reviewLinksEnabled is true, items with a non-empty Review link show "*[review](url)*" between the main link and the description.
func Render(info *Info, categories []Category, itemsByCategory map[string][]*Item, reviewLinksEnabled bool) string {
	var b strings.Builder
	// Title
	b.WriteString("# ")
	b.WriteString(info.Name)
	b.WriteString("\n\n")
	// Badges
	if len(info.Badges) > 0 {
		for i, badge := range info.Badges {
			if badge.URL != "" && badge.Link != "" {
				if i > 0 {
					b.WriteString(" ")
				}
				b.WriteString("[![Awesome](")
				b.WriteString(badge.URL)
				b.WriteString(")](")
				b.WriteString(badge.Link)
				b.WriteString(")")
			}
		}
		b.WriteString("\n\n")
	} else if info.BadgeURL != "" && info.BadgeLink != "" {
		b.WriteString("[![Awesome](")
		b.WriteString(info.BadgeURL)
		b.WriteString(")](")
		b.WriteString(info.BadgeLink)
		b.WriteString(")\n\n")
	}
	// Intro
	b.WriteString(strings.TrimSpace(info.Description))
	b.WriteString("\n\n")
	// Table of contents (categories only)
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
			b.WriteString(")")
			if reviewLinksEnabled && strings.TrimSpace(item.Review) != "" {
				b.WriteString(" *[review](")
				b.WriteString(item.Review)
				b.WriteString(")*")
			}
			b.WriteString(" - ")
			b.WriteString(OnelinerSuffix(item.OnelinerValue()))
			b.WriteString("\n")
		}
		b.WriteString("\n")
	}
	// Footer
	if info.License != "" || info.Contribute != "" || info.Footer != "" {
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
			b.WriteString(") for contribution guidelines.\n\n")
		}
		if info.Footer != "" {
			b.WriteString(strings.TrimSpace(info.Footer))
			b.WriteString("\n")
		}
	}
	return b.String()
}

func escapeMarkdownBracket(s string) string {
	return strings.ReplaceAll(s, "]", "\\]")
}

// GroupByMainCategory builds category -> sorted items from full item set.
// If info.PositionOrder is non-empty, items are sorted by position tier (then by name within tier).
// Unknown or empty position is treated as last tier. If PositionOrder is empty, sort by name only.
func GroupByMainCategory(info *Info, categories []Category, items map[string]*Item) map[string][]*Item {
	out := make(map[string][]*Item)
	for _, item := range items {
		out[item.MainCategory] = append(out[item.MainCategory], item)
	}
	positionIndex := make(map[string]int)
	for i, p := range info.PositionOrder {
		positionIndex[p] = i
	}
	lastIndex := len(info.PositionOrder)
	for _, list := range out {
		sort.Slice(list, func(a, b int) bool {
			if len(info.PositionOrder) == 0 {
				return strings.ToLower(list[a].Name) < strings.ToLower(list[b].Name)
			}
			pi, oki := positionIndex[list[a].Position]
			pj, okj := positionIndex[list[b].Position]
			if !oki || list[a].Position == "" {
				pi = lastIndex
			}
			if !okj || list[b].Position == "" {
				pj = lastIndex
			}
			if pi != pj {
				return pi < pj
			}
			return strings.ToLower(list[a].Name) < strings.ToLower(list[b].Name)
		})
	}
	return out
}
