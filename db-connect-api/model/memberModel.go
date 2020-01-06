package model

type Member struct {
	ID       int    `json:"id"`
	Username string `json:"name"`
	Type     string `json:"type"`
}

type MemberInput struct {
	Id int    `json:"id"`
	T  string `json:"t"`
}
