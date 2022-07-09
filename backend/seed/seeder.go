package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/wawansoer/V-Go-JWT/backend//models"
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

	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
