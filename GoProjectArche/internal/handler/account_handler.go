package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"GoProjectArche/internal/model"
	"GoProjectArche/internal/service"
	"GoProjectArche/pkg/utils"
)

// handler which communicate with service layer
type AccountHandler struct {
	Service *service.AccountService
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var acc model.Account
	//getting data from user which is in JSON form and converting it to struct
	//JSON->struct(DECODE)
	json.NewDecoder(r.Body).Decode(&acc)

	//communicate with service layer
	err := h.Service.CreateAccount(acc)
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	//if sucess send response
	utils.SuccessResponse(w, "Account created successfully")
}

func (h *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	//getting data from URL like
	//http://localhost:8085/api/user?id=2
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	acc, err := h.Service.GetAccountByID(id)
	if err != nil {
		utils.ErrorResponse(w, "Account not found", http.StatusNotFound)
		return
	}
	//JSON response
	//in utils/response we convert stuct->JSON or anytype
	utils.JSONResponse(w, acc)
}
