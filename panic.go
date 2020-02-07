package main

import (
	"errors"
	"fmt"
	"strings"
)

func a(){
	fmt.Println("-------------------")
}
func b(){
	defer func() {
		if err:=recover();err!=nil {
			err = errors.New("b()内有panic")
		}
	}()
	fmt.Println("--------===========-----------")
	panic("自动触发的panic")
}
func c(){
	fmt.Println("--------+++++++++++-----------")
}
func main()  {
	a()
	b()
	c()
	str :=strings.Replace("hellohellohe","he","go",4)
	fmt.Println(str)
}