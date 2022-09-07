package app

import (
	"net/http"
	"github.com/eth/service"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	serv service.UserService
}

func (u UserHandler) register(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	CustomerPassword := vars["password"]

	mnemonic ,err := u.serv.User(CustomerPassword)
	if err!=nil {
		writeReponse(w,err.Code,err.AsMessage())
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(mnemonic.Mnemonic))
}
