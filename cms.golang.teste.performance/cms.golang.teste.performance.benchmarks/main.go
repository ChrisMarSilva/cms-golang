package main

import (
	"log"
	"math"
    "math/big"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.performance.benchmarks
// go get -u XXXXXXXXXXXX
// go mod tidy

// go run main.go

func main() {

	log.Println("")
	log.Println(" 5 =", primeNumbers(5))
	log.Println(" 5 =", primeNumbers2(5))
	log.Println(" 5 =", isPrimeImproved(5))
	log.Println(" 5 =", isPrimeImproved2(5))

	log.Println("")
	log.Println(" 6 =", primeNumbers(6))
	log.Println(" 6 =", primeNumbers2(6))
	log.Println(" 6 =", isPrimeImproved(6))
	log.Println(" 6 =", isPrimeImproved2(6))
}

func primeNumbers(max int) []int {
	var primes []int

	for i := 2; i < max; i++ {
		isPrime := true

		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}

func primeNumbers2(max int) []int {
    f := make([]bool, max)

    for i := 2; i <= int(math.Sqrt(float64(max))); i++ {
        if f[i] == false {
            for j := i * i; j < max; j += i {
                f[j] = true
            }
        }
    }

	var primes []int
    for i := 2; i < max; i++ {
        if f[i] == false {
			primes = append(primes, i)
        }
    }

	return primes
}

func isPrimeImproved(max int) []int {
	b := make([]bool, max)
	var primes []int

	for i := 2; i < max; i++ {
		if b[i] {
			continue
		}

		primes = append(primes, i)

		for k := i * i; k < max; k += i {
			b[k] = true
		}
	}

	return primes
}

func isPrimeImproved2(max int) []int {
	var primes []int

	for i := 2; i < max; i++ {
        if big.NewInt(int64(i)).ProbablyPrime(20) {
            primes = append(primes, i)
        }
    }

	return primes
}