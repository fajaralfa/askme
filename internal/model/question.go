package model

type Question struct {
	Id        int64  `json:"id"`
	Question  string `json:"question"`
	CreatedAt string `json:"created_at"`
	UserId    int64  `json:"user_id"`
}
