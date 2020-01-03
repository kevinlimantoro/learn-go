package controller

import (
	"encoding/json"
	"net/http"

	db "member-db-api/database"
)

// GetAllMember Controller will return all members
var GetAllMember = func(w http.ResponseWriter, r *http.Request) {
	// queries := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")
	data := db.GetDB(db.MemberDBName).GetMembers()
	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "error marshalling data"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
}
