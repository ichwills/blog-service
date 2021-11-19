package model

import (
	"Pro/blog-service/global"
	"Pro/blog-service/pkg/setting"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"fmt"

	"github.com/jinzhu/gorm"
)

type Model struct {
	ID         uint32
	CreatedBy  string
	ModifiedBy string
	CreatedOn  uint32
	ModifiedOn uint32
	DeletedOn  uint32
	IsDel      uint32
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	//FIXME : 最后一个 %s 原项目应该是 %t 但是传递的参数是string 不是bool 占用形参产生报错
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}
