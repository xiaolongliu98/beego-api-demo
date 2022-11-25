package user

import (
	"bee-api-demo/services/user"
)

// GetUserById
// @Description 通过用户id获取用户信息
// @Author 		xiaolong
// @Date		2022/11/24 21:26(create);2022/11/25 16:14(update);
// @Param		uid		int		用户id
// @Return
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
// @Description 通过用户名称获取用户信息
// @Author 		xiaolong
// @Date		2022/11/24 21:26(create);
// @Param		username		string		用户名称
// @Return
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
