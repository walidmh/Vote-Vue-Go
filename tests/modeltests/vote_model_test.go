package modeltests

import (
	"log"
	"testing"
	"time"

	"github.com/Muhammad-Tounsi/Vote-Go-Vue/api/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/go-playground/assert.v1"
)

/*
func TestFindAllVotes(t *testing.T) {

	err := refreshUserAndVoteTable()
	if err != nil {
		log.Fatalf("Error refreshing user and vote table %v\n", err)
	}
	_, _, err = seedUsersAndVotes()
	if err != nil {
		log.Fatalf("Error seeding user and vote table %v\n", err)
	}
	votes, err := voteInstance.FindAllVotes(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the votes: %v\n", err)
		return
	}
	assert.Equal(t, len(votes), 2)
}*/

func TestSaveVote(t *testing.T) {

	err := refreshUserAndVoteTable()
	if err != nil {
		log.Fatalf("Error user and vote refreshing table %v\n", err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user %v\n", err)
	}

	input := "2018-10-01"
	layout := "2006-01-02"
	start, _ := time.Parse(layout, input)

	input2 := "2020-10-01"
	layout2 := "2006-01-02"
	end, _ := time.Parse(layout2, input2)

	newVote := models.Vote{
		ID:        1,
		Title:     "This is the title",
		Desc:      "This is the content",
		AuthorID:  user.ID,
		StartDate: start,
		EndDate:   end,
	}
	savedVote, err := newVote.SaveVote(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the vote: %v\n", err)
		return
	}
	assert.Equal(t, newVote.ID, savedVote.ID)
	assert.Equal(t, newVote.Title, savedVote.Title)
	assert.Equal(t, newVote.Desc, savedVote.Desc)
	assert.Equal(t, newVote.AuthorID, savedVote.AuthorID)

}

func TestGetVoteByID(t *testing.T) {

	err := refreshUserAndVoteTable()
	if err != nil {
		log.Fatalf("Error refreshing user and vote table: %v\n", err)
	}
	vote, err := seedOneUserAndOneVote()
	if err != nil {
		log.Fatalf("Error Seeding table")
	}
	foundVote, err := voteInstance.FindVoteByID(server.DB, vote.ID)
	if err != nil {
		t.Errorf("this is the error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, foundVote.ID, vote.ID)
	assert.Equal(t, foundVote.Title, vote.Title)
	assert.Equal(t, foundVote.Desc, vote.Desc)
}

func TestUpdateAVote(t *testing.T) {

	err := refreshUserAndVoteTable()
	if err != nil {
		log.Fatalf("Error refreshing user and vote table: %v\n", err)
	}
	vote, err := seedOneUserAndOneVote()
	if err != nil {
		log.Fatalf("Error Seeding table")
	}

	input := "2018-10-01"
	layout := "2006-01-02"
	start, _ := time.Parse(layout, input)

	input2 := "2020-10-01"
	layout2 := "2006-01-02"
	end, _ := time.Parse(layout2, input2)

	voteUpdate := models.Vote{
		ID:        1,
		Title:     "modiUpdate",
		Desc:      "modiupdate@gmail.com",
		AuthorID:  vote.AuthorID,
		StartDate: start,
		EndDate:   end,
	}
	updatedVote, err := voteUpdate.UpdateAVote(server.DB, vote.ID)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, updatedVote.ID, voteUpdate.ID)
	assert.Equal(t, updatedVote.Title, voteUpdate.Title)
	assert.Equal(t, updatedVote.Desc, voteUpdate.Desc)
	assert.Equal(t, updatedVote.AuthorID, voteUpdate.AuthorID)
}

func TestDeleteAVote(t *testing.T) {

	err := refreshUserAndVoteTable()
	if err != nil {
		log.Fatalf("Error refreshing user and vote table: %v\n", err)
	}
	vote, err := seedOneUserAndOneVote()
	if err != nil {
		log.Fatalf("Error Seeding tables")
	}
	isDeleted, err := voteInstance.DeleteAVote(server.DB, vote.ID, vote.AuthorID)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	//one shows that the record has been deleted or:
	// assert.Equal(t, int(isDeleted), 1)

	//Can be done this way too
	assert.Equal(t, isDeleted, int64(1))
}
