package compile

// Info is repo-level metadata from meta/info.yml.
type Info struct {
	Name          string   `yaml:"name"`
	Description   string   `yaml:"description"`
	PositionOrder []string `yaml:"position_order"`
	BadgeURL      string   `yaml:"badge_url"`
	BadgeLink     string   `yaml:"badge_link"`
	License       string   `yaml:"license"`
	Contribute    string   `yaml:"contribute"`
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
	Position      string   `yaml:"position"` // tier for ordering: featured, popular, ordinary, new, dead, etc.
	onelinerValue string   // set after load: oneliner or online_description
}

// Oneliner returns the short description for the list item.
func (i *Item) OnelinerValue() string {
	return i.onelinerValue
}
