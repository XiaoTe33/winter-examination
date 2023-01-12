package service

import (
	"winter-examination/src/dao"
	"winter-examination/src/model"
	"winter-examination/src/utils"
)

func Register(username string, password string, password2 string, email string, phone string) string {
	if !utils.IsValidUsername(username) {
		return "用户名需在1~20个字符之间"
	}
	if utils.IsRegisteredUsername(username) {
		return "用户名已存在"
	}
	if !utils.IsValidPassword(password) {
		return "密码由6~20位数字、字母和部分特殊字符组成"
	}
	if password != password2 {
		return "两次密码不一致"
	}
	if !utils.IsValidPhone(phone) {
		return "请输入正确的手机号"
	}
	if utils.IsRegisteredPhone(phone) {
		return "手机号已被注册"
	}
	if !utils.IsValidEmail(email) {
		return "请输入正确的邮箱"
	}
	if utils.IsRegisteredEmail(email) {
		return "邮箱已被注册"
	}

	dao.AddUser(model.User{
		Username: username,
		Password: utils.SHA256Secret(password),
		Phone:    phone,
		Email:    email,
	})
	return "ok"
}

func Login(token string, username string, password string) (msg string) {
	if utils.IsValidJWT(token) {
		return "ok"
	}
	if !utils.IsValidUsername(username) {
		return "用户名需在1~20个字符之间"
	}
	if !utils.IsRegisteredUsername(username) {
		return "用户名不存在"
	}
	if !utils.IsValidPassword(password) {
		return "密码由6~20位数字、字母和部分特殊字符组成"
	}
	if dao.QueryUserByUsername(username).Password != utils.SHA256Secret(password) {
		return "密码不正确"
	}
	return "ok"
}

func QueryUser(token string, id string, username string, phone string, email string) (msg string, user model.User) {
	if token != "" && utils.IsValidJWT(token) {
		username = utils.GetUsernameByToken(token)
	}
	if id != "" {
		if user = dao.QueryUserById(id); user != (model.User{}) {
			return "ok", user
		}
	}
	if username != "" {
		if user = dao.QueryUserByUsername(username); user != (model.User{}) {
			return "ok", user
		}
	}
	if phone != "" {
		if user = dao.QueryUserByPhone(phone); user != (model.User{}) {
			return "ok", user
		}
	}
	if email != "" {
		if user = dao.QueryUserByEmail(email); user != (model.User{}) {
			return "ok", user
		}
	}
	return "not find", model.User{}
}

func UpdateUser(username string, newUsername string, password string, phone string, email string, photo string) (msg string) {
	user := dao.QueryUserByUsername(username)
	if user == (model.User{}) {
		return "user not found"
	}
	if username != "" {
		if !utils.IsValidUsername(newUsername) {
			return "用户名需在1~20个字符之间"
		}
		if utils.IsRegisteredUsername(newUsername) {
			return "用户名已存在"
		}
		user.Username = newUsername
	}
	if password != "" {
		if !utils.IsValidPassword(password) {
			return "密码由6~20位数字、字母和部分特殊字符组成"
		}
		user.Password = utils.SHA256Secret(password)
	}
	if phone != "" {
		if !utils.IsValidPhone(phone) {
			return "请输入正确的手机号"
		}
		if utils.IsRegisteredPhone(phone) {
			return "手机号已被注册"
		}
		user.Phone = phone
	}
	if email != "" {
		if !utils.IsValidEmail(email) {
			return "请输入正确的邮箱"
		}
		if utils.IsRegisteredEmail(email) {
			return "邮箱已被注册"
		}
		user.Email = email
	}
	if photo != "" {
		user.Photo = photo
	}
	dao.UpdateUser(user)
	return "ok"
}

func QueryAllUsers() (msg string, users []model.User) {
	return "ok", dao.QueryAllUsers()
}
