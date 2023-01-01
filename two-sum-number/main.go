package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Buat variabel untuk menyimpan hasil penjumlahan dan carry over
	sum := 0
	carry := 0

	// Buat variabel untuk menyimpan pointer ke node baru yang akan dibuat
	var result *ListNode

	// Buat variabel untuk menyimpan pointer ke node saat ini dari linked list hasil
	var current *ListNode

	// Selama salah satu dari linked list masih memiliki node, lakukan penjumlahan
	for l1 != nil || l2 != nil {
		// Hitung nilai node saat ini dengan menambahkan nilai node dari kedua linked list dan carry over
		sum = carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// Hitung carry over untuk operasi berikutnya
		carry = sum / 10

		// Buat node baru dengan nilai digit saat ini dan tambahkan ke linked list hasil
		newNode := &ListNode{Val: sum % 10}
		if result == nil {
			result = newNode
			current = newNode
		} else {
			current.Next = newNode
			current = newNode
		}
	}

	// Jika terdapat carry over setelah penjumlahan selesai, tambahkan node baru dengan nilai carry over
	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	// Kembalikan pointer ke node pertama dari linked list hasil
	return result
}

func main() {
	// Buat linked list untuk menyimpan angka 243
	l1 := &ListNode{Val: 3}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 2}

	// Buat linked list untuk menyimpan angka 564
	l2 := &ListNode{Val: 4}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 5}

	// Tambahkan kedua linked list dan tampilkan hasilnya
	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Print(result.Val)
		result = result.Next
	}
}
