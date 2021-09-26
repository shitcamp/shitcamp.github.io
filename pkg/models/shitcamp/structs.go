package shitcamp

type Team struct {
	Name  string   `json:"name"`
	Captain string `json:"captain"`
	Users []string `json:"users"`
}

type TeamsInfo struct {
	Teams []*Team `json:"teams"`
}
