package poseidon2

import (
	"log"
	"math/big"
	"testing"
)

func TestAddFullRoundConstants(t *testing.T) {
	// Arrange
	p := big.NewInt(17) // Small prime for testing
	x := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)}
	constants := []*big.Int{big.NewInt(5), big.NewInt(7), big.NewInt(11)}

	// Act
	result := addFullRoundConstants(p, x, constants)

	// Assert
	expected := []*big.Int{big.NewInt(6), big.NewInt(9), big.NewInt(14)}
	// (1+5)%17=6, (2+7)%17=9, (3+11)%17=14

	for i, exp := range expected {
		if result[i].Cmp(exp) != 0 {
			t.Errorf("Position %d: expected %v, got %v", i, exp, result[i])
		}
	}
}

func TestInternalMatrix(t *testing.T) {
	p := big.NewInt(17) // Small prime for Testing
	s := []int{2, 3, 4, 5}

	// Test Rounds
	for i := 0; i < len(s); i++ {
		log.Printf("Matrix Size: %d x %d", s[i], s[i])

		// Act
		result := generateInternalMatrix(s[i], p)

		// Assert
		for j := 0; j < s[i]; j++ {
			expected := new(big.Int).SetInt64(int64(j + 1))
			// This is a diagonal matrix. Position value at i in Row i should be equal to i + 1
			if result[j][j].Cmp(expected) != 0 {
				t.Errorf("Row %d Position %d: expected %v, got %v", j, j, expected, result[j][j])
			}
		}

	}

}


func TestConvertHex(t *testing.T) {
	// Set up
	i := "0x11" // 17 as hex
	p := big.NewInt(19)
	// Act 
	result := convertHex(i, p)

	// Assert 
	expected := big.NewInt(17)

	if result.Cmp(expected) != 0 {
		t.Errorf("Expected %d, but result was %d", expected, result)
	}
}