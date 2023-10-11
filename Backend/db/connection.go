package db

import (
	"api/configs"
	"database/sql"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

func openDatabase() (*gorm.DB, error) {
	conf := configs.GetDB()

	conn, err := sql.Open("mysql", conf.User+":"+conf.Pass+"@tcp("+conf.Host+":"+conf.Port+")/"+conf.Database)
	if err != nil {
		return nil, err
	}

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
		return nil, err
	}

	return gormDB, nil
}

// GetDB retorna uma única instância de conexão com o banco de dados.
func GetDB() (*gorm.DB, error) {
	var err error
	dbOnce.Do(func() {
		db, err = openDatabase()
	})
	return db, err
}
