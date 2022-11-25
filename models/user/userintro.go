package user

import (
	"github.com/astaxie/beego/orm"
)

// 用户简介信息
type UserIntro struct {
	Id           int    `json:"id,omitempty" orm:"description(...);pk;auto"`   // 简介id
	Introduction string `json:"introduction,omitempty" orm:"description(...)"` // 简介内容
	UserId       int    `json:"user_id,omitempty" orm:"description(...)"`      // 用户id
}

// 指定UserIntro结构体默认绑定的表名
func (u *UserIntro) TableName() string {
	return "user_intro_table"
}

// 以下提供一些基础功能

// GetIntroById
// @Description 通过IntroId获取Intro简介信息
// @Author 		xiaolong
// @Date		2022/11/24 21:26(create);
// @Param		iid		int		简介信息id
// @Return      *user.UserIntro, error
func GetIntroById(id int) (*UserIntro, error) {
	o := orm.NewOrm()
	// -------------------------------------
	ui := &UserIntro{Id: id}

	err := o.Read(ui) // select * from user_intro_table where id=?
	if err != nil {
		//fmt.Printf("%v\n", err.Error())
		return nil, err
	}

	return ui, nil
}

// GetIntroByUid
// @Description 通过用户Id获取Intro简介信息
// @Author 		xiaolong
// @Date		2022/11/24 21:26(create);
// @Param		uid		int		用户id
// @Return		*user.UserIntro, error
func GetIntroByUid(uid int) (*UserIntro, error) {
	o := orm.NewOrm()
	// -------------------------------------
	ui := &UserIntro{UserId: uid}

	err := o.Read(ui, "user_id") // select * from user_intro_table where user_id=?
	if err != nil {
		//fmt.Printf("%v\n", err.Error())
		return nil, err
	}

	return ui, nil
}
