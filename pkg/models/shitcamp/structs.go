package shitcamp

type Team struct {
	Name  string   `json:"name"`
	Users []string `json:"users"`
}

type TeamsInfo struct {
	Teams []*Team `json:"teams"`
}
