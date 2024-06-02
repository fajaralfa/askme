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
	Hash     hash.Hash
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
		ApiBadRequestErr(w, "register gagal", map[string]string{"email": "email telah digunakan"})
		return
	}

	payload.Password, err = a.Hash.Make(payload.Password)
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
	if user == nil || !a.Hash.Verify(payload.Password, user.Password) {
		ApiUnauthorizedErr(w, "email atau password salah", nil)
		return
	}

	token, err := jwt.Create(user.Email)
	if err != nil {
		ApiInternalErr(w, err)
		return
	}

	resp := model.Response{
		Status: "success",
		Data:   model.JWT{AccessToken: token},
	}
	helper.WriteJSON(w, resp)
}
