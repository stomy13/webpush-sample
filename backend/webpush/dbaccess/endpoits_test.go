package dbaccess

import (
	"testing"
)

// 疎通確認
func Test_ConnectGorm_1(t *testing.T) {
	db := connectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&Endpoint{})
}
