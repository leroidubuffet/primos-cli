package main

import "testing"

func TestIsPrime(t *testing.T) {
	// Test prime numbers
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	for _, prime := range primes {
		if !isPrime(prime) {
			t.Errorf("isPrime(%d) = false; want true", prime)
		}
	}
	
	// Test non-prime numbers
	nonPrimes := []int{0, 1, 4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20}
	for _, nonPrime := range nonPrimes {
		if isPrime(nonPrime) {
			t.Errorf("isPrime(%d) = true; want false", nonPrime)
		}
	}
	
	// Test negative numbers
	negatives := []int{-10, -1, -100, -999}
	for _, negative := range negatives {
		if isPrime(negative) {
			t.Errorf("isPrime(%d) = true; want false", negative)
		}
	}
	
	// Test edge cases
	if isPrime(0) {
		t.Errorf("isPrime(0) = true; want false")
	}
	if isPrime(1) {
		t.Errorf("isPrime(1) = true; want false")
	}
	if !isPrime(7919) {
		t.Errorf("isPrime(7919) = false; want true")
	}
	if isPrime(7920) {
		t.Errorf("isPrime(7920) = true; want false")
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrime(7919)
	}
}