package dbaccess

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ConnectArgs struct {
	Address  string
	Port     string
	DBName   string
	User     string
	Password string
}

func ConnectGorm(conargs *ConnectArgs) *gorm.DB {
	connectTemplate := "%s:%s@tcp(%s:%s)/%s?parseTime=true"
	connect := fmt.Sprintf(connectTemplate, conargs.User, conargs.Password, conargs.Address, conargs.Port, conargs.DBName)
	log.Println(connect)
	db, err := gorm.Open("mysql", connect)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return db
}
