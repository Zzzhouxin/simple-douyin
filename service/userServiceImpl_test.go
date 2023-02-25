package service

import (
	"fmt"
	"github.com/RaymondCode/simple-douyin/dao"
	"testing"
)

func TestGetTableUserByUsername(t *testing.T) {
	dao.Init()
	impl := UserServiceImpl{}
	list := impl.GetTableUserByUsername("user3")
	fmt.Printf("%v\n", list)
}
