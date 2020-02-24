package dbaccess

import (
	"testing"
)

// 疎通確認
func Test_ConnectGorm_1(t *testing.T) {
	conargs := &ConnectArgs{
		Address:  "localhost",
		Port:     "33306",
		DBName:   "webpush",
		User:     "webpush",
		Password: "webpush"}

	db := ConnectGorm(conargs)
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&Subscription{})
}
