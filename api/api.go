package main

import (
	"encoding/json"
	"net/http"
)

// Coin Balance Params
type CoinBalanceParams struct {
	Username string
}

// Coin Balance response
type CoinBalanceresponse struct {
	// Success Code, usualy 200
	httpCode uint16

	// Acount Balance
	Balance int64
}

type Error struct {
	// error Code
	httpCode uint16

	// error message
	Message string

}

func writeError(w http.ResponseWriter, message string, code uint16) {
	resp := Error{
		httpCode: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(code))

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An unexpected Error occurred", http.StatusInternalServerError)
	}
)