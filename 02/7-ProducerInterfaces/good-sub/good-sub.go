package goodsub

import "fmt"

type sub struct{}

func (s *sub) DoWork() {
	fmt.Println("Do  work good sub")
}
