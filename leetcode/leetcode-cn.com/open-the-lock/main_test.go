package main

import(
	"testing"
	"fmt"
)

func Test_next(t *testing.T){
	ret := next("0000")

	fmt.Println(ret)
}
