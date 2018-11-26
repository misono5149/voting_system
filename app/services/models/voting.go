package models

type Voting struct {
	Id          int    `json:"id" gorm:"not null, default=0, auto_increment"`
	CandidateId int    `json:"candidate_id"`
	ElectionId  int    `json:"election_id"`
	Auto_hash   string `json:"auto_hash"`
}

type Votings []Voting
