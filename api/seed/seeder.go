package seed

import (
	"log"
	"time"
	
	"github.com/Muhammad-Tounsi/Vote-Go-Vue/api/models"
	"github.com/jinzhu/gorm"
)

// Load : Validation and Join
func Load(db *gorm.DB) {

	input1 := "1996-02-08"
	layout1 := "2006-01-02"
	t, _ := time.Parse(layout1, input1)

	input := "2018-10-01"
	layout := "2006-01-02"
	start, _ := time.Parse(layout, input)

	input2 := "2020-10-01"
	layout2 := "2006-01-02"
	end, _ := time.Parse(layout2, input2)

	var users = []models.User{
		models.User{
			Firstname:   "Steven",
			Lastname:    "Victor",
			Email:       "steven@gmail.com",
			Accesslevel: 1,
			Password:    "password",
			Dateofbirth: t,
		},
		models.User{
			Firstname:   "Kevin",
			Lastname:    "Feige",
			Email:       "feige@gmail.com",
			Accesslevel: 1,
			Password:    "password",
			Dateofbirth: t,
		},
		models.User{
			Firstname:   "Karim",
			Lastname:    "Benzema",
			Email:       "k@benzema.io",
			Accesslevel: 1,
			Password:"Mostafa87",
			Dateofbirth: t,
		},
	}

	var votes = []models.Vote{
		models.Vote{
			Title:     "Title 1",
			Desc:      "Hello world 1",
			AuthorID : 1,
			StartDate: start,
			EndDate:   end,
		},
		models.Vote{
			Title:     "Propreté des trottoirs",
			Desc:      "Dans le budget qui sera soumis au vote des conseillers de Paris lundi, 32 M€ seront consacrés à l’achat de nouvelles machines, plus silencieuses, plus propres et plus performantes. Il n’y aura pas d’embauche d’agents supplémentaires.",
			AuthorID : 1,
			StartDate: start,
			EndDate:   end,
		},
		models.Vote{
			Title:     "Title 2",
			Desc:      "Hello world 2",
			AuthorID : 1,
			StartDate: start,
			EndDate:   end,
		},
	}

	err := db.Debug().DropTableIfExists(&models.Vote{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Vote{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Vote{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		votes[i].AuthorID = users[i].ID
		err = db.Debug().Model(&models.Vote{}).Create(&votes[i]).Error
		if err != nil {
			log.Fatalf("cannot seed votes table: %v", err)
		}
	}
	db.Model(&votes[0]).Association("Users")
	err = db.Debug().Model(&votes[0]).Association("Users").Append(&users).Error
	if err != nil {
		log.Fatalf("Association error: %v", err)
	}
		

}
