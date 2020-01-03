package controller

import (
	db "member-db-api/database"
)

var (
	//MemberDB is the DB Context
	memberContext = func() *db.DB { return db.GetDB(db.MemberDBName) }
)
