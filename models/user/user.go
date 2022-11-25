package user

import (
	"github.com/astaxie/beego/orm"
)

// 用户核心信息
type User struct {
	Id       int    `json:"id,omitempty" orm:"description(...);pk;auto"` // 用户id
	Username string `json:"username,omitempty" orm:"description(...)"`   // 用户名
	Password string `json:"password,omitempty" orm:"description(...)"`   // 用户密码
}

// 指定User结构体默认绑定的表名
func (u *User) TableName() string {
	return "user_table"
}

// 以下提供一些基础功能

// GetUserById
// @Description 通过用户id获取用户信息
// @Author 		xiaolong
// @Date		2022/11/24 21:26(create);
// @Param		id		int		用户id
// @Return		*user.User, error
func GetUserById(id int) (*User, error) {
	o := orm.NewOrm()
	// -------------------------------------
	u := &User{Id: id}

	err := o.Read(u) // select * from user_table where id=?
	if err != nil {
		//fmt.Printf("%v\n", err.Error())
		return nil, err
	}

	return u, nil
}

// GetUserByUsername
// @Description 通过用户名称获取用户信息
// @Author 		xiaolong
// @Date		2022/11/24 21:26(create);
// @Param		username		string		用户名称
// @Return		*user.User, error
func GetUserByUsername(username string) (*User, error) {
	o := orm.NewOrm()
	// -------------------------------------
	u := &User{Username: username}

	err := o.Read(u, "username") // select * from user_table where username=?
	if err != nil {
		//fmt.Printf("%v\n", err.Error())
		return nil, err
	}

	return u, err
}
