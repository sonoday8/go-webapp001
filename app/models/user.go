package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	appContext "github.com/sonoday8/webapp001/app/context"
	"golang.org/x/crypto/bcrypt"
	"time"
	"unsafe"
)

// User return user
type User struct {
	//gorm.Model
	ID            int64 `gorm:"primary_key"`
	LoginID       string
	Password      string
	RememberToken string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// TestUser return TestUser
func GetAllUsers(c echo.Context) ([]User, error) {
	dbc := c.(*appContext.DBContext)
	var allUsers []User
	err := dbc.DBConn(func(db *gorm.DB) error {
		db.Find(&allUsers)
		return nil
	})
	if err != nil {
		return allUsers, nil
	}
	return allUsers, nil
}

func ExistsUser(c echo.Context, user User) (bool, error) {
	dbc := c.(*appContext.DBContext)
	if err := dbc.DBConn(func(db *gorm.DB) error {
		db.Where(&User{LoginID: user.LoginID}).First(&user)
		return nil
	}); err == nil && user.ID != 0 {
		return true, nil
	}
	return false, nil
}

func CreateUser(c echo.Context, user User) (bool, error) {
	dbc := c.(*appContext.DBContext)
	if err := dbc.DBTran(func(db *gorm.DB) error {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
		if err != nil {
			c.Logger().Error(err)
			return err
		}
		fmtHash := *(*string)(unsafe.Pointer(&hash))
		user.Password = fmtHash
		if res := db.Create(&user); res.Error != nil {
			return res.Error
		}
		return nil
	}); err == nil && user.ID != 0 {
		return true, nil
	}
	return false, nil
}
