package tools

/**
 * @Author: zhenzhongzhu
 * @Description:工具包
 * @File: tools
 * @Date: 2024/5/18 12:35
 */

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateRandomString
// @Author ZhenZhen Zhu
// @Description: 生成随机字符串
// @param ：length
// @return ：string
func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateRandomNum(start int, end int) int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(end - start)
	random = start + random
	return random
}
