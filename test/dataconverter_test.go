package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestProcessUnit(t *testing.T) {
	databaseConn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root",
		"root",
		"127.0.0.1:3306",
		"explorer")

	db, err := sqlx.Connect("mysql", databaseConn)
	if err != nil {
		fmt.Printf("------------- err : %s", err.Error())
		t.Errorf(" can't open database ,%s , system will quit ... ", err.Error())
		//os.Exit(1)
	}

	r, err := db.Exec("insert into sync_state(id, block_height, unit_mci) values(?,?,?)", 1, 2, 3)
	if err != nil {
		t.Errorf(" insert failed .%v", err)
		db.Close()
		return
	}
	id, _ := r.RowsAffected()
	t.Logf(" inser id :%d", id)
	db.Close()
}
