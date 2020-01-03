package model

type Member struct {
	ID       int    `json:"id"`
	Username string `json:"name"`
	Type     string `json:"type"`
}
