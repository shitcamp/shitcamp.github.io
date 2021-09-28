package shitcamp

type Team struct {
	Name         string   `json:"name"`
	Captain      string   `json:"captain"`
	Users        []string `json:"users"`
	ThumbnailURL string   `json:"thumbnail_url"`
}

type ContestResult struct {
	ContestName string `json:"contest_name"`
	Points      int    `json:"points"`
	Description string `json:"description"`
	VodURL      string `json:"vod_url"`
}

type TeamResults struct {
	TeamName    string           `json:"team_name"`
	TotalPoints int              `json:"total_points"`
	Results     []*ContestResult `json:"results"`
}

type TeamsInfo struct {
	Teams   []*Team        `json:"teams"`
	Results []*TeamResults `json:"team_results"`
}
