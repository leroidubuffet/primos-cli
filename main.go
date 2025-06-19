package main

import (
	"fmt"
	"os"
	"strconv"
)

// isPrime checks if a number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	
	limit := n / 2
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// showHelp displays usage information
func showHelp() {
	fmt.Println("Uso: go run main.go <numero1> [numero2] [numero3] ...")
	fmt.Println("Ejemplo: go run main.go 17 23 100")
	fmt.Println("Verifica si los números dados son primos")
}

func main() {
	args := os.Args[1:] // Skip program name
	
	if len(args) == 0 {
		fmt.Println("Error: Se requiere al menos un número como argumento")
		showHelp()
		os.Exit(1)
	}
	
	// Check for help flags
	if args[0] == "-h" || args[0] == "--help" || args[0] == "help" {
		showHelp()
		os.Exit(0)
	}
	
	// Process each argument
	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("Error: '%s' no es un número válido\n", arg)
			os.Exit(1)
		}
		
		if isPrime(num) {
			fmt.Printf("%d es primo\n", num)
		} else {
			fmt.Printf("%d no es primo\n", num)
		}
	}
}