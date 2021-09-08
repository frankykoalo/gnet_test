package dollars

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type SqlServer struct {
	Db *sqlx.DB
}

func NewSqlServer(db *sqlx.DB) *SqlServer {
	return &SqlServer{
		Db: db,
	}
}

func (ss *SqlServer) InsertData() (stmt *sql.Stmt, err error) {
	stmt, err = ss.Db.Prepare("insert into dollar(item,price,deleted) values (?,?,?) ")
	if err != nil {
		panic(err)
	}
	return
}
