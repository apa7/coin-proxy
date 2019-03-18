package ker_examples_test

func ExampleNewKerl() {}

// i req: length, The length of the Trits to squeeze out. Must be a multiple of 243.
// o: Trits, The squeezed out Trits.
// o: error, Returned for invalid lengths and internal errors.
func ExampleSqueeze() {}

// i req: in, The Trits slice to absorb. Must be a multiple of 243 in length.
// o: error, Returned for invalid slice lengths and internal errors.
func ExampleAbsorb() {}

func ExampleReset() {}

// i req: trits, The Trits to convert to []byte.
// o: []byte, The converted bytes.
// o: error, Returned for invalid lengths and internal errors.
func ExampleKerlTritsToBytes() {}

// i req: b, The bytes to convert.
// o: Trits, The converted Trits.
// o: error, Returned for invalid lengths and internal errors.
func ExampleKerlBytesToTrits() {}
