package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/wawansoer/V-Go-JWT/backend/models"
)

var users = []models.User{
	models.User{
		Nickname: "Sutisna Utamasjida Septiawan",
		Email:    "wawansoer@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Wawan",
		Email:    "wawansoer@hotmail.com",
		Password: "password",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

}
