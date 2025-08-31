package poseidon2

import (
	"math/big"
	"testing"
)

func TestConsistPoseidonHash(t *testing.T) {

}

func TestDifferentInputs(t *testing.T) {

}

func TestFieldBoundsTest(t *testing.T) {
	// This will verify that all ouputs of Poseidon Permutation are
	// within Fp meaning 0 - p-1

	p, _ := new(big.Int).SetString("21888242871839275222246405745257275088548364400416034343698204186575808495617", 10)

	inputs := make([]*big.Int, 3)
	inputs[0] = new(big.Int).SetInt64(1)
	inputs[1] = new(big.Int).SetInt64(2)
	inputs[2] = new(big.Int).SetInt64(3)

	output := Poseidon2Permutation(inputs)

	for i, val := range output {
		if val.Cmp(p) >= 0 {
			t.Errorf("Output %d (%v) is not in field (>= %v)", i, val, p)
		}
		if val.Sign() < 0 {
			t.Errorf("Output %d (%v) is negative", i, val)
		}
	}
}
