package controller

import (
	u "member-db-api/util"
	"net/http"
)

// GetAllMember Controller will return all members
var GetAllMember = func(w http.ResponseWriter, r *http.Request) {
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
