package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/dgrijalva/jwt-go"
	"log"
	"strconv"
	"time"
)

type UserServiceImpl struct {
}

func (usi *UserServiceImpl) GetTableUserByUsername(name string) dao.TableUser {
	tableUser, err := dao.GetTableUserByUsername(name)
	if err != nil {
		log.Println("Err:", err.Error())
		log.Println("User Not Found")
		return tableUser
	}
	log.Printf("Query User Success!")
	return tableUser
}

func GenerateToken(username string) string {
	u := UserService.GetTableUserByUsername(new(UserServiceImpl), username)
	fmt.Printf("generatetoken: %v\n", u)
	token := NewToken(u)
	println(token)
	return token
}

func NewToken(u dao.TableUser) string {
	expiresTime := time.Now().Unix() + int64(config.OneDayOfHours)
	fmt.Printf("expiresTime: %v\n", expiresTime)
	id64 := u.Id
	fmt.Printf("id: %v\n", strconv.FormatInt(id64, 10))
	claims := jwt.StandardClaims{
		Audience:  u.Name,
		ExpiresAt: expiresTime,
		Id:        strconv.FormatInt(id64, 10),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "tiktok",
		NotBefore: time.Now().Unix(),
		Subject:   "token",
	}
	var jwtSecret = []byte(config.Secret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
		token = "Bearer " + token
		println("generate token success!\n")
		return token
	} else {
		println("generate token fail\n")
		return "fail"
	}
}
