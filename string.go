package wlike

import (
	"math/rand"
	"time"
)

type Str struct{}

// 随机生成字符串
func RandomString(num int) (str string) {
	base := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base_bytes := []byte(base)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	v := [num]byte{}

	for i := 0; i < num; i++ {
		v[i] = base_bytes[random.Intn(len(base))]
	}
	str = string(v[:])
	return
}







