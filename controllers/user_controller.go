package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/egnimos/mvc/services"
	"github.com/egnimos/mvc/utils"
	"net/http"
	"strconv"
)

//controllers handle the requests from the client
func GetUsers(w http.ResponseWriter, req *http.Request) {
	userId := req.FormValue("userId")
	//convert the string into int
	id, err := strconv.Atoi(userId)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "UserId must be a number",
			Status: http.StatusBadRequest,
			Code: "bad_request",
		}
		jsonErr, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.Status)
		w.Write(jsonErr)
		return
	}
	//getting the user of that particular ID
	user, apiErr := services.GetUser(id)
	if apiErr != nil {
		//send the bad request
		jsonErr, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.Status)
		w.Write(jsonErr)
		return
	}
	//print the result
	fmt.Println(user)

	jsonUser, _ := json.Marshal(user)
	fmt.Fprintln(w,"USER JSON=> ", string(jsonUser))
}

