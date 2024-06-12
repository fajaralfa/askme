package handler

import (
	"net/http"

	"github.com/fajaralfa/askme/internal/helper"
	"github.com/fajaralfa/askme/internal/model"
	"github.com/fajaralfa/askme/internal/repo"
	"github.com/fajaralfa/askme/internal/service/hash"
	"github.com/fajaralfa/askme/internal/service/jwt"
)

type Auth struct {
	UserRepo repo.UserInterface
}

func (a *Auth) Register(w http.ResponseWriter, r *http.Request) {
	payload := new(model.RegisterReq)
	err := helper.ReadJSON(r.Body, &payload)
	if err != nil {
		ApiInternalErr(w, err)
		return
	}

	user, err := a.UserRepo.FindByEmail(payload.Email)
	if err != nil {
		ApiInternalErr(w, err)
		return
	}
	if user != nil {
		ApiUnauthorizedErr(w, "register gagal", map[string]string{"email": "email telah digunakan"})
		return
	}

	payload.Password, err = hash.Make(payload.Password)
	if err != nil {
		ApiInternalErr(w, err)
		return
	}

	_, err = a.UserRepo.Create(payload.Email, payload.Password, "")
	if err != nil {
		ApiInternalErr(w, err)
		return
	}

	resp := model.Response{
		Status: "success",
		Data:   nil,
	}
	helper.WriteJSON(w, resp)
}

func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	payload := new(model.LoginReq)
	err := helper.ReadJSON(r.Body, &payload)
	if err != nil {
		ApiInternalErr(w, err)
		return
	}

	user, err := a.UserRepo.FindByEmail(payload.Email)
	if err != nil {
		ApiInternalErr(w, err)
		return
	}
	if user == nil || !hash.Verify(payload.Password, user.Password) {
		ApiUnauthorizedErr(w, "email atau password salah", nil)
		return
	}

	token, err := jwt.Create(user.Id, user.Email)
	if err != nil {
		ApiInternalErr(w, err)
		return
	}

	user.Password = ""

	resp := model.Response{
		Status: "success",
		Data: map[string]any{
			"user":        user,
			"accessToken": token,
		},
	}
	helper.WriteJSON(w, resp)
}
