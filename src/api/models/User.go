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
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:255;not null;unique" json:"nickname"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (user *User) BeforeSave() error {
	hashedPassword, err := Hash(user.Password)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return nil
}

func (user *User) Prepare() {
	user.ID = 0
	user.Nickname = html.EscapeString(strings.TrimSpace(user.Nickname))
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}

func (user *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if user.Nickname == "" {
			return errors.New("Required Nickname")
		}

		if user.Password == "" {
			return errors.New("Required Password")
		}

		if user.Email == "" {
			return errors.New("Required email")
		}

		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid email")
		}

		return nil

	case "login":
		if user.Password == "" {
			return errors.New("Required Password")
		}

		if user.Email == "" {
			return errors.New("Required email")
		}

		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid email")
		}

		return nil

	default:
		if user.Nickname == "" {
			return errors.New("Required Nickname")
		}

		if user.Password == "" {
			return errors.New("Required Password")
		}

		if user.Email == "" {
			return errors.New("Required email")
		}

		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid email")
		}

		return nil
	}
}

func (user *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error

	err = db.Debug().Create(&user).Error

	if err != nil {
		return &User{}, err
	}

	return user, nil
} 

func (user *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error

	users := []User{}

	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error

	if err != nil {
		return &[]User{}, err
	}

	return &users, err
}

func (user *User) FindUserById(db *gorm.DB, uid uint32) (*User, error) {
	var err error

	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&user).Error

	if err != nil {
		return &User{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User not found")
	}

	return user, err
}

func (user *User) UpdateUser(db *gorm.DB, uid int32) (*User, error) {

	// hash the password
	err := user.BeforeSave()

	if err != nil {
		log.Fatal(err)
	}

	// proses update data user
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string] interface{} {
			"password": user.Password,
			"nickname": user.Nickname,
			"email": user.Email,
			"updated_at": time.Now(),
		},
	)

	if db.Error != nil {
		return &User{}, db.Error
	}

	// tampilan hasil update user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&user).Error

	if err != nil {
		return &User{}, err
	}

	return user, nil
}

func (user *User) DeleteUser(db *gorm.DB, uid int32) (int64, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
