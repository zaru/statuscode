package model

import (
	"database/sql"
	"log"
)

var db = connection()

func connection() *sql.DB {
	db, err := sql.Open("mysql", "root:@/statuscode")
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close() これクローズしなくていいのか？　参考 https://github.com/haluanice/social_network_golang/blob/92fcaa28c6e1a9489edf7dbb8f1b8627fb413c8a/src/service/ExecutionDB.go

	return db
}

func Create(fqdn string) {
	stmtIns, err := db.Prepare("insert sites (fqdn) values(?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(fqdn)
	if err != nil {
		panic(err.Error())
	}
}

func Select() []string {
	rows, err := db.Query("select fqdn from sites")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var fqdns []string
	for rows.Next() {
		var fqdn string
		if err := rows.Scan(&fqdn); err != nil {
			panic(err.Error())
		}
		fqdns = append(fqdns, fqdn)
	}

	return fqdns
}
