package dal

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// go get github.com/go-sql-driver/mysql

type MySQLConfig struct {
	User     string
	Password string
	Address  string
	DBName   string
}

func NewMySQLDB(conf MySQLConfig) (*sql.DB, error) {
	dbConf := mysql.NewConfig()
	dbConf.User = conf.User
	dbConf.Passwd = conf.Password
	dbConf.Addr = conf.Address
	dbConf.DBName = conf.DBName

	connector, err := mysql.NewConnector(dbConf)
	if err != nil {
		return nil, fmt.Errorf("failed to create MySQL connector: %w", err)
	}

	return sql.OpenDB(connector), nil
}
