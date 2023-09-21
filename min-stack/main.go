package main

// type Shape interface {
//     Area() float64
//     Perimeter() float64
// }

// type Circle struct {
//     radius float64
// }

// func (c *Circle) Area() float64 {
//     return 3.14 * c.radius * c.radius
// }

// func (c *Circle) Perimeter() float64 {
//     return 2 * 3.14 * c.radius
// }

type MinStack struct {
	stack []struct {
		val, min int
	}
}

// Constructor adalah konstruktor untuk membuat objek MinStack baru.
func Constructor() MinStack {
	return MinStack{}
}

// Push digunakan untuk menambahkan elemen baru ke dalam tumpukan.
// Fungsi ini memperbarui nilai minimum saat ini jika diperlukan.
func (ms *MinStack) Push(val int) {
	min := val
	// Memeriksa jika tumpukan tidak kosong dan val lebih besar dari nilai minimum saat ini.
	if len(ms.stack) > 0 && val > ms.GetMin() {
		min = ms.GetMin()
	}
	// Menambahkan elemen baru ke dalam tumpukan dengan nilai dan nilai minimum saat ini yang sesuai.
	ms.stack = append(ms.stack, struct{ val, min int }{val, min})
}

// Pop digunakan untuk menghapus elemen teratas dari tumpukan.
func (ms *MinStack) Pop() {
	if len(ms.stack) > 0 {
		ms.stack = ms.stack[:len(ms.stack)-1]
	}
}

// Top digunakan untuk mengambil nilai elemen teratas dari tumpukan tanpa menghapusnya.
func (ms *MinStack) Top() int {
	if len(ms.stack) > 0 {
		return ms.stack[len(ms.stack)-1].val
	}
	// Handle kasus ketika tumpukan kosong (return 0 atau nilai default lainnya).
	return 0
}

// GetMin digunakan untuk mengambil nilai minimum saat ini dalam tumpukan.
// Nilai ini selalu merepresentasikan elemen terkecil dalam tumpukan saat ini.
func (ms *MinStack) GetMin() int {
	if len(ms.stack) > 0 {
		return ms.stack[len(ms.stack)-1].min
	}
	// Handle kasus ketika tumpukan kosong (return 0 atau nilai default lainnya).
	return 0
}

func main() {
	minStack := Constructor()
	minStack.GetMin()
}
