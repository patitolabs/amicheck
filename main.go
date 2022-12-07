package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

const MAX int = 1000000

func sumOfProperDivisors(n int) int {
	sum := 0
	for i := 1; i <= n/2+1; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

func amicable(n int, wg *sync.WaitGroup) {
	sum := sumOfProperDivisors(n)
	if sumOfProperDivisors(sum) == n && sum != n {
		fmt.Printf("%d and %d are amicable numbers\n", n, sum)
	}
	wg.Done()
}

func main() {
	var max int
	if len(os.Args) > 1 {
		max, _ = strconv.Atoi(os.Args[1])
	} else {
		fmt.Print("Enter a number: ")
		fmt.Scanf("%d", &max)
		if max > MAX {
			fmt.Println("Number too large exiting...")
			os.Exit(1)
		}
	}

	var wg sync.WaitGroup
	wg.Add(max)
	for i := 0; i < max; i++ {
		go amicable(i, &wg)
	}
	wg.Wait()
}
