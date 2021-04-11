package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Env struct {
	Asset interface {
		All() ([]Asset, error)
	}
}

func main() {
	db, err := sql.Open("postgres", "Postgres://")
	if err != nil {
		log.Fatal(err)
	}
	env := &Env{
		assets: AssetModel{DB: db},
	}

	htpp.HandleFunc("/assets", env.AssetIndex)
	http.ListenAndServe(":3000", nil)
}

func (env *Env) AssetIndex(w http.ResponseWriter, r *http.Request) {
	tks, err := model.Globalbooks()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, a := range a {
		fmt.Fprintf(w, "%s, \n, %s, %s, %s", a.Org, a.TID, a.Content, a.Timestamp, a.Signature)
	}
}
