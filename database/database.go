package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gnet_test/server"
	"log"
	"os"
)

func ConnectDatabase() (ss *server.SqlServer) {
	databaseConn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root",
		"localhost:3306", "test")
	db, err := sqlx.Connect("mysql", databaseConn)
	if err != nil {
		log.Fatalf(" can't open database ,%s , system will quit ... ", err.Error())
		os.Exit(1)
	}
	ss = server.NewSqlServer(db)
	return
}
