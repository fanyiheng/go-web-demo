package persist

import (
	"errors"
	"fmt"
	"github.com/fanyiheng/go-web-demo/util"
	"github.com/jinzhu/gorm"
)

type Account struct {
	Model
	Name     string `json:"name"`
	Password string `json:"-"`
}

func (a *Account) checkById() error {
	err := db.Select("id").Where("id = ?", a.ID).Take(&Account{}).Error
	if err == gorm.ErrRecordNotFound {
		return errors.New(fmt.Sprintf("账户ID不存在[%d]", a.ID))
	}
	return err
}

func (a *Account) Add() error {
	return db.Create(a).Error
}

func (a *Account) Update(data map[string]interface{}) error {
	if err := a.checkById(); err != nil {
		return err
	}
	return db.Model(a).Updates(data).Error
}

func (a *Account) Delete() error {
	if err := a.checkById(); err != nil {
		return err
	}
	return db.Delete(a).Error
}

func (a *Account) GetById() error {
	if err := db.Take(a).Error; err == gorm.ErrRecordNotFound {
		return errors.New(fmt.Sprintf("账户ID不存在[%d]", a.ID))
	}
	return nil
}

func (a *Account) Find(page *util.Page) error {
	var accounts []Account
	page.Data = &accounts
	dbw := db.Model(a)
	if a.Name != "" {
		dbw.Where("name like ?", "%"+a.Name+"%")
	}
	err := util.PageFind(dbw, page)
	//page.Data = accounts
	return err
}
