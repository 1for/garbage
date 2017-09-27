package main

import "fmt"

func main(){
	h := NewBinHeap()
	h.Insert(13)
	h.Insert(21)
	h.Insert(16)
	h.Insert(24)
	h.Insert(31)
	h.Insert(19)
	h.Insert(68)
	h.Insert(65)
	h.Insert(26)
	h.Insert(32)
	h.Print();

	fmt.Println("min is:", h.DeleteMin())

	h.Print()
}
