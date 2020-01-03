package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

var (
	// MemberDBName specify member database name in MySQL
	MemberDBName = "member"

	memberDB *sql.DB
)

// DB is a wrapper for sql.DB
type DB struct {
	*sql.DB
}

func getConnString(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("db_username"), os.Getenv("db_password"), os.Getenv("db_address"), os.Getenv("db_port"), dbName)
}

// InitDB this function need to be called first to initialize each DB context
func InitDB(dbName string) {
	fmt.Printf("Start Conencting to %s DB\n", dbName)
	db, err := sql.Open("mysql", getConnString(dbName))

	// defer db.Close()
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	// Set the maximum number of concurrently open connections to 5. Setting this
	// to less than or equal to 0 will mean there is no maximum limit (which
	// is also the default setting).
	db.SetMaxOpenConns(5)
	// Set the maximum number of concurrently idle connections to 5. Setting this
	// to less than or equal to 0 will mean that no idle connections are retained.
	db.SetMaxIdleConns(5)
	// Set the maximum lifetime of a connection to 1 hour. Setting it to 0
	// means that there is no maximum lifetime and the connection is reused
	// forever (which is the default behavior).
	db.SetConnMaxLifetime(time.Hour)
	switch dbName {
	case MemberDBName:
		memberDB = db
	}
}

// GetDB will return db context of a specific DB name
func GetDB(dbName string) *DB {
	var res *DB
	switch dbName {
	case MemberDBName:
		res = &DB{memberDB}
	}
	return res
}
