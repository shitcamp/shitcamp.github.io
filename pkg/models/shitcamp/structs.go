package shitcamp

type Team struct {
	Name         string   `json:"name"`
	Captain      string   `json:"captain"`
	Users        []string `json:"users"`
	ThumbnailURL string   `json:"thumbnail_url"`
}

type TeamsInfo struct {
	Teams []*Team `json:"teams"`
	// Winner- 100 team points
	// Trivia + Roulette
	// Hide n seek winner
}
