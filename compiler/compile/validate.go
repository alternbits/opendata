package compile

import (
	"fmt"
	"strings"
)

// Validate checks meta and data; returns all errors.
func Validate(info *Info, categories []Category, items map[string]*Item) []error {
	var errs []error
	if info.Name == "" {
		errs = append(errs, fmt.Errorf("meta/info.yml: name is required"))
	}
	if info.Description == "" {
		errs = append(errs, fmt.Errorf("meta/info.yml: description is required"))
	}
	catIDs := make(map[string]string) // id -> name
	for _, c := range categories {
		if c.ID == "" {
			errs = append(errs, fmt.Errorf("meta/categories.yml: category id is required"))
			continue
		}
		catIDs[c.ID] = c.Name
	}
	for filenameSlug, item := range items {
		if item.Slug == "" {
			errs = append(errs, fmt.Errorf("data/%s.yml: slug is required", filenameSlug))
		} else if item.Slug != filenameSlug {
			errs = append(errs, fmt.Errorf("data/%s.yml: slug %q must match filename stem %q", filenameSlug, item.Slug, filenameSlug))
		}
		if item.Name == "" {
			errs = append(errs, fmt.Errorf("data/%s.yml: name is required", filenameSlug))
		}
		if item.URL == "" {
			errs = append(errs, fmt.Errorf("data/%s.yml: url is required", filenameSlug))
		}
		if item.OnelinerValue() == "" {
			errs = append(errs, fmt.Errorf("data/%s.yml: oneliner or online_description is required", filenameSlug))
		}
		if item.MainCategory == "" {
			errs = append(errs, fmt.Errorf("data/%s.yml: main_category is required", filenameSlug))
		} else if _, ok := catIDs[item.MainCategory]; !ok {
			errs = append(errs, fmt.Errorf("data/%s.yml: main_category %q not in meta/categories.yml", filenameSlug, item.MainCategory))
		}
		for _, cat := range item.Categories {
			if _, ok := catIDs[cat]; !ok {
				errs = append(errs, fmt.Errorf("data/%s.yml: category %q not in meta/categories.yml", filenameSlug, cat))
			}
		}
	}
	return errs
}

// OnelinerSuffix ensures the short description ends with a period for consistency.
func OnelinerSuffix(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	if !strings.HasSuffix(s, ".") {
		return s + "."
	}
	return s
}
