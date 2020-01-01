package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	age "github.com/bearbin/go-age"
	"github.com/jinzhu/gorm"
	"github.com/badoux/checkmail"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User Struct
type User struct {
	UUID        uuid.UUID `gorm:"type:uuid;unique_index;" json:"uuid"`
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Email       string    `gorm:"size:100;not null;unique" valid:"email" json:"email"`
	Firstname   string    `valid:"required,alpha,length(2|255)" json:"firstname"`
	Lastname    string    `valid:"required,alpha,length(2|255)" json:"lastname"`
	Accesslevel int       `valid:"range(0|1),numeric" json:"access_level"`
	Dateofbirth time.Time `gorm:"not null;" validate:"min=18" json:"date_of_birth"`
	Password    string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName : Gorm related
func (u *User) TableName() string {
	return "users"
}

// Hash Password : simple password hashing method
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword : This method compare the password with the hash
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Validate : Validation rule on the age of the user
// and if the email already exist
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Firstname == "" {
			return errors.New("Required Firstname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if age.Age(u.Dateofbirth) < 18 {
			return errors.New( "Age User needs to be 18+")
		}
		if u.Firstname == "" {
			return errors.New("Required Firstname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

// Prepare : prepare  statements
func (u *User) Prepare() {
	u.UUID = uuid.NewV4()
	u.ID = 0
	u.Firstname = html.EscapeString(strings.TrimSpace(u.Firstname))
	u.Lastname = html.EscapeString(strings.TrimSpace(u.Lastname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// BeforeSave : Method before Save
func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.UUID = uuid.NewV4()
	u.CreatedAt = time.Now()
	return nil
}

// BeforeUpdate is gorm hook that is triggered on every updated on user struct
func (u *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

// SaveUser : Method Save User, triggered on every saved on user struct
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// FindAllUsers : function to find all users
func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

// FindUserByID : function to find a user with an ID
func (u *User) FindUserByID(db *gorm.DB, uid uint64) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

// UpdateAUser : update an user
func (u *User) UpdateAUser(db *gorm.DB, uid uint64) (*User, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"firstname": u.Firstname,
			"lastname":  u.Lastname,
			"email":     u.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// DeleteAUser delete an user
func (u *User) DeleteAUser(db *gorm.DB, uid uint64) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
