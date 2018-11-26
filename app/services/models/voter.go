package models

type Voter struct {
	Id        string `json:"id"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	StudentId string `json:"student_id"`
	Major     string `json:"major"`
	College   string `json:"college"` // ex) 공과 대학, 어문 대학, ...
	Mobile    string `json:"mobile"`
	Address   string `json:"address"`
	Email     string `json:"email"`
	Sex       string `json:"sex"`
	Birth     string `json:"birth"`
}
