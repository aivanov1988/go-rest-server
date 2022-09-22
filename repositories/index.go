package repositories

import (
	pg "github.com/go-pg/pg/v10"
	//"github.com/go-pg/pg/v10/orm"

	"github.com/aivanov1988/go-rest-server/config"
)

var db *pg.DB

func ConnectToDB(options config.DatabaseConfig) {
	db = pg.Connect(&pg.Options{
		Addr:     options.Addr,
		Database: options.Database,
		User:     options.User,
		Password: options.Password,
	})
}

func CloseConnection() {
	db.Close()
}
