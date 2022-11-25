package user

import "bee-api-demo/services/user"

/**
用户简介获取
*/

// GetIntroById
// @Title GetUser
// @Author xiaolong
// @Date 2022/11/24 21:26(create);
// @Description 通过IntroId获取Intro简介信息
// @Param iid query int true 用户简介id
// @Success 200 {object} user.UserIntro
// @Failure 500 {string}
// @router /intro/id/:iid [get]
func (u *UserController) GetIntroById() {
	iid, err := u.GetInt(":iid")
	if err != nil {
		u.Data["json"] = err.Error()
		u.ServeJSON()
		return
	}
	intro, err := user.GetIntroById(iid)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = *intro
	}
	u.ServeJSON()
}

// GetIntroByUid
// @Title GetUser
// @Author xiaolong
// @Date 2022/11/24 21:26(create);
// @Description 通过用户Id获取Intro简介信息
// @Param uid query int true 用户id
// @Success 200 {object} user.UserIntro
// @Failure 500 {string}
// @router /intro/uid/:uid [get]
func (u *UserController) GetIntroByUid() {
	uid, err := u.GetInt(":uid")
	if err != nil {
		u.Data["json"] = err.Error()
		u.ServeJSON()
		return
	}
	intro, err := user.GetIntroByUid(uid)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = *intro
	}
	u.ServeJSON()
}
