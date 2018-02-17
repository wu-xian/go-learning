package main

import (
	"fmt"
)

type human interface {
	sleep() bool
	jerkoff() bool
	showSelf()
}

type wuxian struct {
	name string
	cnt  int
}

func (wu wuxian) sleep() bool {
	return wu.cnt > 10
}

func (wu wuxian) jerkoff() bool {
	return wu.name == "smoker"
}

func (wu wuxian) showSelf() {
	fmt.Println("in show self", wu.name)
	fmt.Println("in show self", wu.cnt)
}

func showItByInterface(hu human) {
	//fmt.Println("in sub function", wu.name)
	//fmt.Println("in sub function", wu.cnt)
	//fmt.Println("in sub function", )
	hu.showSelf()
	fmt.Println("jerkoff:", hu.jerkoff())
}

func main5() {
	fmt.Println("interface")
	var w1 = wuxian{name: "wuxian", cnt: 9}
	var w2 = wuxian{name: "smoker", cnt: 11}
	result1 := w1.sleep()
	fmt.Println("result1", result1)
	result2 := w2.sleep()
	fmt.Println("result2", result2)
	showItByInterface(w1)
}
