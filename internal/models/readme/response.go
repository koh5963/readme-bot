package models

type Response struct {
	ReadmeLatestChange string    `json:"readme_latest_change"`
	ChangelogEntry     Changelog `json:"changelog_entry"`
}

type Changelog struct {
	Version string   `json:"version"`
	Date    string   `json:"date"`
	Changes []string `json:"changes"`
}
