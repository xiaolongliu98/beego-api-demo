package user

import "bee-api-example1/models/user"

/**
业务逻辑层 - user/user业务
*/

// GetUserById
// @Description 通过用户id获取用户信息
// @Author 		xiaolong
// @Date		2022/11/24 21:26(create);2022/11/25 16:14(update);
// @Param		id		int		用户id
// @Return		*user.User, error
func GetUserById(id int) (*user.User, error) {
	return user.GetUserById(id)
}

// GetUserByUsername
// @Description 通过用户名称获取用户信息
// @Author 		xiaolong
// @Date		2022/11/24 21:26(create);2022/11/25 16:14(update);
// @Param		username		string		用户名称
// @Return		*user.User, error
func GetUserByUsername(username string) (*user.User, error) {
	return user.GetUserByUsername(username)
}
