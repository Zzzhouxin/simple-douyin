package dao

import "log"

// TableUser 数据库中的User表结构
type TableUser struct {
	Id       int64
	Name     string
	Password string
}

// TableName 修改表名映射
func (tableUser TableUser) TableName() string {
	return "users"
}

func GetTableUserByUsername(name string) (TableUser, error) {
	tabelUser := TableUser{}
	if err := Db.Where("name = ?", name).First(&tabelUser).Error; err != nil {
		log.Printf(err.Error())
		return TableUser{}, err
	}
	return tabelUser, nil
}
