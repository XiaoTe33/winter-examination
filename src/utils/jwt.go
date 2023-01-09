package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"winter-examination/src/conf"
)

func CreateJWT(username string) string {

	header, err := json.Marshal(map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	})
	if err != nil {
		fmt.Println("marshal err")
		return ""
	}
	jwtHeader := base64.URLEncoding.EncodeToString(header)

	body, err := json.Marshal(map[string]string{
		"aud": username,
		"exp": strconv.FormatInt(time.Now().Add(time.Second*conf.JWTLastTime).Unix(), 10),
		"nbf": strconv.FormatInt(time.Now().Unix(), 10),
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
	var data = map[string]string{}
	decodeString, err2 := base64.StdEncoding.DecodeString(arr[1])
	if err2 != nil {
		fmt.Println("decodeString failed ...")
	}
	err := json.Unmarshal(decodeString, &data)
	if err != nil {
		fmt.Println("json unmarshal failed...", err)
		return false
	}
	i, err := strconv.ParseInt(data["exp"], 10, 64)
	if err != nil {
		fmt.Println("strconv ParseInt failed ...")
		return false
	}
	return time.Unix(i, 0).After(time.Now())
}
