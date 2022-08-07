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
func addNode(head *node, newNode *node) *node {
	if head == nil {
		head = newNode
		return head
	} else {
		var p *node = head
		for head.next != nil {
			head = head.next
		}

		head.next = newNode
		head = p
	}
	return head
}
func deleteByName(head *node, name string) *node {
	var p *node = head
	if p.val.name == name {
		return head.next
	} else {
		var temp *node
		for p.next != nil {
			temp = p
			p = p.next
			if p.val.name == name {
				temp.next = p.next
				p.next = nil
				return head
			}
		}
		temp.next = nil
	}
	return head
}
func main() {
	var head *node
	var user1 node
	var user2 node
	var user3 node
	user1.val = user{name: "Leonid", age: 32}
	user2.val = user{name: "Shlomi", age: 40}
	user3.val = user{name: "Michael", age: 55}
	head = addNode(head, &user1)
	head = addNode(head, &user2)
	head = addNode(head, &user3)
	printMyList(head)
	fmt.Println("after delete test")
	head = deleteByName(head, "Shlomi")
	printMyList(head)
	fmt.Printf("%v", head)

}
