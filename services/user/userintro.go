package user

import "bee-api-demo/models/user"

/**
业务逻辑层 - user/userintro业务
*/

// GetIntroById
// @Description 通过IntroId获取Intro简介信息
// @Author 		xiaolong
// @Date		2022/11/24 21:26(create);
// @Param		iid		int		简介信息id
// @Return      *user.UserIntro, error
func GetIntroById(id int) (*user.UserIntro, error) {
	return user.GetIntroById(id)
}

// GetIntroByUid
// @Description 通过用户Id获取Intro简介信息
// @Author 		xiaolong
// @Date		2022/11/24 21:26(create);
// @Param		uid		int		用户id
// @Return		*user.UserIntro, error
func GetIntroByUid(uid int) (*user.UserIntro, error) {
	return user.GetIntroByUid(uid)
}
