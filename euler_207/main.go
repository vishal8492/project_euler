package main

import (
	"fmt"
	"math"
)

func main() {
	var input int
	fmt.Scanf("%d", &input)
	var lms = make([]int, input)
	var ms = make([]float64, input)

	for i := 0; i < input; i++ {
		//fmt.Printf("Inside scan")
		var a int
		var b int
		fmt.Scanf("%d", &a)
		fmt.Scanf("%d", &b)
		ms[i] = float64(a) / float64(b)
	}
	m := math.MaxInt32
	//k := 2

	/**
4^t = 2^t  + k
k starts from 1
2^t starts from 3
k<=m
 */
	mtemp := 3
	cnt := 0

	for mtemp != m {
		total := 1
		perfect := 1
		for l := 3; l <= m; l++ {
			ktemp := l*l - l

			if ktemp <= mtemp {
				//check if log l is integer
				base := math.Log2(float64(l))
				if base == math.Trunc(base) {
					perfect += 1
				}
				total += 1

				for i, thresh := range ms {
					if thresh > (float64(perfect) / float64(total)) {
						if lms[i] == 0 {
							lms[i] = mtemp
							cnt += 1
						}
					}
				}

			} else {
				break
			}
			if cnt == input {
				break
			}

		}
		//log.Printf("total:%v,perfect:%v", total, perfect)

		mtemp += 1
		if cnt == input {
			break
		}
	}

	for _, v := range lms {
		fmt.Println(v)
	}
}
