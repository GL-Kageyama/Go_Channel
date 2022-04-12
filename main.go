package main

import (
	"fmt"
)

// Prime Number detection function
func checkPrime(number int) bool {
	for i := 2; i <= number/2; i++ {
		// If divisible & 0 and 1 are not prime numbers
		if number%i == 0 || number == 0 || number == 1 {
			return false
		}
	}
	// Prime numbers except the above
	return true
}

// Function to find prime numbers greater than or equal to a certain number
func primeRoutine(targetLine int, primeChan chan int) {
	i := 0
	for {
		// Increment value
		i++
		// prime number check
		if checkPrime(i) && targetLine < i {
			// Send data to channel
			primeChan <- i // [Key Points]
		}
	}
}

// main function
func main() {

	// Base value of prime number to be sought
	targetLine := 100000

	// Channel Creation
	primeChan := make(chan int)

	// Finding Prime Numbers in Goroutine
	go primeRoutine(targetLine, primeChan)

	// Receive data from channel
	primeResult := <-primeChan // [Key Points]

	// Value Output
	fmt.Println(primeResult)
}

// =================================
//           Output Sample
// =================================
// ~ $ go build -o prime main.go 
// ~ $ ./prime 
// 100003
