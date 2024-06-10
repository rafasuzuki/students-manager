package models

type Student struct {
	ID        int64  `json:"id"`
	Nome      string `json:"nome"`
	Graduacao string `json:"graduacao"`
	Idade     int64  `json:"idade"`
	Email     string `json:"email"`
	Senha     string `json:"senha"`
}
