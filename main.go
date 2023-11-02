package main

import (
	"fmt"
	"github.com/lirongyangtao/mygo/base"
	"strconv"
)

type Score int

func (s Score) Less(element base.SkipElement) bool {
	return s < element.(Score)
}

func (s Score) Key() string {
	return strconv.Itoa(int(s))
}

func main() {
	arr := []Score{99, 100, 1, 2, 3, 454, 6, 7, 8}
	s := base.NewSortSet[Score]()
	for _, v := range arr {
		fmt.Println(s.Add(v))
	}
	for _, v := range arr {
		fmt.Println(s.Add(v))
	}
	for _, v := range arr {
		fmt.Println(v, s.GetRank(v.Key()))
	}
	fmt.Println("======================")
	for i := 0; i <= len(arr)+3; i++ {
		fmt.Println(i, s.GetElementRank(i))
	}

}
