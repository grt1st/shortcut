package core

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"strings"
)

var DB *gorm.DB

func initDB() {
	var err error
	parts := strings.SplitN(Config.Mysql, "//", 2)
	DB, err = gorm.Open(parts[0], parts[1])
	if err != nil {
		log.Fatalln("[-] form open database error.", err)
	}

	// log
	//DB.LogMode(true)
}
