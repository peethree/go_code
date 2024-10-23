package main

import (
	"fmt"	
)

func printPrimes(max int) {		
	for n := 2; n <= max; n++ {		
		isPrime := true
			
		if n == 2 {
			fmt.Println(n)
		}

		if n % 2 == 0 {
			isPrime = false
			continue
		}

		for i := 3; i * i <= n; i+=2 {		
			if n % i == 0 {
				isPrime = false
				continue
			}					
		}

		if isPrime == true {
			fmt.Println(n)
		}
	}	
}



func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
