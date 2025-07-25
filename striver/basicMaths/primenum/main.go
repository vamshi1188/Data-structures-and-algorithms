package main

import "fmt"

func main() {
	num := 10
	countPrimes(num)
}
func countPrimes(n int) {

	count := 0

	for i := 1; i <= n; i++ {

		if n%i == 0 {
			count++
		}

		// if count == 2 {
		// 	fmt.Println("its a prime number")
		// 	break
		// }
	}
	fmt.Println(count)

}
