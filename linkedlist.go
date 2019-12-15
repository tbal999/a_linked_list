package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	index int
	item  interface{}
	next  *Node
	prev  *Node
}

func (l *Node) add(str string) {
	node := *l
	newnode := Node{}
	switch node.index == 0 {
	case true: //we are at index 0
		if node.next == nil { //next node doesn't exist
			node.item = str
			node.next = &newnode
			newnode.prev = &node
			newnode.index = node.index + 1
			*l = node
			return
		}
		if node.next != nil {
			node.next.add(str)
		}
	case false: //we are not at index 0
		if node.next == nil { //next node doesn't exist
			node.item = str
			node.next = &newnode
			newnode.prev = &node
			newnode.index = node.index + 1
			*l = node
			return
		}
		if node.next != nil { //next node does exist
			node.next.add(str)
		}

	}
	*l = node
}

func (l *Node) delete(index int) {
	node := *l
	if index == 0 && node.index == 0 && node.next == nil {
		fmt.Println("Nothing to delete")
		return
	}
	switch index == 0 && node.index == 0 {
	case true: // If we are at the first node and index to be deleted is 0.
		fmt.Println("first node")
		fmt.Println(node)
		fmt.Println(node.next)
		node = *node.next
		node.index = 0
		*l = node
		if node.next != nil {
			node.next.indexshift(index + 1)
		}
	case false: // If we are not at the first node
		switch node.index == index {
		case true:
			switch node.next.next == nil {
			case true: // If we are at the last node and need to delete
				fmt.Println("last node")
				fmt.Println(node)
				node.next.prev = node.prev
				node = *node.next
				node.index = index
				*l = node
				return
			case false: // If we are not at the last node and need to delete
				fmt.Println("a middle node")
				fmt.Println(node)
				node.next.prev = node.prev
				node = *node.next
				node.index = index
				node.next.index = index + 1
				*l = node
				return
			}
		case false:
			switch node.next == nil {
			case true: // If we are at the last node and we haven't found the index
				fmt.Printf("The index %d doesn't exist in list\n", index)
				return
			case false: // if we haven't found the index but not at last node
				node.next.delete(index)
			}
		}
	}
}

func (l *Node) indexshift(index int) {
	node := *l
	if node.next != nil {
		node.index = index
		*l = node
		node.next.indexshift(index + 1)
	}
	*l = node
}

func (l *Node) view() {
	node := *l
	if node.next != nil {
		fmt.Println(node.index, node.item)
		node.next.view()
	}
}

func (l *Node) reverseview(iterator, index int) {
	node := *l
	if node.index == index {
		fmt.Println(node.index, node.item)
		node.prev.reverseview2(index - 1)
		return
	}
	if node.index == iterator {
		node.next.reverseview(iterator+1, index)
	}
}

func (l *Node) reverseview2(index int) {
	node := l
	if node.index == index && node.index != 0 {
		fmt.Println(node.index, node.item)
		node.prev.reverseview2(index - 1)
	}
	if node.index == 0 {
		fmt.Println(node.index, node.item)
		return
	}
}

func main() {
	fmt.Println("Start: add, delete, print, quit.")
	List := Node{}
	Scanner := bufio.NewScanner(os.Stdin)
reset:
	Scanner.Scan()
	answer := Scanner.Text()

	if answer == "" {
		goto reset
	}
	result := strings.Fields(answer)
	text := strings.Join(result[1:], " ")
	switch result[0] {
	case "add":
		if len(result) == 1 {
			fmt.Println("You need to type in some text after 'add'.")
			goto reset
		}
		List.add(text)
		List.view()
		goto reset
	case "delete":
		if len(result) == 1 {
			fmt.Println("You need to type in an index (number) with delete.")
			goto reset
		}
		number, err := strconv.Atoi(result[1])
		if err != nil {
			fmt.Println("You need to type in an index (number) with delete.")
			goto reset
		}
		List.delete(number)
		List.view()
		goto reset
	case "print":
		List.view()
		goto reset
	case "rprint":
		if len(result) == 1 {
			fmt.Println("You need to type in an index (number) with rprint.")
			goto reset
		}
		number, err := strconv.Atoi(result[1])
		if err != nil {
			fmt.Println("You need to type in an index (number) with rprint.")
			goto reset
		}
		List.reverseview(0, number)
		goto reset
	case "quit":
		goto ending
	}
	fmt.Println("invalid input")
	goto reset
ending:
	fmt.Println("Bye!")
}
