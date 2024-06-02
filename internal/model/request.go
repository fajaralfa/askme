package model

type RegisterReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AskQuestionReq struct {
	Question    string `json:"question"`
	TargetEmail string `json:"targetEmail"`
}
