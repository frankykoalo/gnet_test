package dollars

import (
	"github.com/jmoiron/sqlx"
	"gnet_test/model"
	"log"
)

type SqlServer struct {
	Db *sqlx.DB
}
type Tps struct {
	content []byte
}

func NewSqlServer(db *sqlx.DB) *SqlServer {
	return &SqlServer{
		Db: db,
	}
}

func InsertServer(item string, price float32, ss SqlServer) (err error) {
	_, err = ss.Db.Exec("insert into dollar(item,price,deleted) values (?,?,?)", item, price, 0)
	if err != nil {
		log.Fatalf("%s", err)
	}
	return
}

func ListServer(ss SqlServer) (d []model.Dollar) {
	err := ss.Db.Select(&d, "select id,item,price from dollar where deleted = 0 order by price")
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
