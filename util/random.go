package util

import (
	"math/rand"
	"time"
)

var myRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandNum ...
func RandNum(num int32) int32 {
	return myRand.Int31n(num)
}

//生成不重复的随机数，随机数范围0~max,数量num
func GenDiffRandomNum(num, max int) (outIntArr []int) {
	if num >= max {
		for i := 0; i < int(max); i++ {
			outIntArr = append(outIntArr, i)
		}
		return
	}
	var tmpArr []int
	for i := 0; i < int(max); i++ {
		tmpArr = append(tmpArr, 1)
	}

	var rnd int

	for i := 0; i < int(num); i++ {
		for {
			rnd = int(RandNum(int32(max)))
			if tmpArr[rnd] != -1 {
				break
			}
		}
		outIntArr = append(outIntArr, rnd)
		tmpArr[rnd] = -1
	}
	return
}