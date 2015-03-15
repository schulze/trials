package main

import "fmt"

type yfunc func(yfunc, int) int

func main() {
	fmt.Println(func(n int) int {
		return func(fact yfunc) int {
			return fact(fact, n)
		}(func(f yfunc, k int) int {
			if k == 1 {
				return 1
			}
			return k * f(f, k-1)
		})
	}(10))
}
