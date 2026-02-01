package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/alternbits/opendata/compiler/compile"
	"gopkg.in/yaml.v3"
)

type config struct {
	Output            string `yaml:"output"`
	ReviewLinksEnabled bool   `yaml:"review_links_enabled"`
}

func main() {
	dataDir := flag.String("data", "data", "directory containing data/*.yml")
	metaDir := flag.String("meta", "meta", "directory containing meta/info.yml and meta/categories.yml")
	outPath := flag.String("out", "readme.md", "output path for generated readme (overridden by config.yml if set)")
	printOut := flag.Bool("print-out", false, "print effective output path and exit (for CI)")
	flag.Parse()

	out := resolveOutputPath(*outPath)

	if *printOut {
		abs, _ := filepath.Abs(out)
		fmt.Println(abs)
		return
	}

	if err := run(*dataDir, *metaDir, out); err != nil {
		fmt.Fprintf(os.Stderr, "compile: %v\n", err)
		os.Exit(1)
	}
}

// resolveOutputPath returns the output path, applying config.yml if present.
func resolveOutputPath(flagOut string) string {
	out := flagOut
	if cwd, err := os.Getwd(); err == nil {
		if cfg := loadConfig(filepath.Join(cwd, "config.yml")); cfg != nil && strings.TrimSpace(cfg.Output) != "" {
			out = strings.TrimSpace(cfg.Output)
		}
	}
	return out
}

func run(dataDir, metaDir, outPath string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	if !filepath.IsAbs(dataDir) {
		dataDir = filepath.Join(cwd, dataDir)
	}
	if !filepath.IsAbs(metaDir) {
		metaDir = filepath.Join(cwd, metaDir)
	}
	if !filepath.IsAbs(outPath) {
		outPath = filepath.Join(cwd, outPath)
	}

	info, err := compile.LoadInfo(metaDir)
	if err != nil {
		return err
	}
	categories, err := compile.LoadCategories(metaDir)
	if err != nil {
		return err
	}
	items, err := compile.LoadData(dataDir)
	if err != nil {
		return err
	}

	if errs := compile.Validate(info, categories, items); len(errs) > 0 {
		for _, e := range errs {
			fmt.Fprintf(os.Stderr, "%v\n", e)
		}
		return fmt.Errorf("validation failed: %d error(s)", len(errs))
	}

	reviewLinksEnabled := false
	if cfg := loadConfig(filepath.Join(cwd, "config.yml")); cfg != nil {
		reviewLinksEnabled = cfg.ReviewLinksEnabled
	}

	byCategory := compile.GroupByMainCategory(info, categories, items)
	md := compile.Render(info, categories, byCategory, reviewLinksEnabled)
	if err := os.WriteFile(outPath, []byte(md), 0644); err != nil {
		return fmt.Errorf("write %s: %w", outPath, err)
	}
	return nil
}

// loadConfig reads config.yml from path. Returns nil if file is missing or invalid.
func loadConfig(path string) *config {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var c config
	if err := yaml.Unmarshal(data, &c); err != nil {
		return nil
	}
	return &c
}
