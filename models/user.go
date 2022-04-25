package models

import (
	"fmt"
	"learngo/goland/yangpanping/bloggin/databases"
	"strconv"
)

//用户信息结构体
type Userinform struct {
	Usid   int
	Usname string
	Uspswd string
	Usroot int
}

//用户set方法

//用户的get方法
func (u *Userinform) Getuser() string {
	return "[" + strconv.Itoa(u.Usid) + "," + u.Usname + "," + u.Uspswd + "," + strconv.Itoa(u.Usroot) + "]"
}

//插入数据
func InsertUser(u Userinform) (code int) {
	re := databases.Db.Select(`Usname`, `Uspswd`, `Usroot`).Create(&u)
	code = int(re.RowsAffected)
	return code
}

//查询数据
func QueryUser() {
	var user Userinform
	databases.Db.First(&user)
	fmt.Println(user)
}

//查询用户，返回权限值
func CheckUser(u Userinform) Userinform {
	var user Userinform
	databases.Db.First(&user, "usname= ?", u.Usname)
	fmt.Println(user)
	return user
}

//更新数据
func UpdataUser() {
	// 条件更新
	res := databases.Db.Model(&Userinform{}).Where("usroot = ?", 1).Update("usroot", 3)
	fmt.Println(res.RowsAffected)
}

//删除数据
func DeleteUser() {
	var user Userinform
	res := databases.Db.Where("usroot=?", 1).Delete(&user)
	fmt.Println(res.RowsAffected)

}
