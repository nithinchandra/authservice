package app

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/nithinchandra/authservice/models"
	"net/http"
	u "github.com/nithinchandra/authservice/utils"
	"os"
	"strings"
)

var JwtAutherntication = func(next http.Handler)http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		notAuth := []string{"/api/user/new", "/api/user/login"} //Endpoints that doesnt require Authentication
		requestPath := r.URL.Path // Request path


		//checking if request needs authentcation
		for _,value := range notAuth{
			if value == requestPath{
				next.ServeHTTP(w,r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization") // Taking token from Authorization header

		// Token missing
		if tokenHeader == ""{
			response = u.Message(false,"Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type","application/json")
			u.Respond(w,response)
			return
		}

		splitted := strings.Split(tokenHeader," ") //The token comes in format `Bearer {token-body}`
		if len(splitted) != 2{
			response = u.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type","application/json")
			u.Respond(w,response)
			return
		}

		tokenPart := splitted[1] //take token part
		tk := &models.Token{}

		token,err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token)(interface{},error){
			return []byte(os.Getenv("token_password")),nil
		})

		if err != nil{ //Malformed token
			response = u.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type","application/json")
			u.Respond(w, response)
			return
		}

		if !token.Valid{ // Token is invalid
			response = u.Message(false,"Token is not valid")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type","application/json")
			u.Respond(w,response)
			return

		}
		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		fmt.Sprintf("User %",tk.UserId) // for monitoring
		ctx := context.WithValue(r.Context(),"user",tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w,r) //proceed in the middleware chain!

	})
}
