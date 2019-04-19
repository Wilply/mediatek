package main

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gchaincl/dotsql"
	_ "github.com/mattn/go-sqlite3"
)

type database struct {
	db      *sql.DB
	queries *dotsql.DotSql
}

var errQuery = errors.New("db: error while querying database")
var errAddUser = errors.New("db: cannot create user")

func newDB() database {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	dot, err := dotsql.LoadFromFile("./form/rqst.sql")
	if err != nil {
		log.Fatal(err)
	}
	_, err = dot.Exec(db, "create-users-table")
	if err != nil {
		log.Fatal(err)
	}
	_, err = dot.Exec(db, "create-groups-table")
	if err != nil {
		log.Fatal(err)
	}
	return database{
		db:      db,
		queries: dot,
	}
}

func (d database) init() {
	ok, _ := d.getgroupbyname("root")
	if !ok {
		d.addgroup(group{
			name:      "root",
			readonly:  []string{"/"},
			readwrite: []string{"/"},
		})
	}
}

func (d database) close() {
	d.db.Close()
}

func (d database) test() {
	d.adduser(user{
		name:      "root",
		pass:      "root",
		active:    true,
		groups:    []string{"1"},
		readonly:  []string{"/temp/ro"},
		readwrite: []string{"/temp/rw"},
	})
	fmt.Println(d.userlist())
	//fmt.Println(d.getuserbyid(2))
	d.addgroup(group{
		name:      "root",
		readonly:  []string{"/tmp", "/ro"},
		readwrite: []string{"/tmp", "/rw"},
	})
	fmt.Println(d.grouplist())
	//fmt.Println(d.getgroupbyname("root"))
}
