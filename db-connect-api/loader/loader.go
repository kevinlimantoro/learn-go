package loader

import (
	"member-db-api/database"
)

var accessableDBs = []string{database.MemberDBName}

func InitAllContext() {
	for _, item := range accessableDBs {
		database.InitDB(item)
	}
}
