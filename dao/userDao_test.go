package dao

import (
	"fmt"
	"testing"
)

func TestGetUserTableByUsername(t *testing.T) {
	Init()
	list, err := GetTableUserByUsername("user2")
	fmt.Println("&v\n", list)
	fmt.Println("&v\n", err)
}

func TestGetTableUserById(t *testing.T) {
	Init()
	list, err := GetTableUserById(int64(2))
	fmt.Printf("%v\n", list)
	fmt.Printf("%v\n", err)
}
