package models

import (
	"bee-api-demo/models/db"
	"bee-api-demo/models/user"
	"fmt"
	"github.com/astaxie/beego/orm"
)

// registerDB
// @Description 绑定DB配置，并向ORM注册
// @Author 		xiaolong
// @Date		2022/11/24 21:26(create);
// @Param		configFile		string		db-json配置文件路径
// @Return      error
func registerDB(configFile string) error {
	// read database config
	dbConfig, err := db.NewDBConfigAndBind(configFile)
	if err != nil {
		return err
	}

	// register database into orm
	err = orm.RegisterDataBase(dbConfig.Alias, dbConfig.Driver, dbConfig.GetDataSource())
	if err != nil {
		return err
	}

	return nil
}

// registerModels
// @Description 向ORM注册Models，请将所有需要数据库表映射的Model在此注册
// @Author 		xiaolong
// @Date		2022/11/24 21:26(create);
func registerModels() {
	orm.RegisterModel(new(user.User))
	orm.RegisterModel(new(user.UserIntro))
}

func registerORM() {

	orm.Debug = true
	err := registerDB("./conf/database-dev-config.json")
	if err != nil {
		panic(fmt.Sprintf("%v\n", err.Error()))
		return
	}
	registerModels()
}
