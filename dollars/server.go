package dollars

import (
	"gnet_test/model"
	"log"
)

func InsertServer(item string, price float32, ss SqlServer) {
	stmt, err := ss.InsertData()
	if err != nil {
		log.Fatalf("%s", err)
	}
	_, err = stmt.Exec(item, price, 0)
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func ListServer(ss SqlServer) (d []model.Dollar) {
	err := ss.Db.Select(&d, "select item,price from dollar where deleted = 0")
	if err != nil {
		log.Fatalf("%s", err)
	}
	return
}
