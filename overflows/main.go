package main

import (
	"fmt"
	"math"
)

func main() {
	var minInt8 int8 = int8(math.MinInt8)
	var maxInt8 int8 = int8(math.MaxInt8)
	fmt.Println(minInt8, maxInt8)

	// When int underflows value is reset to the max value
	// When int overflows value is reset to the min value
	minInt8--
	maxInt8++
	fmt.Println(minInt8, maxInt8)

	var maxFloat32 float32 = float32(math.MaxFloat32)
	fmt.Println(maxFloat32)

	// When float overflows value is reset to +Inf (infinity)
	maxFloat32 = maxFloat32 * 1.256
	fmt.Println(maxFloat32)
}
