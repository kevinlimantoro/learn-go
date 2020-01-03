package database

import (
	m "member-db-api/model"

	_ "github.com/go-sql-driver/mysql"
)

// GetMembers = This method will return a list of members
func (db *DB) GetMembers() []m.Member {
	// perform a db.Query insert
	rsMemberData, err := db.Query(`SELECT id, username, type FROM member.members M
	INNER JOIN member.credits C ON C.memberId = M.id
	 LIMIT 200`)
	defer rsMemberData.Close()

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	var memberList = make([]m.Member, 0, 0)
	for rsMemberData.Next() {
		var m m.Member
		err := rsMemberData.Scan(&m.ID, &m.Username, &m.Type)
		if err != nil {
			panic(err.Error())
		}
		memberList = append(memberList, m)
	}
	return memberList
	// be careful deferring Queries if you are using transactions
}
