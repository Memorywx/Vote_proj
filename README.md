# Vote_proj

### 数据库表设置
1.user 存放用户的账号密码与id

2.vote 存放具体的投票项目，例如“今天晚上吃什么”,
以及投票项目id,与多选或单选类型。

3.vote_opt 进入具体的投票项目后，可以选择的选项，比如
“红烧肉”， 以及对应的id还有它们被投的次数。

4.vote_opt_user 存在登录的用户的id，以及他们进入哪个
投票项目和投了什么选项的id。

### logic文件夹放着的是 
响应投票页面和登录页面的函数

### model文件夹放着的是 
1.db.go 连接和关闭mysql数据库

2.model.go 存放着mysql建表语句转换为gorm语句的结构体

3.user.go 存放着从数据库获取用户数据的函数

4.vote.go 存放着获取数据库中的投票项目的函数和增加表vote_user_opt
的函数


## router文件夹放着的是
所以的 url 以及检查 cookie 的功能

## tools文件夹放着的是
返回状态的json，message，和code

## view文件夹放着的是
登录页面的html文件
投票项目的html文件
进入投票项目后选择选项的html文件

## app.go
连接数据库，启动路由

## main.go 
直接启动app.go 间接启动整个运行

