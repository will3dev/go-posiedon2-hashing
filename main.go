package main

import (
	"fmt"
	"log"
	"math/big"
	"poseidon2-hashing/pkg/poseidon2"
)

/*

d>=3 that satisfies gcd (d, p-1)

d = 5
p = 2^255


Step 1 - Setup
2^30 < p < 2^255 => p must be a prime
t = valid elements. a single element in range t must be less than p. works best between 2-24

ce = external round constants
ci = internal round constant

generate Me and Mi

Mi must be MDS (maximum distance separable)
B(M) =
this means that B(M) = t + 1

x = input values


ROUNDS CONSIST OF SAME THREE STEPS. ALTERNATE EXTERNAL AND INTERNAL
Step 2 - First external round; add round constants
add a value external constants to each of the values in the input

for internal only add constant to x0 (first position)


Step 3 - Apply S-Box (raise each to d = 5 mod 17)

(x[i]**d)%p

this means go through the current inputs values and raise each to d mod p
these each calculation against a position will replace the current position

Step 4 - Apply External Matric
Multiple each matrix value across the row to the corresponding input value
Add the row then mod p. This row becomes the new xi value

Matrix must be the same length as the input value

Complete the first set of external rounds:
1. Add the constants to each input value and mod p so (x[i]+c[i])mod p....(x[i-1]+c[i-1]) mod p
2. s-box each value in x[] meaning x[i]^d.....x[i-1]^d
3. multiply the matrix by the external linear layer

Complete the partial rounds:
1. Add the constant to just the first input value and mod p while all others stay the same. meaning x[0] + c[0]
2. s-box just value x^d
3. multiply the matrix by the internal linear layer

Complete the last set of external rounds:
1. Add the constants to each input value and mod p so (x[i]+c[i])mod p....(x[i-1]+c[i-1]) mod p
2. s-box each value in x[] meaning x[i]^d.....x[i-1]^d
3. multiply the matrix by the external linear layer




*/

func main() {
	fmt.Println("Poseidon2 Hashing Project")
	x := make([]*big.Int, 3)

	x[0] = big.NewInt(1)
	x[1] = big.NewInt(2)
	x[2] = big.NewInt(3)

	// Your main application logic will go here
	log.Println("Application started successfully")

	output := poseidon2.Poseidon2Hash(x)

	log.Println("Output: ", output)

	log.Println("Application completed successfully")
}

/*

 */

func modularArithmeticAdd(a, b, p int) int {
	return (a + b) % p
}

func modularAritmethicMultiply(a, b, p int) int {
	return (a * b) % p
}

func modularArithmeticPower(a, b, p int) int {
	var result int = 0
	for i := 0; i < b; i++ {
		result = (result * a)
	}

	return result % p
}

func matrixMultiplication(matrix [][]int, input []int, p int) []int {
	result := make([]int, len(input))

	for i := 0; i < len(input); i++ {
		var v int = 0
		for j := 0; j < len(input); j++ {
			v = (v + (matrix[i][j] * input[j]))
		}
		result[i] = v % p
	}

	return result
}

func calculateBranchNumber(matrix [][]int, input []int, p int) int {
	hw := 0  // hammering weight of input
	hwm := 0 // hammering weight of matrix

	// calculate hammering weight for input
	for _, val := range input {
		if val%p != 0 {
			hw++
		}
	}

	// calculate hammering weight for matrix
	result := make([]int, len(input))
	for i := 0; i < len(input); i++ {

		val := 0
		for j := 0; j < len(matrix); j++ {
			val = val + (matrix[i][j] * input[j])
		}
		result[i] = val % p
		if result[i] != 0 {
			hwm++
		}
	}

	return hw + hwm
}
