package models

type EditPageData struct {
	Title     string          `json:"title"`
	User      User            `json:"user"`
	Countries []string        `json:"countries"`
	SportsMap map[string]bool `json:"sportsMap"`
}
