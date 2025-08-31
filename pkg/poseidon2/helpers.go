package poseidon2

import (
	"fmt"
	"math/big"
)

// used to convert a hex value to a bigInt for use in
func convertHex(i string, p *big.Int) *big.Int {
	z, isValid := new(big.Int).SetString(i, 0)
	if !isValid {
		panic("bad hex: " + i)
	}
	z.Mod(z, p)
	return z
}

// used to chunk the round constants
func chunkRoundConstants(allConsts []string, p *big.Int, t int, Rf int, Rp int) ([][]*big.Int, []*big.Int) {
	if len(allConsts) != (Rf+Rp)*t {
		panic("Bad constant length")
	}

	// break into rounds
	rounds := make([][]*big.Int, 0, Rf+Rp)
	for i := 0; i < len(allConsts); i += t {
		// build thte row
		currRow := make([]*big.Int, 0, t)
		for j := 0; j < t; j++ {
			// one-by-one get the values and convert from hex to bigInt
			val := convertHex(allConsts[i+j], p)
			// append to the temporary row
			currRow = append(currRow, val)
		}
		rounds = append(rounds, currRow)
	}

	// break the external rounds into proper rows
	// Get the first set of external rounds
	externalRounds := make([][]*big.Int, 0, Rf)
	for i := 0; i < Rf/2; i++ {
		externalRounds = append(externalRounds, rounds[i])
	}

	// get the second set of external rounds which are after the internal
	for i := (Rf / 2) + Rp; i < (Rf/2)+Rp+(Rf/2); i++ {
		externalRounds = append(externalRounds, rounds[i])
	}

	// break out the internal rounds extracting only position 0 for each
	internalRounds := make([]*big.Int, 0, Rp)
	for i := Rf / 2; i < (Rf/2)+Rp; i++ {
		internalRounds = append(internalRounds, rounds[i][0]) // only needs to extract lane 0 for each partial round
	}

	return externalRounds, internalRounds
}

func formatMatrix(m [][]string, p *big.Int, t int) [][]*big.Int {
	fmt.Println("length of m: ", len(m))
	if len(m) != t {
		panic("Malformed Matrix, doesn't match size TxT")
	}

	newMatrix := make([][]*big.Int, 0, len(m))
	for i := 0; i < len(m); i++ {
		newRow := make([]*big.Int, 0, len(m[i]))
		for j := 0; j < len(m[i]); j++ {
			out := convertHex(m[i][j], p)
			newRow = append(newRow, out)
		}
		newMatrix = append(newMatrix, newRow)
	}

	return newMatrix
}

func generateInternalMatrix(t int, p *big.Int) [][]*big.Int {
	output := make([][]*big.Int, t)
	for i := 0; i < t; i++ {
		row := make([]*big.Int, t)
		for j := 0; j < t; j++ {
			if i == j {
				val := new(big.Int).SetInt64(int64(i + 1))
				row[j] = new(big.Int).Mod(val, p)

			} else {
				row[j] = new(big.Int).SetInt64(0)
			}
		}
		output[i] = row
	}
	return output
}

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
