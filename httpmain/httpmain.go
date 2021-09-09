package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gnet_test/dollars"
	ghttp "gnet_test/http"
	"log"
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
	ghttp.HttpHandler(*ss)
	ghttp.ListenAndServe(":10232", nil)
}
