package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetDb() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=Sn4ck3d dbname=snacked sslmode=disable password=Sn4ck3d host=db")

	if err != nil {
		panic(err)
	}

	return db
}
