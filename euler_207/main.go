package euler_207

import (
	"math"
	"log"
)

func main() {
	m := 185
	//k := 2
	total := 1
	perfect := 1

	/**
	4^t = 2^t  + k
	k starts from 1
	2^t starts from 3
	k<=m
	 */

	for l := 3; l <= m; l++ {
		ktemp := l*l - l

		if ktemp <= m {
			//check if log l is integer
			base := math.Log2(float64(l))
			if base == math.Trunc(base) {
				perfect += 1
			}
			total += 1
		}

	}
	log.Printf("total:%v,perfect:%v", total, perfect)
}
