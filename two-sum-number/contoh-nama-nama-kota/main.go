// program linked list menyimpan data barang

package main

import "fmt"

type Barang struct {
	Nama  string
	Harga int
}

type Node struct {
	Barang Barang
	Next   *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) TambahBarang(b Barang) {
	node := &Node{Barang: b}
	if l.Head == nil {
		l.Head = node
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
}

func (l *LinkedList) TampilBarang() {
	current := l.Head
	for current != nil {
		fmt.Println(current.Barang)
		current = current.Next
	}
}

func main() {
	var list LinkedList
	list.TambahBarang(Barang{"Beras", 10000})
	list.TambahBarang(Barang{"Gula", 15000})
	list.TampilBarang()
}
