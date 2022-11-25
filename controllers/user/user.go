package user

import (
	"bee-api-demo/services/user"
)

/**
完成用户注册登录功能
*/

// GetUserById
// @Title GetUser
// @Author xiaolong
// @Date 2022/11/24 21:26(create);
// @Description 通过用户id获取用户信息
// @Param uid query int true 用户id
// @Success 200 {object} user.User
// @Failure 500 {string}
// @router /id/:uid [get]
func (u *UserController) GetUserById() {
	uid, err := u.GetInt(":uid")
	if err != nil {
		u.Data["json"] = err.Error()
		u.ServeJSON()
		return
	}

	o, err := user.GetUserById(uid)

	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = *o
	}
	u.ServeJSON()
}

// GetUserByUsername
// @Title GetUser
// @Author xiaolong
// @Date 2022/11/24 21:26(create);
// @Description 通过用户名称获取用户信息
// @Param username query string true 用户名称
// @Success 200 {object} user.User
// @Failure 500 {string}
// @router /username/:username [get]
func (u *UserController) GetUserByUsername() {
	username := u.GetString(":username")
	if username == "" {
		u.Data["json"] = "username is empty"
		u.ServeJSON()
		return
	}

	o, err := user.GetUserByUsername(username)

	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = *o
	}
	u.ServeJSON()
}
