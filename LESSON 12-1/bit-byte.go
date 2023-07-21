package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("Min int8 : %d - Max int8 : %d\n", math.MinInt8, math.MaxInt8)

	fmt.Printf("Min Uint8 : %d - Max Uint8 : %d\n", 0, math.MaxUint8)

	fmt.Printf("Min int16 : %d - Max int16 : %d\n", math.MinInt16, math.MaxInt16)

	fmt.Printf("Min Uint16 : %d - Max Uint16 : %d\n", 0, math.MaxUint16)

	fmt.Printf("Min int32 : %d - Max int32 : %d\n", math.MinInt32, math.MaxInt32)

	fmt.Printf("Min Uint32 : %d - Max Uint32 : %d\n", 0, math.MaxUint32)

	fmt.Printf("Min int64 : %d - Max int64 : %d\n", math.MinInt64, math.MaxInt64)

	fmt.Printf("Min Uint64 : 0 - Max Uint64 : 18446744073709551615\n")
}
