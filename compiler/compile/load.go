package compile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// LoadInfo reads meta/info.yml.
func LoadInfo(metaDir string) (*Info, error) {
	data, err := os.ReadFile(filepath.Join(metaDir, "info.yml"))
	if err != nil {
		return nil, fmt.Errorf("meta/info.yml: %w", err)
	}
	var info Info
	if err := yaml.Unmarshal(data, &info); err != nil {
		return nil, fmt.Errorf("meta/info.yml: %w", err)
	}
	return &info, nil
}

// LoadCategories reads meta/categories.yml.
// Supports list of objects {id, name} or list of strings (id/name derived).
func LoadCategories(metaDir string) ([]Category, error) {
	data, err := os.ReadFile(filepath.Join(metaDir, "categories.yml"))
	if err != nil {
		return nil, fmt.Errorf("meta/categories.yml: %w", err)
	}
	var raw []yaml.Node
	if err := yaml.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("meta/categories.yml: %w", err)
	}
	var out []Category
	for i, node := range raw {
		if node.Kind == yaml.ScalarNode {
			name := strings.TrimSpace(node.Value)
			id := slugify(name)
			out = append(out, Category{ID: id, Name: name})
			continue
		}
		if node.Kind == yaml.MappingNode {
			var c Category
			if err := node.Decode(&c); err != nil {
				return nil, fmt.Errorf("meta/categories.yml[%d]: %w", i, err)
			}
			if c.ID == "" {
				c.ID = slugify(c.Name)
			}
			out = append(out, c)
			continue
		}
		return nil, fmt.Errorf("meta/categories.yml[%d]: expected string or object", i)
	}
	return out, nil
}

func slugify(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' || r == '_' {
			b.WriteRune(r)
		} else if r == ' ' || r == '/' {
			b.WriteRune('-')
		}
	}
	return strings.Trim(b.String(), "-")
}

// LoadData reads all data/*.yml and returns items keyed by filename stem (slug).
func LoadData(dataDir string) (map[string]*Item, error) {
	entries, err := os.ReadDir(dataDir)
	if err != nil {
		return nil, fmt.Errorf("data: %w", err)
	}
	out := make(map[string]*Item)
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".yml") {
			continue
		}
		stem := strings.TrimSuffix(e.Name(), ".yml")
		path := filepath.Join(dataDir, e.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", path, err)
		}
		var item Item
		if err := yaml.Unmarshal(data, &item); err != nil {
			return nil, fmt.Errorf("%s: %w", path, err)
		}
		if item.Oneliner != "" {
			item.onelinerValue = item.Oneliner
		} else {
			item.onelinerValue = item.OnlineDesc
		}
		if item.Slug == "" {
			item.Slug = stem
		}
		out[stem] = &item
	}
	return out, nil
}
