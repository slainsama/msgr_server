package globals

import (
	"fmt"
	"github.com/slainsama/msgr_server/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func initDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		UnmarshaledConfig.DB.User,
		UnmarshaledConfig.DB.Pass,
		UnmarshaledConfig.DB.Host,
		UnmarshaledConfig.DB.Port,
		UnmarshaledConfig.DB.Name,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Println(err)
	}
}
