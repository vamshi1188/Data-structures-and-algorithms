package main

// import "fmt"

func main() {

	// Sumofn(2, 0)
	println(Suminfunction(3))

}

// func Sumofn(i int, n int) {
// 	if i < 1 {
// 		fmt.Println(n)
// 		return
// 	}
// 	Sumofn(i-1, n+i)

// }
func Suminfunction(n int) int {
	if n == 1 {
		return 1
	}
	return n * Suminfunction(n-1)
}
