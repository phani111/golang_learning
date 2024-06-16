package main

import (
	"13-UtilsPackage/stringset"
	"13-UtilsPackage/utils"
)

func main() {
	s := utils.NewStringSet()
	_ = utils.SortStringSet(s) //name of package doesnt tell us any thing

	s2 := stringset.New() //Better
	_ = s2.Sort()
}
