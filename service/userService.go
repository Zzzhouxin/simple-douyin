package service

import "github.com/RaymondCode/simple-demo/dao"

type UserService interface {
	// GetTableUserByUsername 获取所有的TableUser对象
	GetTableUserByUsername(name string) dao.TableUser
}
