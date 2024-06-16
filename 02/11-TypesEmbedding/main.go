package main

import (
	"11-TypesEmbedding/bad"
	"11-TypesEmbedding/good"
)

func main() {
	i := bad.NewInMem()
	i.Lock() // Bad, user shouldn't have access to mutex outside struct
	i.Add("a", 1)
	i.Add("b", 2)
	i.Add("c", 3)
	i.Unlock()

	i2 := good.NewInMem()
	i2.Add("a", 1)
	i2.Add("b", 2)
	i2.Add("c", 3)
}
