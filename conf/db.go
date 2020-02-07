package conf

import (
	"Dorm/model"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func db() {
	model.MysqlInit(os.Getenv("MYSQL_DSN"))

}
