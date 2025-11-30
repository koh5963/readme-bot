// Package readme defines data models used by the readme-bot.
package readme

// Response represents the response payload produced by the readme-bot.
type Response struct {
	ReadmeLatestChange string    `json:"readme_latest_change"`
	ChangelogEntry     Changelog `json:"changelog_entry"`
}

// Changelog describes a single changelog entry for the README.
type Changelog struct {
	Version string   `json:"version"`
	Date    string   `json:"date"`
	Changes []string `json:"changes"`
}
