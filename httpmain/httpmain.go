package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gnet_test/dollars"
	ghttp "gnet_test/http"
	"gnet_test/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	databaseConn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root",
		"localhost:3306", "test")
	db, err := sqlx.Connect("mysql", databaseConn)
	defer db.Close()
	if err != nil {
		log.Fatalf(" can't open database ,%s , system will quit ... ", err.Error())
		os.Exit(1)
	}
	ss := dollars.NewSqlServer(db)

	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		d := dollars.ListServer(*ss)
		j, err := json.MarshalIndent(d, "", "  ")
		if err != nil {
			log.Fatalf("Marshal failed,%v\n ", err)
		}
		fmt.Fprintf(w, string(j))
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Printf(r.Method)
		body, err := ioutil.ReadAll(r.Body)
		var a model.Dollar
		if err = json.Unmarshal(body, &a); err != nil {
			log.Fatalf("Unmarshal err,%v\n", err)
		}
		dollars.InsertServer(a.Item, a.Price, *ss)
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Printf(r.Method)
		body, err := ioutil.ReadAll(r.Body)
		var a struct {
			Item string `json:"item"`
		}
		if err = json.Unmarshal(body, &a); err != nil {
			log.Fatalf("Unmarshal err,%v\n", err)
		}
		d := dollars.GetDataServer(a.Item, *ss)
		j, _ := json.MarshalIndent(d, "", "  ")
		fmt.Fprintf(w, string(j))
	})

	ghttp.ListenAndServe(":8080", nil)
}
