package models

type Election struct {
	Title             string `json:"title"`
	Major             string `json:"major"`
	College           string `json:"college"`
	Content           string `json:"content"`
	ElectionStartTime string `json:"start_time"`
	ElectionEndTime   string `json:"end_time"`
	State             int    `json:"state"`
	Id                int    `json:"election_id"`
	AdminId           string `json:"admin"`
}

/*
[state] 1: 투표 전 2: 투표 중 3: 투표 완료
*/

type EndElectionCandidateInfo struct {
	ElectionId  int    `json:"election_id"`
	All_vote    int    `json:"all_vote"`
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
