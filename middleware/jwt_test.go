package middleware

import (
	"fmt"
	"testing"
)

func TestGenToken(t *testing.T) {
	str, _ := GenToken("zpp0")
	fmt.Println(str)
}
