package main

import (
	"fmt"
	"web3-base/base"
)

func main() {
	//data := base.SingleNumber([]int{4, 1, 2, 1, 2})
	//fmt.Println(data)

	data := base.IsValid("(({}{}))")
	fmt.Println(data)
}
