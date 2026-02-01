# Contributing

This list is generated from YAML data. **Do not edit `readme.md` directly.**

## How to contribute

1. **Add or edit entries** in the `data/` folder. Each file is named by slug (e.g. `altern.ai.yml`).
2. **Categories** are defined in `meta/categories.yml`. Use an existing category's `id` for `main_category` in your entry.
3. **Validate** by running the compiler locally:

   ```bash
   make run
   ```
   Or `make build` then `./compiler/bin/compile`.

   If validation fails, fix the reported errors in your YAML before submitting a PR.

4. **Commit** your `.yml` changes and the updated `readme.md` (run `make run` to regenerate; CI will fail if `readme.md` is out of date).

## Config (`config.yml`)

Optional **`config.yml`** in the project root can set:

- **`output`** (string): Filename or path for the generated list. Default is `readme.md`; if non-empty, the compiler writes to that file instead. Example: `output: LIST.md`.
- **`review_links_enabled`** (boolean): If `true`, items with a **`review`** link in data YAML show it in the readme between the main link and the description as: `*[review](url)*`. Default is `false`.

## Data file format (`data/{slug}.yml`)

- **name** (required): Display name.
- **slug** (required): Must match the filename (without `.yml`), e.g. `altern.ai`.
- **url** (required): Primary link.
- **oneliner** or **online_description** (required): Short description for the list item.
- **description** (optional): Longer description.
- **main_category** (required): Must be an `id` from `meta/categories.yml`.
- **categories** (optional): Additional category ids from `meta/categories.yml`.
- **position** (optional): Tier for ordering within the category. Allowed values and their order are defined in `meta/info.yml` under `position_order` (e.g. `featured`, `popular`, `ordinary`, `new`, `dead`). Items with the same position are sorted by name. If omitted, the item is treated as the last tier. If `position_order` is omitted in meta, all items are sorted by name only.
- **review** (optional): URL to a review. Shown only when **`review_links_enabled`** is `true` in `config.yml`, as `*[review](url)*` between the main link and the description.
- **date_added** (optional): Date the entry was added, e.g. `2025-02-01` (ISO 8601).
- **date_modified** (optional): Date the entry was last updated, e.g. `2025-02-01`.

## Meta: badges (`meta/info.yml`)

**`badges`** (list of objects) adds one or more badges after the title. Each entry has **`url`** (image URL) and **`link`** (click target). They are rendered in order on one line, separated by spaces. Example:

```yaml
badges:
  - url: https://awesome.re/badge.svg
    link: https://awesome.re
  - url: https://other.example/badge.svg
    link: https://other.example
```

For a single badge you can still use **`badge_url`** and **`badge_link`** (backward compatible).

## Meta: position order (`meta/info.yml`)

Optional **`position_order`** (list of strings) defines the display order of position tiers within each category. First in the list appears at the top (e.g. featured, then popular, then ordinary, then new, then dead). Omit to use name-only sort (current default behavior).

## Meta: footer (`meta/info.yml`)

Optional **`footer`** (string, multiline with `|`) is raw markdown rendered at the bottom of the readme, after the horizontal rule. Use for extra links, credits, or notes.

## Pull requests

CI runs on every push and pull request. It will:

1. Run the compiler to validate all YAML structure and references.
2. Fail if `readme.md` would change (i.e. you must run `make run` and commit the updated `readme.md` with your data changes).

Thanks for contributing.
