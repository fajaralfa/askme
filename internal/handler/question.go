package handler

import (
	"net/http"

	"github.com/fajaralfa/askme/internal/helper"
	"github.com/fajaralfa/askme/internal/repo"
	"github.com/fajaralfa/askme/internal/service/jwt"
)

type Question struct {
	QRepo repo.QuestionInterface
}

func (q *Question) FindAllAssociatedWithUser(w http.ResponseWriter, r *http.Request) {
	token, err := jwt.Parse(r.Header.Get("Authorization")[len("Bearer "):])
	if err != nil {
		ApiInternalErr(w, err)
		return
	}

	question, err := q.QRepo.FindAssociatedWithUser(token.Claims.(*jwt.Claims).Email)
	if err != nil {
		ApiInternalErr(w, err)
		return
	}

	helper.WriteJSON(w, question)
}
func (q *Question) Find(w http.ResponseWriter, r *http.Request) {

}
