package controllers

import (
	"Goblog/models"
	"Goblog/utils"
	"encoding/json"
	"net/http"
)

var SignUp = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}
	resp := user.Register() //Create account
	utils.Respond(w, resp)
}

var SignIn = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}
	resp := models.Login(user.Email, user.Password)
	utils.Respond(w, resp)
}