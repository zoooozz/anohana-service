package common

import (
	"anohana-service/model"
	"crypto/md5"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"regexp"
	"strings"
	"time"
)

func PhoneLegitimate(phone string) bool {

	lens := strings.Count(phone, "") - 1
	if lens > 11 {
		return false
	}
	reg := regexp.MustCompile(`1[3|5|7|8|][\d]{9}`)
	return reg.MatchString(phone)
}

func TokenProduction(phone string) (token string) {
	r := jwt.New(jwt.SigningMethodHS256)
	claims := r.Claims.(jwt.MapClaims)
	claims["phone"] = phone
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token, _ = r.SignedString([]byte(model.JwtSalt))
	return
}

func PassProduction(phone, password string) (pass string) {
	pass = fmt.Sprintf("%x", md5.Sum([]byte(password+model.PassSalt+phone)))
	return
}
