package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:100;not null;unique" json:"username"`
	Password  string    `gorm:"size:100;not null" json:"password"`
	Email     string    `gorm:"size:255;not null;unique" json:"email"`
	Gender    string    `gorm:"size:10;not null" json:"gender"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.Id = 0                                                      // initial id
	u.Username = html.EscapeString(strings.TrimSpace(u.Username)) // trim space for username
	u.Password = html.EscapeString(strings.TrimSpace(u.Password)) // trim space for password
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(command string) error {
	switch strings.ToLower(command) {
	// updated data: required username, password, gender, email
	case "update":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Gender == "" {
			return errors.New("Required Gender")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email Format")
		}

		return nil
	// login data: required username, password
	case "login":
		if u.Username == "" {
			return errors.New("Required Username")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}

		return nil
	// default data : required username, password, gender, email
	default:
		if u.Username == "" {
			return errors.New("Required Username")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Gender == "" {
			return errors.New("Required Gender")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email Format")
		}

		return nil
	}
}

// Create
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	// create & handler error
	err = db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// Search by id
func (u *User) FindUserById(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	// find user by id
	err = db.Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	// handle if not record
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

// Update
func (u *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {
	// hash password before save
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	// update data -> password, email only
	db = db.Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"email":     u.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}

	// Query again -> to show update values
	err = db.Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// Delete
func (u *User) DeleteUser(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
