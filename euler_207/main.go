package euler_207

import (
	"fmt"
	"math"
)

func main() {
	var input int
	fmt.Scanf("%d", &input)
	var lms []int
	var a int
	var b int
	for i := 0; i < input; i++ {
		fmt.Scanf("%d", &a)
		fmt.Scanf("%d", &b)
		lms = append(lms, getLeastM(a, b))
	}

	//log.Printf("total:%v,perfect:%v", total, perfect)

	for _, v := range lms {
		fmt.Println(v)
	}
}
func getLeastM(a int, b int) int {
	m := math.MaxInt64

	//k := 2
	/**
		4^t = 2^t  + k
		k starts from 1
		2^t starts from 3
		k<=m
 */
	total := 1
	perfect := 1
	for l := 3; l <= m; l++ {
		ktemp := l*l - l

		if ktemp <= m {
			//check if log l is integer
			base := l & (l - 1)
			if base == 0 {
				perfect += 1
				//	b+=b
			}
			total += 1

			//a+=a

			if (a * total) > (b * perfect) {
				return ktemp
			}

		}

	}
	return -1
}
