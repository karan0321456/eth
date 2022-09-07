package dto

import ()

type UserRequest struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
	Mnemonic string `json:"mnemonic"`
	Path     string `json:"path"`
}

type UserResponse struct {
	Mnemonic string `json:"mnemonic"`
}

type TransactionRequest struct {
	UserID          string `json:"user_id"`
	TransactionHash string `json:"transaction_hash"`
}
