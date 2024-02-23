package main

func fizzBuzz(n int32) {
	for i := int32(n); i <= n; i++ {
		divBy3 := i%3 == 0
		divBy5 := i%5 == 0

		if divBy3 && divBy5 {
			println("FizzBuzz")
		} else if divBy3 {
			println("Fizz")
		} else if divBy5 {
			println("Buzz")
		} else {
			println(i)
		}
	}
}

func main() {
	fizzBuzz(21)
}
