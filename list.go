package List

import "fmt"

type Node struct {
	next *Node
	data int
}

type List struct {
	head *Node
	tail *Node
}

func Create() *List {
	new_list := &List{head: nil, tail: nil}
	return new_list
}

func Append(list *List, value int) {
	new_node := &Node{next: nil, data: value}

	if list.head == nil {
		list.head = new_node
		list.tail = new_node
	} else {
		list.tail.next = new_node
		list.tail = new_node
	}
}

func Display(list *List) {
	current := list.head
	fmt.Printf("[")
	for current != nil {
		if current == list.tail {
			fmt.Printf("%d", current.data)
			break
		}
		fmt.Printf("%d,", current.data)
		current = current.next
	}
	fmt.Println("]")
}

func Remove(list *List, index int) int { //caçar erro dps
	if index == 0 {
		value := list.head.data
		list.head = list.head.next
		return value
	}

	count := 0
	current := list.head
	for current != nil {
		if index == count-1 {
			value := current.next.data
			current.next = current.next.next
			return value
		}
		current = current.next
		count++
	}

	fmt.Println("ERRO! INVALID INDEX")
	return -1

}

func Get(list *List, index int) int {
	if index == 0 {
		return list.head.data
	}

	count := 0
	for current := list.head; current != nil; current = current.next {
		if index == count {
			return current.data
		}
		count++
	}
	fmt.Println("Invalid index")
	return -1
}

func Set(list *List, index int, value int) {
	if index == 0 {
		list.head.data = value
	}

	count := 0
	for current := list.head; current != nil; current = current.next {
		if index == count {
			current.data = value
		}
		count++
	}
	fmt.Println("Invalid index")
}

// func Insert (list *List, index int, value int) {
// 	if index == 0 {
// 		new_node := &Node
// 	}
// }
