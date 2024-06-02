package model

type Question struct {
	Id     int64  `json:"id"`
	Date   string `json:"date"`
	Text   string `json:"text"`
	UserId int64  `json:"user_id"`
}
