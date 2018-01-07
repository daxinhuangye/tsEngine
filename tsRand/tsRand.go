package tsRand

import (
	"math/rand"
	"time"
)

type RandMaker struct {
}

func Seed(value int64) {
	rand.Seed(value)
}

/*
 * @brief 获得随机数，在设定范围
 *
 * @param min 最小
 * @param max 最大
 * @return 随机数
 */
func RandInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

//随机数获取
func RandNum(intn int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(intn)
}
