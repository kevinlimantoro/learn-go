package controller

import (
	"fmt"
	u "member-db-api/util"
	"net/http"
)

// GetAllMember Controller will return all members
var GetAllMember = func(w http.ResponseWriter, r *http.Request) {
	rawtoken := r.Context().Value("ckiftk")
	mytoken, _ := rawtoken.(map[string]interface{})
	fmt.Println("FROM HERE I GET", mytoken["message"])
	data := memberContext().GetMembers()
	if data != nil {
		message := u.BaseSuccessResponse
		message["data"] = data
		u.Respond(w, message)
	} else {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.BaseFailedResponse)
	}
}
