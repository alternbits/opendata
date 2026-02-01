# Altern Open Data

Data-driven awesome list. Edit YAML, run the compiler, get markdown.

## What it is

- Curated list content lives in **YAML**, not in the readme.
- A **Go compiler** reads that YAML and writes the list (e.g. `readme.md` or whatever you set).
- You change **data** and **meta**; the readme is generated.

## Layout

- **`data/`** — One `.yml` per item. Filename = slug (e.g. `altern.ai.yml`). Fields: name, slug, url, oneliner, main_category, position, date_added, etc.
- **`meta/`** — `info.yml` (title, description, badges, footer, position_order) and `categories.yml` (category id + name).
- **`config.yml`** — Optional. Set **`output`** to change the generated filename (default `readme.md`).
- **`compiler/`** — Go app. Builds from `data/` + `meta/`, validates, then writes the list.

## How to run

```bash
make run
```

Builds the compiler and writes the list to the path from `config.yml` (or `readme.md`).  
Or: `make build` then `./compiler/bin/compile`. Use `-data`, `-meta`, `-out` to override paths.

## CI

GitHub Action on push/PR: runs the compiler and fails if the generated file is out of date. Uses **`config.yml`** for the output path.

## Docs

- **CONTRIBUTING.md** — Data format, meta fields, how to add entries and badges.
