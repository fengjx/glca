package db

import (
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"

	"github.com/fengjx/glca/connom/config"
)

var dbMap = make(map[string]*sqlx.DB)
var defaultDB *sqlx.DB

var toLowerMapper = reflectx.NewMapperFunc("json", strings.ToLower)

func Init() {
	for k, c := range config.GetConfig().DB {
		db, err := sqlx.Open(c.Type, c.Dsn)
		if err != nil {
			log.Panicf("create db connection err - %s, %s, %s", c.Type, c.Dsn, err.Error())
		}
		err = db.Ping()
		if err != nil {
			log.Panicf("db ping err - %s, %s, %s", c.Type, c.Dsn, err.Error())
		}
		if c.MaxIdle != 0 {
			db.SetMaxIdleConns(c.MaxIdle)
		}
		if c.MaxConn != 0 {
			db.SetMaxOpenConns(c.MaxConn)
		}
		db.Mapper = toLowerMapper
		dbMap[k] = db
	}
	defaultDB = dbMap["default"]
}

func GetDefaultDB() *sqlx.DB {
	return defaultDB
}

func GetDB(name string) *sqlx.DB {
	return dbMap[name]
}
