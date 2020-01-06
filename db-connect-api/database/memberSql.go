package database

import (
	"fmt"
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

// GetMemberByID = This method will return a list of members
func (db *DB) GetMemberByID(id int, t string) *m.Member {
	fmt.Println(id)
	// perform a db.Query insert
	stmt, _ := db.Prepare(`
	SELECT id, username, type 
	FROM member.members M
	INNER JOIN member.credits C ON C.memberId = M.id
	WHERE M.id = ? AND (? = '' or M.type = ?)
	 LIMIT 200`)

	// var parameterMap = map[string]interface{}{
	// 	"id":   id,
	// 	"type": t,
	// }

	// stmt.SetValuesFromMap(parameterMap)

	rsMemberData, err := stmt.Query(id, t, t)
	defer rsMemberData.Close()

	// if there is an error inserting, handle it
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	var member *m.Member
	for rsMemberData.Next() {
		var m m.Member
		err := rsMemberData.Scan(&m.ID, &m.Username, &m.Type)
		if err != nil {
			panic(err.Error())
		}
		member = &m
		fmt.Println(member)
	}
	return member
}
