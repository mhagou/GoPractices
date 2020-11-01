package main

import "fmt"

type node struct {
	next  *node
	value int
}

func Newnode() *node {
	return new(node)
}

func (n *node) Getvalue() int {
	return n.value
}

func (n *node) Setvalue(v int) {
	n.value = v
}

type llist struct {
	firstnode *node
	lastnode  *node
}

func Newllist() *llist {
	return new(llist)
}

func (ll *llist) add(v int) {
	ele := &node{value: v}
	if ll.firstnode == nil {
		ll.firstnode = ele
		ll.lastnode = ele
	} else {
		ll.lastnode.next = ele
		ll.lastnode = ele
	}
}

func main() {

	l := Newllist()
	l.add(33)
	l.add(66)
	l.add(99)
	l.add(111)
	l.add(222)
	l.add(333)
	l.add(444)
	l.add(555)

	n := l.firstnode
	for n.next != nil {
		fmt.Println(n.value)
		n = n.next
	}
	fmt.Println(n.value)
}
