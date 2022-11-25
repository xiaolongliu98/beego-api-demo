package user

import "bee-api-demo/services/user"

// GetIntroById
// @Description 通过IntroId获取Intro简介信息
// @Author 		xiaolong
// @Date		2022/11/24 21:26(create);
// @Param		iid		int		简介信息id
// @Return
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
// @Description 通过用户Id获取Intro简介信息
// @Author 		xiaolong
// @Date		2022/11/24 21:26(create);
// @Param		uid		int		用户id
// @Return
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
