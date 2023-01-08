package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

func CreateJWT(data any) string {

	header, err := json.Marshal(map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	})
	if err != nil {
		fmt.Println("marshal err")
		return ""
	}
	jwtHeader := base64.URLEncoding.EncodeToString(header)

	body, err := json.Marshal(map[string]interface{}{
		"iss": data,
		"exp": "20231216",
		"jti": "100",
	})
	if err != nil {
		fmt.Println("marshal err")
		return ""
	}
	jwtBody := base64.URLEncoding.EncodeToString(body)

	sign := SHA256Secret(jwtHeader + "." + jwtBody)
	return jwtHeader + "." + jwtBody + "." + sign
}

func IsValidJWT(jwt string) bool {
	arr := strings.Split(jwt, ".")
	if len(arr) != 3 {
		return false
	}
	if arr[2] != SHA256Secret(arr[0]+"."+arr[1]) {
		return false
	}
	return true
}
