package modeltests

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/Muhammad-Tounsi/Vote-Go-Vue/api/controllers"
	"github.com/Muhammad-Tounsi/Vote-Go-Vue/api/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
)

var server = controllers.Server{}
var userInstance = models.User{}
var voteInstance = models.Vote{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())
}

func Database() {

	var err error

	TestDbDriver := os.Getenv("TestDbDriver")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbUser"), os.Getenv("TestDbName"), os.Getenv("TestDbPassword"))
	server.DB, err = gorm.Open(TestDbDriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", TestDbDriver)
	}
}

func refreshUserTable() error {
	err := server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneUser() (models.User, error) {

	refreshUserTable()

	input1 := "1996-02-08"
	layout1 := "2006-01-02"
	t, _ := time.Parse(layout1, input1)

	user := models.User{
		UUID:        uuid.NewV4(),
		Firstname:   "Pet",
		Lastname:    "Last",
		Email:       "pet@gmail.com",
		Password:    "password",
		Accesslevel: 1,
		Dateofbirth: t,
	}

	err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	return user, nil
}

func seedUsers() error {
	input1 := "1996-02-08"
	layout1 := "2006-01-02"
	t, _ := time.Parse(layout1, input1)

	users := []models.User{
		models.User{
			UUID:        uuid.NewV4(),
			Firstname:   "Steven",
			Lastname:    "Victor",
			Email:       "steven@gmail.com",
			Password:    "password",
			Accesslevel: 1,
			Dateofbirth: t,
		},
		models.User{
			UUID:        uuid.NewV4(),
			Firstname:   "Kenny",
			Lastname:    "Morris",
			Email:       "kenny@gmail.com",
			Password:    "password",
			Accesslevel: 1,
			Dateofbirth: t,
		},
	}

	for i := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func refreshUserAndVoteTable() error {

	err := server.DB.DropTableIfExists(&models.User{}, &models.Vote{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}, &models.Vote{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed tables")
	return nil
}

func seedOneUserAndOneVote() (models.Vote, error) {

	err := refreshUserAndVoteTable()
	if err != nil {
		return models.Vote{}, err
	}

	input1 := "1996-02-08"
	layout1 := "2006-01-02"
	t, _ := time.Parse(layout1, input1)

	user := models.User{
		UUID:        uuid.NewV4(),
		Firstname:   "Sam",
		Lastname:    "Phil",
		Email:       "sam@gmail.com",
		Password:    "password",
		Accesslevel: 1,
		Dateofbirth: t,
	}
	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.Vote{}, err
	}

	input := "2018-10-01"
	layout := "2006-01-02"
	start, _ := time.Parse(layout, input)

	input2 := "2020-10-01"
	layout2 := "2006-01-02"
	end, _ := time.Parse(layout2, input2)

	vote := models.Vote{
		Title:     "This is the title sam",
		Desc:      "This is the content sam",
		AuthorID:  user.ID,
		StartDate: start,
		EndDate:   end,
	}
	err = server.DB.Model(&models.Vote{}).Create(&vote).Error
	if err != nil {
		return models.Vote{}, err
	}
	return vote, nil
}

func seedUsersAndVotes() ([]models.User, []models.Vote, error) {

	var err error

	if err != nil {
		return []models.User{}, []models.Vote{}, err
	}

	input1 := "1996-02-08"
	layout1 := "2006-01-02"
	t, _ := time.Parse(layout1, input1)

	users := []models.User{
		models.User{
			UUID:        uuid.NewV4(),
			Firstname:   "Steven",
			Lastname:    "Victor",
			Email:       "steven@gmail.com",
			Password:    "password",
			Accesslevel: 1,
			Dateofbirth: t,
		},
		models.User{
			UUID:        uuid.NewV4(),
			Firstname:   "Kenny",
			Lastname:    "Morris",
			Email:       "kenny@gmail.com",
			Password:    "password",
			Accesslevel: 1,
			Dateofbirth: t,
		},
	}

	input := "2018-10-01"
	layout := "2006-01-02"
	start, _ := time.Parse(layout, input)

	input2 := "2020-10-01"
	layout2 := "2006-01-02"
	end, _ := time.Parse(layout2, input2)

	votes := []models.Vote{
		models.Vote{
			Title:     "This is the title sam",
			Desc:      "This is the content sam",
			StartDate: start,
			EndDate:   end,
		},
		models.Vote{
			Title:     "This is the title second",
			Desc:      "This is the content second",
			StartDate: start,
			EndDate:   end,
		},
	}

	for i := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		votes[i].AuthorID = users[i].ID

		err = server.DB.Model(&models.Vote{}).Create(&votes[i]).Error
		if err != nil {
			log.Fatalf("cannot seed votes table: %v", err)
		}
	}
	return users, votes, nil
}
