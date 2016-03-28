package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/zaru/statuscode/model"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	_ "github.com/go-sql-driver/mysql"

	//"gopkg.in/gorp.v1"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fqdns := strings.Split(r.Host, ".")
	fmt.Println(fqdns[0])
	fmt.Println(r.Host)
	fmt.Fprintf(w, "Hello, %s! This page is %s host.", c.URLParams["name"], fqdns[0])
	fmt.Println("hoge")
}

func view_mysql(c web.C, w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@/statuscode")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select * from sites")
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Fprintf(w, "%s: %s", columns[i], value)
		}
	}
}

func site_lists(c web.C, w http.ResponseWriter, r *http.Request) {
	fqdns := model.Select()
	for i := range fqdns {
		fmt.Fprintf(w, "%s\n", fqdns[i])
	}
}

func site_create(c web.C, w http.ResponseWriter, r *http.Request) {
	if r.FormValue("fqdn") == "" {
		fmt.Fprintln(w, "error")
	} else {
		model.Create(r.FormValue("fqdn"))
		fmt.Fprintln(w, "created")
	}
}

func main() {

	goji.Get("/api/v1/sites", site_lists)
	goji.Post("/api/v1/sites", site_create)

	goji.Get("/hello/:name", hello)
	goji.Get("/mysql", view_mysql)
	goji.Serve()
}
