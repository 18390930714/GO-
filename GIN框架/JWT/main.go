package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	username string `json:"username"`
	jwt.StandardClaims
}

func main() {

	mySigningKey := []byte("woshizzzzlovego")

	c := MyClaims{
		username: "zzzz",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() - 60*60*2,
			Issuer:    "zzzz",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	s, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)

}
