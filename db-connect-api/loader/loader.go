package loader

import (
	"member-db-api/database"
)

var accessableDBs = []string{database.MySQLDBName}

func InitAllContext() {
	for _, item := range accessableDBs {
		database.InitDB(item)
	}
}
