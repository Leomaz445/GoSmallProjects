package main

import "fmt"

type user struct {
	name string
	age  int
}

type node struct {
	val  user
	next *node
}

func printMyList(listToPrint *node) {
	var p *node = listToPrint
	for counter := 1; p != nil; counter++ {
		fmt.Printf("The %d user name is: %v, and the user age is: %d\n",
			counter, p.val.name, p.val.age)
		p = p.next
	}
}
func addNode(head *node, newNode node) {
	if head == nil {
		head = &newNode

	} else {
		var p *node = head
		for head.next != nil {
			head = head.next
		}

		head.next = &newNode
		head = p
	}
}
func main() {
	var head *node
	var user1 node
	var user2 node
	var user3 node
	head = &user1
	user1.val = user{name: "Leonid", age: 32}
	user2.val = user{name: "Shlomi", age: 40}
	user3.val = user{name: "Michael", age: 55}
	addNode(head, user1)
	addNode(head, user2)
	addNode(head, user3)
	printMyList(head)
	fmt.Printf("%v", head)

}
