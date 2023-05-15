package rest

import (
	"awesomeProject2/internal/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	token, err := getTokenFromRequest(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
	}
	id, err := h.usersService.ParseToken(token)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	var inp models.UserUpdate
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		log.Println("updateUser", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.usersService.UpdateUserInfo(int(id), inp)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type User1 struct {
	Name     string
	Email    string
	Password string
}

func (h *Handler) chooseRole(w http.ResponseWriter, r *http.Request) {

	token, err := getTokenFromRequest(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
	}
	id, err := h.usersService.ParseToken(token)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	var inp models.User
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		log.Println("updateUser", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch inp.UserType {
	case "client":
		user_id := id
		err := h.clientService.Create(user_id)
		if err != nil {
			return
		}
	case "owner":
		user_id := id
		err := h.ownerService.Create(user_id)
		if err != nil {
			return
		}
	case "agent":
		var agent models.Agent
		agent.Id = int(id)
		err := h.agentService.Create(agent)
		if err != nil {
			return
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
	err = h.usersService.ChooseRole(int(id), inp.UserType)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
