package models

import (
	"errors"
	"html"
	"strings"
	"time"
	"fmt"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Vote : Vote struct
type Vote struct {
	UUID      uuid.UUID `gorm:"type:uuid;unique_index;" json:"uuid"`
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `json:"title" valid:"required"`
	Desc      string    `json:"desc" valid:"required"`
	Author    User      `json:"author"`
	AuthorID  uint64    `gorm:"not null" json:"author_id"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Users []*User `gorm:"many2many:vote_users;association_foreignkey:id;foreignkey:id" json:"users,omitempty"`
}

// TableName : Gorm related
func (v *Vote) TableName() string {
	return "votes"
}

// Validate : validations struct Vote
func (v *Vote) Validate() error {

	if v.Title == "" {
		return errors.New("Required Title")
	}
	if v.Desc == "" {
		return errors.New("Required Content")
	}
	if v.AuthorID < 1 {
		return errors.New( fmt.Sprintf("Required Creator %v", v.AuthorID) ) 
	}
	return nil
}

// BeforeSave : Method before Save
func (v *Vote) BeforeSave(scope *gorm.Scope) error {
	scope.SetColumn("UUID", uuid.NewV4())
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}

// BeforeUpdate is gorm hook that is triggered on every updated on vote struct
func (v *Vote) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

// Prepare : prepare the statements
func (v *Vote) Prepare() {
	v.UUID = uuid.NewV4()
	v.ID = 0
	v.Title = html.EscapeString(strings.TrimSpace(v.Title))
	v.Desc = html.EscapeString(strings.TrimSpace(v.Desc))
	v.Author = User{}
	v.AuthorID = 1
	v.CreatedAt = time.Now()
	v.UpdatedAt = time.Now()
}

// SaveVote : function to save a Vote
func (v *Vote) SaveVote(db *gorm.DB) (*Vote, error) {
	var err error
	err = db.Debug().Model(&Vote{}).Create(&v).Error
	if err != nil {
		return &Vote{}, err
	}
	if v.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", v.AuthorID).Take(&v.Author).Error
		if err != nil {
			return &Vote{}, err
		}
	}
	return v, nil
}

// FindAllVotes : funtion to find all votes
func (v *Vote) FindAllVotes(db *gorm.DB) (*[]Vote, error) {
	var err error
	votes := []Vote{}
	err = db.Debug().Model(&Vote{}).Limit(100).Preload("Users").Find(&votes).Error
	if err != nil {
		return &[]Vote{}, err
	}

	if len(votes) > 0 {
		for i := range votes {
			err := db.Debug().Model(&User{}).Where("id = ?", votes[i].AuthorID).Take(&votes[i].Author).Error
			if err != nil {
				return &[]Vote{}, err
			}
		}
	}
	return &votes, nil
}

// FindVoteByID : function to fin a vote with an ID
func (v *Vote) FindVoteByID(db *gorm.DB, pid uint64) (*Vote, error) {
	var err error
	err = db.Debug().Model(&Vote{}).Where("id = ?", pid).Take(&v).Error
	if err != nil {
		return &Vote{}, err
	}
	if v.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", v.AuthorID).Take(&v.Author).Error
		if err != nil {
			return &Vote{}, err
		}
	}
	return v, nil
}

// UpdateAVote : function to update a vote
func (v *Vote) UpdateAVote(db *gorm.DB, pid uint64) (*Vote, error) {

	var err error
	db = db.Debug().Model(&Vote{}).Where("id = ?", pid).Take(&Vote{}).UpdateColumns(
		map[string]interface{}{
			"title":      v.Title,
			"desc":       v.Desc,
			"updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Vote{}).Where("id = ?", pid).Take(&v).Error
	if err != nil {
		return &Vote{}, err
	}
	if v.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", v.AuthorID).Take(&v.Author).Error
		if err != nil {
			return &Vote{}, err
		}
	}
	return v, nil
}

// DeleteAVote : function to delete a vote
func (v *Vote) DeleteAVote(db *gorm.DB, pid uint64, uid uint64) (int64, error) {

	db = db.Debug().Model(&Vote{}).Where("id = ? and author_id = ?", pid, uid).Take(&Vote{}).Delete(&Vote{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Vote not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
