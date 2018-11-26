package models

type Candidate struct {
	StudentId  string `json:"strudent_id"`
	Name       string `json:"name"`
	Major      string `json:"major"`
	College    string `json:"college"`
	Thumbnail  string `json:"thumbnail"`
	Resume     string `json:"resume"`
	Id         int    `json:"candidate_id"`
	ElectionId int    `json:"election_id"`
	//Poll       int    `json:"poll"`
}

type Candidates []Candidate
