package db

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() (*gorm.DB, error) {
	// conf := configs.GetDB()

	// sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=/%s sslmode=disable",
	// conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	// sqlDB, err := sql.Open("mysql", "mydb_dsn")

	// log.Printf("teste %v", sc)
	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/dea")

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: conn,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}

	// err = gormDB.Ping()

	return gormDB, err
}
