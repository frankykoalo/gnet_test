package main

import (
	_ "github.com/go-sql-driver/mysql"
	"gnet_test/database"
	ghttp "gnet_test/http"
)

func main() {
	ss := database.ConnectDatabase()
	ghttp.HttpHandler(*ss)
}
