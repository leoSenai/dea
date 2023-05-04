package db

import (
	"api/configs"
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() (*gorm.DB, error) {

	conf := configs.GetDB()

	conn, err := sql.Open("mysql", conf.User+":"+conf.Pass+"@tcp("+conf.Host+":"+conf.Port+")/"+conf.Database)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: conn,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
	})

	if err != nil {
		log.Printf("Cannot connect to database named '%s' using the host '%s' on port '%s' with user '%s'", conf.Database, conf.Host, conf.Port, conf.User)
		panic(err)
	}

	return gormDB, err
}
