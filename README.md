# beego-api-demo

前提
- 我们基于Bee api进行后端开发，与前端分离
- Beego项目文件
├─conf           放置配置文件
├─controllers    请求控制层
│  └─user
├─models         数据模型与访问层
│  ├─db
│  └─user
├─routers        定义路由、绑定Controller、启用filter
├─filters        定义filter规则函数
├─services       业务逻辑层
│  └─user
├─swagger        
├─tests
├─utils          工具
└─main.go        程序入口
- 基于提供的demo基础上进行开发
分层规范
- Controller：请求控制层。负责API请求处理，向下调用Service层功能组织完成，以JSON格式返回数据。
- Service：业务逻辑层。负责相应的业务逻辑实现，向下调用Model层功能组织完成，向上对Controller层提供功能。在该层需注意数据事务的实现。
- Model：数据模型与访问层，与DAO类似。负责数据库表、文件、远程API等与Go-Struct模型映射，并提供模型数据增删改查的基本功能。在该层需注意功能的原子特性。
拒绝跨层访问调用，即使某项功能简单，Model实现后也不允许直接在Controller调用，正确方式是在Service层进行简单封装后，给Controller调用。

具体开发时，应按照具体功能、模型进行文件分解，并在三层尽力做到对应。参照下图，以user业务为例：
[图片]
注释规范
- Controller
// [对应Func名称]
// @Description   [描述Func实现的功能]
// @Author        [作者名称]
// @Date          [xxxx/xx/xx xx:xx](create);[xxxx/xx/xx xx:xx](update)...
// @param         [变量名]       [数据类型]      [解释]
// @Return        [返回数据描述]
// @Router        [beego路由注解]
- Service&Model&utils&...
// [对应Func名称]
// @Description   [描述Func实现的功能]
// @Author        [作者名称]
// @Date          [xxxx/xx/xx xx:xx](create);[xxxx/xx/xx xx:xx](update)...
// @Param         [变量名]       [数据类型]      [解释]
// @Return        [返回数据描述]
- Struct注释
要求在结构体上方注释一行，并对每个属性做注释。
例如：
// User 用户对象，定义了用户的基础信息
type User struct{
    Id        int    // 用户Id
    Username  string // 用户名
    Email     string // 邮箱
}
配置文件规范
要求配置文件统一放置于./conf下，一般都建议以JSON文件格式进行配置文件编写，例如下方DB的配置：
1. 准备配置文件
[图片]
2. 在models/db中编写对应配置Model：
[图片]
3. 并提供绑定方法（使用JSON工具以及utils工具）：
[图片]
4. 若需要启动注册、或者全局变量注册（初始化），请在models/register.go下完成，或者创建相应的models/XXXregister.go来编写注册方法，并在models/register.go的RegisterAll方法中调用，具体参考models/dbregister.go
