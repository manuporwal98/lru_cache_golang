package main

import (
	"fmt"
)

const SIZE = 5 //signifies the size of hash map or queue
type Node struct {
	Val   string
	Left  *Node //pointer which point to a struct , that is Node
	Right *Node
}

type Queue struct {
	Head   *Node //Head of type node
	Tail   *Node
	Length int //length of type int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

type Hash map[string]*Node

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{} //head points to the first elemnt of the queue and tail points to the last element , basically they do not have a val , if head and tail points to each other, that means queue is empty.
	tail := &Node{} //Does it mean var head is a pointer  //head has the address of empty node ///think when we declare a linked list in cpp, we say ListNode* head , meaning head is of type ListNode8 , always used the val of a node is used to approach pther elements
	head.Right = tail
	tail.Left = head
	return Queue{Head: head, Tail: tail}
}

func (c *Cache) Check(str string) {
	node := &Node{}
	if val, ok := c.Hash[str]; ok { //this is thr way to see if a element exists in hash
		node = c.Remove(val) //need to define remove and add function
	} else {
		node = &Node{Val: str}
	}
	c.Add(node)
	c.Hash[str] = node
}
func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove: %s\n", n.Val)
	left := n.Left
	right := n.Right
	left.Right = right
	right.Left = left
	c.Queue.Length -= 1
	delete(c.Hash, n.Val)
	return n //returning the n which is removed as it needs to be added back
}

func (c *Cache) Add(n *Node) {
	fmt.Printf("add: %s\n", n.Val)
	tmp := c.Queue.Head.Right //basically the first element of queue
	c.Queue.Head.Right = n    //the new first node will be n
	n.Left = c.Queue.Head     //basically this is the empty pointer or node before first node
	n.Right = tmp
	tmp.Left = n
	c.Queue.Length++
	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}

}
func (c *Cache) Display() {
	c.Queue.Display()
}
func (q *Queue) Display() {
	node := q.Head.Right //queue's first val
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Val)
		if i < q.Length-1 {
			fmt.Printf("<->")
		}
		node = node.Right
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("START CACHE")
	cache := NewCache()
	letters := []string{"parrot", "carrot", "water", "tree", "onion", "milk"}
	for _, word := range letters { //_ is for index(as we don't need index, word is for the exact value in slice)
		cache.Check(word)
		cache.Display()
	}
}
