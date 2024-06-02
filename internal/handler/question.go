package handler

import (
	"fmt"
	"net/http"

	"github.com/fajaralfa/askme/internal/helper"
	"github.com/fajaralfa/askme/internal/model"
	"github.com/fajaralfa/askme/internal/repo"
	"github.com/fajaralfa/askme/internal/service/jwt"
	"github.com/gorilla/mux"
)

type Question struct {
	QRepo    repo.QuestionInterface
	UserRepo repo.UserInterface
}

func (q *Question) AskQuestion(w http.ResponseWriter, r *http.Request) {
	payload := new(model.AskQuestionReq)
	err := helper.ReadJSON(r.Body, payload)
	if err != nil {
		ApiInternalErr(w, err)
		return
	}

	targetUser, err := q.UserRepo.FindByEmail(payload.TargetEmail)
	if err != nil {
		ApiInternalErr(w, err)
		return
	}

	if targetUser == nil {
		ApiNotFoundErr(w, "pengguna dengan email ini tidak ditemukan")
		return
	}

	question, err := q.QRepo.Create(payload.Question, targetUser.Id)
	if err != nil {
		ApiInternalErr(w, err)
		return
	}

	resp := model.Response{
		Status: "success",
		Data: map[string]*model.Question{
			"question": question,
		},
	}

	helper.WriteJSON(w, resp)
}

func (q *Question) FindAllAssociatedWithUser(w http.ResponseWriter, r *http.Request) {
	claim, err := jwt.ParseGetClaims(r.Header.Get("Authorization")[len("Bearer "):])
	if err != nil {
		ApiInternalErr(w, err)
		return
	}

	questions, err := q.QRepo.FindAllByUserEmail(claim.Email)
	if err != nil {
		ApiInternalErr(w, err)
		return
	}

	resp := model.Response{
		Status: "success",
		Data: map[string][]model.Question{
			"questions": questions,
		},
	}

	helper.WriteJSON(w, resp)
}

func (q *Question) RemoveAssociatedWithUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	claim, err := jwt.ParseGetClaims(r.Header.Get("Authorization")[len("Bearer "):])
	if err != nil {
		ApiInternalErr(w, err)
		return
	}

	user, err := q.UserRepo.FindByEmail(claim.Email)
	if err != nil {
		ApiInternalErr(w, err)
		return
	}

	affected, err := q.QRepo.DeleteByQIdUserId(vars["id"], fmt.Sprint(user.Id))
	if err != nil {
		ApiInternalErr(w, err)
		return
	}

	if affected == 0 {
		ApiUnauthorizedErr(w, "pertanyaan ini tidak ditanyakan ke kamu, maaf aja ya", nil)
		return
	}

	resp := model.Response{
		Status: "success",
		Data:   nil,
	}

	helper.WriteJSON(w, resp)
}
