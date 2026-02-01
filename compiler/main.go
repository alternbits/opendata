package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alternbits/awesome-altern/compiler/compile"
)

func main() {
	dataDir := flag.String("data", "data", "directory containing data/*.yml")
	metaDir := flag.String("meta", "meta", "directory containing meta/info.yml and meta/categories.yml")
	outPath := flag.String("out", "readme.md", "output path for generated readme.md")
	flag.Parse()

	if err := run(*dataDir, *metaDir, *outPath); err != nil {
		fmt.Fprintf(os.Stderr, "compile: %v\n", err)
		os.Exit(1)
	}
}

func run(dataDir, metaDir, outPath string) error {
	// Resolve paths relative to current directory
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

	byCategory := compile.GroupByMainCategory(categories, items)
	md := compile.Render(info, categories, byCategory)
	if err := os.WriteFile(outPath, []byte(md), 0644); err != nil {
		return fmt.Errorf("write %s: %w", outPath, err)
	}
	return nil
}
