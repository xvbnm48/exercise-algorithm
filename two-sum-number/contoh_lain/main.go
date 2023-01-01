package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Buat linked list dengan 3 node
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}

	// Tampilkan isi dari linked list
	fmt.Println("Isi linked list:")
	node := head
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
	fmt.Println()

	// Tambahkan node baru dengan nilai 4 di awal linked list
	newNode := &ListNode{Val: 4}
	newNode.Next = head
	head = newNode

	// Tampilkan isi dari linked list setelah node baru ditambahkan
	fmt.Println("Isi linked list setelah node baru ditambahkan:")
	node = head
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
	fmt.Println()
}
