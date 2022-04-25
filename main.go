package main

import (
	"learngo/goland/yangpanping/bloggin/databases"
	"learngo/goland/yangpanping/bloggin/routers"
)

func main() {

	databases.DbInit()
	routers.GinInit()
	//databases.RedisInit()
	//u :=models.Userinform{
	//	Usid:   202,
	//	Usname: "ypp1",
	//	Uspswd: "ypp12345",
	//	Usroot: 2,
	//}
	//code :=models.InsertUser(u)
	//fmt.Println(code)

}
