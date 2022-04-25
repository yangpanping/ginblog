package models

import (
	"fmt"
	"testing"
)

func TestCheckUser(t *testing.T) {
	user := Userinform{
		Usid:   0,
		Usname: "zpp0",
		Uspswd: "12345",
		Usroot: 0,
	}
	x := CheckUser(user)
	fmt.Println(x)
}

func TestInsertUser(t *testing.T) {

	u := Userinform{
		Usid:   202,
		Usname: "ypp1",
		Uspswd: "ypp12345",
		Usroot: 2,
	}
	InsertUser(u)
}
