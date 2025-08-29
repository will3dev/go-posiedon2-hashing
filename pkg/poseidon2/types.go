package poseidon2

import "math/big"

// Need a type for the field params
type FieldParams struct {
	P *big.Int // This is the field value. 
	D uint64// This is the value that is used as exponent
}


// field to track the number of rounds
type Rounds struct {
	Rf int // this is a full round user for external
	Rp int // used to track the internal rounds which are partial
}


// type for the matrices
type LinearLayers struct {
	Me [][] *big.Int // external mixing matrix
	Mi [][] *big.Int // internal mixing matrix 
}


// type to track the constants
type Constants struct {
	// [Rf][t]
	Ce [][] *big.Int // used for external rounds. Sized based on number of rounds and size T
	// [Rp]
	Ci [] *big.Int // used for internal rounds. only use one value. Sized based on number of internal rounds
}

// type for all the poseidon params
type Poseidon2Params struct {
	T int // the size of the input and as a result the matrices
	FP FieldParams
	R Rounds 
	M LinearLayers
	C Constants 
}