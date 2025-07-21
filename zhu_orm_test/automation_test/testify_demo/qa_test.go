package testify_demo

/**
 * @Author: zhenzhongzhu
 * @Description:
 * @File: qatest
 * @Date: 2025/7/21 下午7:07
 */

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddWithTestify(t *testing.T) {
	result := Add(1, 2)
	assert.Equal(t, 3, result, "Add(1, 2) should return 3")
}

func Add(a, b int) int {
	return a + b
}
