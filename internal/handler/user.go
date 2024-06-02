package handler

import (
	"net/http"

	"github.com/fajaralfa/askme/internal/helper"
	"github.com/fajaralfa/askme/internal/model"
	"github.com/fajaralfa/askme/internal/repo"
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
