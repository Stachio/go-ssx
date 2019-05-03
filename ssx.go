package ssx

import (
	"log"
	"math/rand"

	"github.com/Stachio/go-printssx"
)

// Printer - Generic printer object provided by stachio/printerssx
var Printer = printssx.New("SSX", log.Println, log.Printf, printssx.Moderate, printssx.Loud)

// BinaryResult - Return type for BinarySearchRecursion
type BinaryResult int

// BinaryFail - Causes BinarySearchRecursion to fail
const BinaryFail BinaryResult = 0

// BinarySuccess - Causes BinarySearchRecursion to succeed
const BinarySuccess BinaryResult = 1

// BinaryContinue - Allows BinarySearchRecursion to continue on to the next index
const BinaryContinue BinaryResult = 2

// BinarySearchRecursion - Walks through n-indices randomly, uniquely
func BinarySearchRecursion(startIndex, endIndex, level uint64, op func(uint64) BinaryResult) (out uint64, result BinaryResult) {
	// Randomly found a node
	Printer.Printf(printssx.Loud, "Binarily checking %d/%d/%d", startIndex, endIndex, level)
	if startIndex == endIndex {
		Printer.Println(printssx.Moderate, "Operating on index", startIndex)
		return startIndex, op(startIndex)
	}

	midIndex := (startIndex + endIndex) / 2
	hilo := rand.Uint32() & 1
	lowSet := [2]uint64{startIndex, midIndex}
	highSet := [2]uint64{midIndex + 1, endIndex}

	var setA *[2]uint64
	var setB *[2]uint64

	if hilo == 1 {
		// Search hi
		setA = &highSet
		setB = &lowSet
	} else {
		// Search lo
		setA = &lowSet
		setB = &highSet
	}

	out, result = BinarySearchRecursion(setA[0], setA[1], level+1, op)
	if result == BinaryContinue {
		out, result = BinarySearchRecursion(setB[0], setB[1], level+1, op)
	}
	return
}
