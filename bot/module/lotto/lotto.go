package lotto

import (
	"math/rand"
	"strconv"
)

// Generate lotto numbers
func GenLottoNum() string {
	var (
		sArr []int  = []int{0, 0, 0, 0, 0, 0}
		sStr string = ""
	)

	for i := 0; i < 6; i++ {
		sArr[i] = rand.Intn(45)

		for j := 0; j < i; j++ {
			if sArr[i] == sArr[j] {
				i--
			}
		}
		sStr += strconv.Itoa(sArr[i]) + " "
	}

	return sStr
}
