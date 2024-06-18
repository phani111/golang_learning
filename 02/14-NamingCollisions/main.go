package main

import (
	"container/list"
)

func main() {
	list := list.New() // From now you cant use list.New due to collision with local variable
	_ = list.PushBack(1)

}
