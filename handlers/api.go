package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/arthurkay/lucid/utils"
)

// ApiHomeHandler returns the home route api endpoint
func ApiHomeHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "success"}
	resp, err := json.Marshal(response)
	utils.CheckError(err)
	w.Write(resp)
}

// ApiRegisterHandler creates an user account
func ApiRegisterHandler(w http.ResponseWriter, r *http.Request) {}

// ApiResetPasswordHandler trigger a password reset request
func ApiResetPasswordHandler(w http.ResponseWriter, r *http.Request) {}

// ApiNewPasswordHandler create a new user password
func ApiNewPasswordHandler(w http.ResponseWriter, r *http.Request) {}

// ApiLoginHandler handles all login requests via API calls
func ApiLoginHandler(w http.ResponseWriter, r *http.Request) {}

// ApiServiceRequest handles all API request and routes them accordingly
func ApiServiceRequest(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"status":  "success",
		"message": "Middleware auth success",
	}

	data, err := json.Marshal(resp)
	utils.CheckError(err)
	w.Write(data)
}
