package compile

// Badge is a single badge (image URL + link) for the readme header.
type Badge struct {
	URL  string `yaml:"url"`
	Link string `yaml:"link"`
}

// Info is repo-level metadata from meta/info.yml.
type Info struct {
	Name          string   `yaml:"name"`
	Description   string   `yaml:"description"`
	PositionOrder []string `yaml:"position_order"`
	Badges        []Badge  `yaml:"badges"`       // multiple badges (preferred)
	BadgeURL      string   `yaml:"badge_url"`    // single badge (backward compat)
	BadgeLink     string   `yaml:"badge_link"`   // single badge (backward compat)
	License       string   `yaml:"license"`
	Contribute    string   `yaml:"contribute"`
	Footer        string   `yaml:"footer"` // optional markdown content for footer section
}

// Category is a single category from meta/categories.yml.
// Supports both object form {id, name} and string form (id/name derived).
type Category struct {
	ID   string `yaml:"id"`
	Name string `yaml:"name"`
}

// Item is a single list entry from data/{slug}.yml.
type Item struct {
	Name          string   `yaml:"name"`
	Slug          string   `yaml:"slug"`
	URL           string   `yaml:"url"`
	Oneliner      string   `yaml:"oneliner"`
	OnlineDesc    string   `yaml:"online_description"`
	Description   string   `yaml:"description"`
	MainCategory  string   `yaml:"main_category"`
	Categories    []string `yaml:"categories"`
	Position      string   `yaml:"position"`      // tier for ordering: featured, popular, ordinary, new, dead, etc.
	DateAdded     string   `yaml:"date_added"`    // optional, e.g. 2025-02-01
	DateModified  string   `yaml:"date_modified"`  // optional, e.g. 2025-02-01
	onelinerValue string   // set after load: oneliner or online_description
}

// Oneliner returns the short description for the list item.
func (i *Item) OnelinerValue() string {
	return i.onelinerValue
}
