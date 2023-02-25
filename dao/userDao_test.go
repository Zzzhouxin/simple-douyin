package dao

import (
	"fmt"
	"testing"
)

func TestGetUserTableByUsername(t *testing.T) {
	Init()
	list, err := GetTableUserByUsername("user2")
	fmt.Println("&v", list)
	fmt.Println("&v", err)
}
