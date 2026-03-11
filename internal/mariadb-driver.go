package internal

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v3/log"
)

func NewDatabase(user string, password string, host string, port string, dbname string) *sql.DB {
	var intPort, errorToConvert = strconv.Atoi(port)

	if errorToConvert != nil {
		log.Fatal("Error converting port to int")
	}

	var sqlInfo = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, intPort, dbname)

	conn, err := sql.Open("mysql", sqlInfo)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Debugf("(CONFIG) Database Drive Initialized - %s", dbname)

	return conn
}
