package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/rizal282/api-golang-mux/src/api/models"
)

var users = []models.User{
	{
		Nickname: "Rizal Ramdani",
		Email: "rizal@gmail.com",
		Password: "1234",
	},
	{
		Nickname: "Rizal Ramdan",
		Email: "rizal2@gmail.com",
		Password: "1234",
	},
}

var posts = []models.Post{
	{
		Title: "Title 1",
		Content: "Test 1",
	},

	{
		Title: "Title 2",
		Content: "Test 2",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error

	if err != nil {
		log.Fatalf("Cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error

	if err != nil {
		log.Fatalf("Cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error

	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("Cannot seed users table: %v", err)
		}

		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error

		if err != nil {
			log.Fatalf("Cannot seed posts table: %v", err)
		}
	}


}