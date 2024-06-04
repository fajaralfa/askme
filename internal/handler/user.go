package handler

import (
	"net/http"

	"github.com/fajaralfa/askme/internal/helper"
	"github.com/fajaralfa/askme/internal/model"
	"github.com/fajaralfa/askme/internal/repo"
	"github.com/fajaralfa/askme/internal/service/jwt"
	"github.com/gorilla/mux"
)

type User struct {
	UserRepo repo.UserInterface
}

func (h *User) FindByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user, err := h.UserRepo.FindByEmail(vars["email"])
	if err != nil {
		ApiInternalErr(w, err)
		return
	}

	if user == nil {
		ApiErr(w, "fail", "pengguna tidak ditemukan", nil, 404)
		return
	}

	resp := model.Response{
		Status: "success",
		Data: map[string]model.User{
			"user": {
				Id:    user.Id,
				Email: user.Email,
				Photo: user.Photo,
			},
		},
	}

	helper.WriteJSON(w, resp)
}

func (h *User) CurrentUserInfo(w http.ResponseWriter, r *http.Request) {
	claim, err := jwt.ParseGetClaims(r.Header.Get("Authorization")[len("Bearer "):])
	if err != nil {
		ApiInternalErr(w, err)
		return
	}

	user, err := h.UserRepo.FindByEmail(claim.Email)
	if err != nil {
		ApiInternalErr(w, err)
		return
	}

	user.Password = ""

	resp := model.Response{
		Status: "success",
		Data: map[string]*model.User{
			"user": user,
		},
	}

	helper.WriteJSON(w, resp)
}
