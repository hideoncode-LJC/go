package connect

import (
	"database/sql"
	"go/goorm/config"
)


func GetDB() *sql.DB{
	
	slc := config.GetSqliteConfig()
	
	db, _ := sql.Open(slc.Dbc, slc.Db)

	defer db.Close()

	return db
}