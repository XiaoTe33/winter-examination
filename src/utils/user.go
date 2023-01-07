package utils

import (
	"fmt"
	"regexp"
)

func IsValidEmail(email string) bool {
	email2 := "test@gmail.com"
	match, err := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`, email2)
	if err != nil {
		fmt.Println(err)
	}
	if match {
		fmt.Println("email is valid")
		return match
	} else {
		fmt.Println("email is invalid")
		return match
	}
	return true
}
