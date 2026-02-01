# Altern Open Data

**Repo:** [github.com/alternbits/opendata](https://github.com/alternbits/opendata)

Data-driven awesome list. Edit YAML, run the compiler, get markdown.

## Use as a GitHub Action

Use this in **any repo** to validate and compile your awesome list. No need to copy the compiler.

### 1. Create your repo layout

```
your-repo/
├── data/
│   └── example.yml
├── meta/
│   ├── info.yml
│   └── categories.yml
├── config.yml          # Optional
└── .github/workflows/validate.yml
```

### 2. Add meta files

**`meta/info.yml`** — List title, description, badges:
```yaml
name: Awesome Tools
description: A curated list of awesome tools.
badges:
  - url: https://awesome.re/badge.svg
    link: https://awesome.re
position_order:
  - featured
  - popular
  - ordinary
```

**`meta/categories.yml`** — Categories for your list:
```yaml
- id: tools
  name: Tools
- id: resources
  name: Resources
```

### 3. Add data files

One file per item in `data/`. Filename = slug (e.g. `my-tool.yml`).

**`data/my-tool.yml`**:
```yaml
name: My Tool
slug: my-tool
url: https://example.com
oneliner: A great tool for developers.
main_category: tools
position: featured
date_added: 2025-02-01
```

### 4. Add the workflow

**`.github/workflows/validate.yml`**:
```yaml
name: Validate and compile
on:
  push:
    branches: [main]
  pull_request:
    branches: [main]
jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: alternbits/opendata@v1
```

With custom inputs:
```yaml
      - uses: alternbits/opendata@v1
        with:
          data-dir: data
          meta-dir: meta
          fail-if-outdated: 'true'
```

### 5. Optional: config.yml

Set output filename (default is `readme.md`):
```yaml
output: LIST.md
```

### Action inputs

| Input | Default | Description |
|-------|---------|-------------|
| `data-dir` | `data` | Directory with data YAML files |
| `meta-dir` | `meta` | Directory with info.yml and categories.yml |
| `fail-if-outdated` | `true` | Fail if generated file differs from committed |

### What the action does

1. Builds the compiler
2. Reads `data/`, `meta/`, and optional `config.yml` from your repo
3. Validates YAML structure and category references
4. Writes the generated list (e.g. `readme.md`)
5. Fails if the file changed (so you know to commit the update)

---

## What it is

- Curated list content lives in **YAML**, not in the readme.
- A **Go compiler** reads that YAML and writes the list (e.g. `readme.md` or whatever you set).
- You change **data** and **meta**; the readme is generated.

## Layout

- **`template/`** — Standalone test template (separate project). Has everything a consumer repo needs: `data/`, `meta/`, `config.yml`, `.github/workflows/validate.yml` (uses the action), `README.md`, `CONTRIBUTING.md`, `.gitignore`, and the generated `readme-sample.md`. Copy this folder into a new repo to start an awesome list; no compiler required.
- **`compiler/`** — Go app. Builds from `data/` + `meta/`, validates, then writes the list.

## How to run

```bash
make run
```

Builds the compiler and runs it on **`template/`** (writes `template/readme-sample.md`).  
To run on another dir: `make build` then `cd your-dir && ../compiler/bin/compile`. Use `-data`, `-meta`, `-out` to override paths.

## CI

Root workflow runs the compiler on `template/` and fails if `template/readme-sample.md` is out of date.

## Docs

- **CONTRIBUTING.md** — Data format, meta fields, how to add entries and badges.
