package poseidon2

import (
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
