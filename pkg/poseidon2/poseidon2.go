package poseidon2

import (
	"log"
	"math/big"
)

// This will be used to actually perform the operation

// This function assumes that the input of constants are ONLY the constants needing to be modified
func addFullRoundConstants(p *big.Int, x []*big.Int, extConst []*big.Int) []*big.Int {
	if len(x) != len(extConst) {
		panic("Input length does not equal constant width")
	}

	for i := 0; i < len(extConst); i++ {
		x[i].Add(x[i], extConst[i])
		x[i].Mod(x[i], p)
	}

	return x
}

func matrixMultiply(p *big.Int, x []*big.Int, extMatrix [][]*big.Int) []*big.Int {
	output := make([]*big.Int, len(x))

	// each row of the matrix will be added with the the three values
	// The sum of those values mod p, takes the place of x[i] with i representing the row number
	for i := 0; i < len(x); i++ {
		acc := new(big.Int)
		for j := 0; j < len(extMatrix[i]); j++ {
			prod := new(big.Int).Mul(extMatrix[i][j], x[j])
			acc = acc.Add(acc, prod)
		}

		output[i] = acc.Mod(acc, p)
	}

	return output

}

func Poseidon2Permutation(x []*big.Int) []*big.Int {
	var params Poseidon2Params
	if len(x) == 3 {
		params = poseidon2ParamsT3
	} else if len(x) == 5 {
		panic("Not available yet")
	} else {
		panic("Input length unsupported")
	}

	log.Println("Starting Poseidon2 Hash......")
	log.Println("Input values: ")
	for pos, val := range x {
		log.Println("Postion: ", pos, "Value: ", val)
	}

	output := make([]*big.Int, params.T)
	copy(output, x)

	
	// PRE-MIX
	// Before the first external round, Posiedon2 applies a linear layer
	log.Println("Starting Pre-Mix..........")
	for pos, val := range matrixMultiply(params.FP.P, output, params.M.Me) {
		output[pos] = val
	}

	log.Println("Values after pre-mix......")
	for pos, val := range output {
		log.Println("Postion: ", pos, "Value: ", val)
	}

	
	// FIRST SET OF EXTERNAL ROUNDS

	log.Println("Starting External Round 1......")
	ext1Const := params.C.Ce[:len(params.C.Ce)/2]

	log.Println("ext1Const: ", ext1Const)
	
	for i := 0; i < len(ext1Const); i++ {
		// Handle the constants
		for pos, val := range addFullRoundConstants(params.FP.P, output, ext1Const[i]) {
		output[pos] = val
		}

		// handle the S-Box
		for j := 0; j < len(output); j++ {
			output[j] = new(big.Int).Exp(output[j], new(big.Int).SetUint64(params.FP.D), params.FP.P)
		}

		// perform matrix multiplication
		for pos, val := range matrixMultiply(params.FP.P, output, params.M.Me) {
			output[pos] = val
		}

	}
	
	log.Println("Values after external round 1......")
	for pos, val := range output {
		log.Println("Postion: ", pos, "Value: ", val)
	}

	// INTERNAL ROUND
	log.Println("Starting Internal Round......")
	
	// Set up loop for all of the internal rounds
	// should be length of the internal constants
	for i := 0; i < len(params.C.Ci); i++ {
		// add constant to position [0]
		output[0] = new(big.Int).Add(output[0], params.C.Ci[i])
		output[0] = new(big.Int).Mod(output[0], params.FP.P)

		// S-Box the first value
		output[0] = new(big.Int).Exp(output[0], new(big.Int).SetUint64(params.FP.D), params.FP.P)

		// multiply the matrix by the internal linear layer
		for pos, val := range matrixMultiply(params.FP.P, output, params.M.Mi) {
			output[pos] = val
		}
	} 
	
	log.Println("Values after internal round......")
	for pos, val := range output {
		log.Println("Postion: ", pos, "Value: ", val)
	}

	// EXTERNAL ROUND 2

	log.Println("Starting External Round 2......")

	ext2Const := params.C.Ce[len(params.C.Ce)/2:]

	for i := 0; i < len(ext2Const); i++ {
		// Handle the constants
		for pos, val := range addFullRoundConstants(params.FP.P, output, ext2Const[i]) {
			output[pos] = val
		}

		// handle the S-Box
		for j := 0; j < len(output); j++ {
			output[j] = new(big.Int).Exp(output[j], new(big.Int).SetUint64(params.FP.D), params.FP.P)
		}

		// perform matrix multiplication
		for pos, val := range matrixMultiply(params.FP.P, output, params.M.Me) {
			output[pos] = val
		}

	}

	log.Println("Values after external round 2......")
	for pos, val := range output {
		log.Println("Postion: ", pos, "Value: ", val)
	}

	return output
}

// Used to generate the poseidon hash of a set of inputs.
func Poseidon2Hash(x []*big.Int) *big.Int {
	out := Poseidon2Permutation(x)
	// The hash value is just position 0
	return out[0]
}