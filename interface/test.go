// 接口实现多态
package main

import "fmt"

type speaker interface {
	speak()
}

type china struct {
	word string
}

func (c *china) speak() {
	fmt.Println(c.word)
}

type african struct {
	word string
}

func (a *african) speak() {
	fmt.Println(a.word)
}

type european struct {
	word string
}

func (e *european) speak() {
	fmt.Println(e.word)
}

// 只要实现了speak方法的类型都可以当做参数
func peopleSay(s speaker) {
	s.speak()
}

// ...变长参数列表 可接受n个T类型的参数
func manyPeopleSay(s ...speaker) {
	for _, p := range s {
		p.speak()
	}
}
