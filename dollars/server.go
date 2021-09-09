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
	err := ss.Db.Select(&d, "select id,item,price from dollar where deleted = 0")
	if err != nil {
		log.Fatalf("%s", err)
	}
	return
}

func GetDataOfItemServer(item string, ss SqlServer) (d []model.Dollar) {
	err := ss.Db.Select(&d, "select id, item, price from dollar where item = ? and deleted = 0", item)
	if err != nil {
		log.Fatalf("%s", err)
	}
	return
}

func GetDataOfIdServer(id int, ss SqlServer) (d []model.Dollar) {
	err := ss.Db.Select(&d, "select id, item, price from dollar where id = ? and deleted = 0", id)
	if err != nil {
		log.Fatalf("%s", err)
	}
	return
}

func DeleteServer(id int, ss SqlServer) (err error) {
	_, err = ss.Db.Exec("update dollar set deleted = 1 where id = ? and deleted = 0", id)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	return
}

func UpdatePrice(id int, price float32, ss SqlServer) (err error) {
	_, err = ss.Db.Exec("update dollar set price = ? where id = ? and deleted = 0", price, id)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	return
}
