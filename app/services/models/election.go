package models

import "time"

type Election struct {
	Title             string    `json:"title"`
	Major             string    `json:"major"`
	College           string    `json:"college"`
	Content           string    `json:"content"`
	ElectionStartTime time.Time `json:"start_time"`
	ElectionEndTime   time.Time `json:"end_time"`
	State             int       `json:"state"`
	Id                int       `json:"election_id" gorm:"default=0, primary_key, auto_increment"`
	AdminId           string    `json:"admin"`
}

/*
[state] 1: 투표 전 2: 투표 중 3: 투표 완료
*/

type EndElectionCandidateInfo struct {
	ElectionId  int    `json:"election_id" gorm:"primary_key, not_null"`
	All_vote    int    `json:"all_vote" gorm:"not_null"`
	CandidateId int    `json:"candidate_id"`
	Poll        int    `json:"poll"`
	StudentId   string `json:"strudent_id"`
	Name        string `json:"name"`
	Major       string `json:"major"`
	College     string `json:"college"`
	Thumbnail   string `json:"thumbnail"`
	Resume      string `json:"resume"`
}

type Elections []Election
type EndElectionResult []EndElectionCandidateInfo
