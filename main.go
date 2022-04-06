package main

import (
	"fmt"
	"math"
)

var a, c = 100, -200
var b = "sylar"

func main() {

	fmt.Printf("当前 value 是：%d", a)
	fmt.Printf("Hello,my master：%v", b)
	fmt.Printf("绝对值：%v", math.Abs(float64(c)))

}

// some tools needed to be installed

//  go-outline
//   gotests
//   gomodifytags
//   impl
//   goplay
//   dlv
//   staticcheck
// gopls







