package server

import (
	"encoding/hex"
	"encoding/json"
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

func InsertServer(content model.Content, ss SqlServer) (err error) {
	jsonByte, _ := json.Marshal(content)
	encodeStr := hex.EncodeToString(jsonByte)
	_, err = ss.Db.Exec("insert into dollar (content,deleted) values(?,0) ", encodeStr)
	return
}

func ListServer(ss SqlServer) (d []model.Dollar) {
	err := ss.Db.Select(&d, "select id,content from dollar where deleted = 0 ")
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

func ListBlock(ss SqlServer) (b []model.Block) {
	err := ss.Db.Select(&b, "select * from block order by block_time desc ")
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	return
}

func ListChainStatus(ss SqlServer) (c []model.ChainStatus) {
	err := ss.Db.Select(&c, "select * from chain_status where id = ?", 1)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	return
}

func ListMessage(ss SqlServer, msgType int) {
	switch msgType {
	case 0:
		ListIndexMessage(ss)
	case 1:
	case 2:
	case 3:
	}
}

func ListIndexMessage(ss SqlServer) (im []model.IndexMessage) {
	return
}
