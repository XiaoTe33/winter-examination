package utils

import (
	"crypto/sha256"
	"fmt"
	"regexp"

	"winter-examination/src/dao"
	"winter-examination/src/model"
)

func IsValidUsername(username string) bool {
	return len([]rune(username)) > 0 && len([]rune(username)) <= 20
}
func IsValidEmail(email string) bool {
	return regexp.
		MustCompile(`^(\w+(.\w+)*)*@(\w+(.\w+)+)+$`).
		MatchString(email)
}

func IsValidPhone(phone string) bool {
	return regexp.
		MustCompile("^1[3-9][0-9]{9}$").
		MatchString(phone)
}

func IsValidPassword(password string) bool {
	return regexp.
		MustCompile(`^[\w!-~]{6,20}$`).
		MatchString(password)
}

func IsRegisteredUsername(username string) bool {
	return dao.QueryUserByUsername(username) != model.User{}
}

func IsRegisteredPhone(phone string) bool {
	return dao.QueryUserByPhone(phone) != model.User{}
}

func IsRegisteredEmail(email string) bool {
	return dao.QueryUserByEmail(email) != model.User{}
}

func SHA256Secret(str string) string {
	h := sha256.New()
	h.Write([]byte(str + "secret"))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GetIdByUsername(username string) (id string) {
	return dao.QueryUserByUsername(username).Id
}
