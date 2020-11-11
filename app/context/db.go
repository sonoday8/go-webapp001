package context

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/sonoday8/webapp001/app/env"
)

// context.DBContext
type DBContext struct {
	echo.Context
}

func (c *DBContext) DBConn(fn func(*gorm.DB) error) error {
	dsn := getDsn()
	driver := env.GetStr("DB_DRIVER", "mysql")
	db, err := gorm.Open(driver, dsn)
	if err != nil {
		return err
	}
	if err = fn(db); err != nil {
		return err
	}
	if err = db.Close(); err != nil {
		return err
	}
	return nil
}

func (c *DBContext) DBTran(fn func(*gorm.DB) error) error {
	dsn := getDsn()
	driver := env.GetStr("DB_DRIVER", "mysql")
	db, err := gorm.Open(driver, dsn)
	if err != nil {
		return err
	}
	tx := db.Begin()
	if err = tx.Error; err != nil {
		return err
	}
	err = fn(tx)
	defer func() {
		if p := recover(); p != nil {
			c.Logger().Error(p)
			if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
				c.Logger().Error(rollbackErr)
			}
			panic(p)
		} else if err != nil {
			c.Logger().Warn(err)
			if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
				c.Logger().Error(rollbackErr)
				panic(rollbackErr)
			}
		} else {
			err = tx.Commit().Error
		}
	}()
	return err
}

func getDsn() string {
	user := env.GetStr("DB_USERNAME", "developer")
	passwd := env.GetStr("DB_PASSWORD", "secret")
	dbHost := env.GetStr("DB_HOST", "main")
	dbPort := env.GetStr("DB_PORT", "3306")
	dbName := env.GetStr("DB_DATABASE", "main")
	loc := "Asia%2FTokyo"
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s", user, passwd, dbHost, dbPort, dbName, loc)
}
