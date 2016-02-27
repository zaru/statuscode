package main

import (
  "fmt"
  "net/http"
  "strings"

  "github.com/zenazn/goji"
  "github.com/zenazn/goji/web"

  //"gopkg.in/gorp.v1"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
  fqdns := strings.Split(r.Host, ".")
  fmt.Println(fqdns[0])
  fmt.Println(r.Host)
  fmt.Fprintf(w, "Hello, %s! This page is %s host.", c.URLParams["name"], fqdns[0])
}

func main() {
  //db, err := sql.Open("mymysql", "tcp:localhost:3306*statuscode/root/")
  //dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
  //t1 := dbmap.AddTableWithName(Site{}, "sites").SetKeys(true, "Id")


  goji.Get("/hello/:name", hello)
  goji.Serve()
}

type Site struct {
    Id      int64  `db:"site_id"`
    Created int64
    Fqdn    string `db:",size:512"`
}
