package controllers

import (
	"encoding/json"
	"github.com/nithinchandra/authservice/models"
	u "github.com/nithinchandra/authservice/utils"
	"net/http"
)

var CreateAccount  =  func(w http.ResponseWriter, r *http.Request){

	account := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(account) // decode request body to account struct

	if err != nil{
		u.Respond(w,u.Message(false,"Invalid Request"))
		return
	}

	resp := account.Create()//create account
	u.Respond(w,resp)

}

var Authenticate = func(w http.ResponseWriter, r *http.Request){

	account := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(account) // decode request body to account struct

	if err != nil{
		u.Respond(w, u.Message(false,"Invalid Request"))
		return
	}

	resp := models.Login(account.Email,account.Password)

	u.Respond(w, resp)
}
