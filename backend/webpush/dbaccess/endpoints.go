package dbaccess

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// type conn struct {
// 	db *gorm.DB
// }

// func (c conn) Insert(model interface{}) {
// }

type Endpoint struct {
	gorm.Model
	UserID   int    `gorm:"size:11"`
	Endpoint string `gorm:"size:2048"`
	Key      string `gorm:"size:255"`
	Token    string `gorm:"size:255"`
}

func ConnectGorm() *gorm.DB {
	connectTemplate := "%s:%s@%s/%s"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)
	log.Println(connect)
	db, err := gorm.Open(Dialect, connect)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return db
}

const (
	// データベース
	Dialect = "mysql"

	// ユーザー名
	DBUser = "webpush"

	// パスワード
	DBPass = "webpush"

	// プロトコル
	DBProtocol = "tcp(localhost:33306)"

	// DB名
	DBName = "webpush"
)
