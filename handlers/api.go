package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/arthurkay/lucid-framework/models"
	"github.com/arthurkay/lucid-framework/utils"
)

// ApiHomeHandler returns the home route api endpoint
func ApiHomeHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "success"}
	resp, _ := json.Marshal(response)
	w.Write(resp)
}

// ApiRegisterHandler creates an user account
func ApiRegisterHandler(w http.ResponseWriter, r *http.Request) {}

// ApiResetPasswordHandler trigger a password reset request
func ApiResetPasswordHandler(w http.ResponseWriter, r *http.Request) {}

// ApiNewPasswordHandler create a new user password
func ApiNewPasswordHandler(w http.ResponseWriter, r *http.Request) {}

// ApiLoginHandler handles all login requests via API calls
func ApiLoginHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := ioutil.ReadAll(r.Body)
	utils.CheckError(err)
	var userModel models.User
	jerr := json.Unmarshal(resp, &userModel)
	utils.CheckError(jerr)
	user, er := utils.FindUserByEmail(userModel.Email)
	utils.CheckError(er)

	ok := utils.PasswordHashCompare(userModel.Password, user.Password)

	if ok {
		// Success, now create a jwt token
		token, err := utils.CreateToken()
		utils.CheckError(err)
		responseData := map[string]string{
			"status": "success",
			"token":  token,
		}
		data, er := json.Marshal(responseData)
		utils.CheckError(er)
		w.Write(data)
	} else {
		// Login failure, fail with error
		responseData := map[string]string{
			"status":  "failed",
			"message": "Authentication failure",
		}
		data, err := json.Marshal(responseData)
		utils.CheckError(err)
		w.Write(data)
	}
}

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
